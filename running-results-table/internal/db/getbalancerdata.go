package db

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/machinebox/graphql"
)

func estimate_future_balancer_volume_and_pool_sz(histvolume BalancerHistVolumeQuery) (float32, float32) {
	future_volume_est := 0.0
	future_sz_est := 0.0

	var count float64
	var count_sz float64
	count = 0
	count_sz = 0

	for i := 0; i < len(histvolume.Pool.Swaps); i++ {
		/*
			fmt.Print("ESTIMATING BALANCER FUTURE VOLUME + POOL SZ: ")
			fmt.Print("transaction volume: ")
			fmt.Print(histvolume.Pool.Swaps[i].TokenAmountIn)
			fmt.Print(" | pool sz (liquidity): ")
			fmt.Println(histvolume.Pool.Swaps[i].PoolLiquidity)
		*/
		v, _ := strconv.ParseFloat(histvolume.Pool.Swaps[i].TokenAmountIn, 64) // double check the TokenAmounIn, only used to compile;
		sz, _ := strconv.ParseFloat(histvolume.Pool.Swaps[i].PoolLiquidity, 64)

		future_volume_est += v
		future_sz_est += sz

		if v != 0.0 {
			count++
		}

		if sz != 0.0 {
			count_sz++
		}

	}

	// APPLY ADJUSTOR? 	// MEDIAN?	// TAKE OUT EXTREME VALUES TO NORMALISE?
	if count > 0 {
		future_volume_est = future_volume_est / count
	} else {
		future_volume_est = 0.0
	}

	if count_sz > 0 {
		future_sz_est = future_sz_est / count_sz
	} else {
		future_sz_est = 0.0
	}

	if math.IsNaN(float64(future_volume_est)) {
		// should never happen
		fmt.Println("ERROR IN FUTURE VOLUME - 999999999999999999555555555555555555")
		future_volume_est = -995.0
	}
	if math.IsNaN(float64(future_sz_est)) {
		// should never happen
		fmt.Println("ERROR IN FUTURE SZ - 999999999999999999666666666666666666")
		future_sz_est = -996.0
	}

	if math.IsInf(float64(future_volume_est), 0) {
		fmt.Println("ERROR IN FUTURE VOLUME - 999999999999999999555555555555555555")
		future_volume_est = -993.0
	}
	if math.IsInf(float64(future_sz_est), 0) {
		fmt.Println("ERROR IN FUTURE SZ - 999999999999999999666666666666666666")
		future_sz_est = -994.0
	}
	/*
		fmt.Print("Future volume est: ")
		fmt.Print(future_volume_est)
		fmt.Print(" | ")
		fmt.Print("Future sz est: ")
		fmt.Print(future_sz_est)
	*/
	return float32(future_volume_est), float32(future_sz_est) // USD
}

func getBalancerData(database *Database, uniswapreqdata UniswapInputStruct) {

	fmt.Println("trying to get balancer data")
	clientBalancer := graphql.NewClient("https://api.thegraph.com/subgraphs/name/balancer-labs/balancer")

	// 2 - declare queries
	reqBalancerListOfPools := graphql.NewRequest(`
	query {
		pools(first: 100, orderDirection: desc, orderBy: liquidity, where: {publicSwap: true}) {
		  id
		  tokensList
		  tokens {
			id
			address
			balance
			symbol
		  }
		}
	  }
	`)

	reqBalancerByPoolID := graphql.NewRequest(`
		query ($poolid:String!){
			pool(id:$poolid) {
				id
				swapFee
				totalSwapVolume
				liquidity
				totalWeight
				tokensList
				tokens {
					id
					address
					balance
					symbol
				}	
			}
		}
 	`)

	// get historical volume
	reqBalancerHistVolume := graphql.NewRequest(`
	query($pairid:String!){
		pool(id:$pairid) {
			swaps(first: 1000, skip: 0, orderBy: timestamp, orderDirection: desc){
				timestamp
				feeValue
				tokenInSym
				tokenOutSym
				tokenIn
				tokenOut
				tokenAmountIn
				tokenAmountOut
				poolLiquidity
			}
		}  
	}
`)

	// get TVL
	// get this pool % TVL
	// get BAL token price

	reqBalancerListOfPools.Var("key", "value")
	reqBalancerListOfPools.Header.Set("Cache-Control", "no-cache")
	reqBalancerByPoolID.Header.Set("Cache-Control", "no-cache")

	reqBalancerHistVolume.Var("key", "value")
	reqBalancerHistVolume.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	var respBalancerPoolList BalancerPoolList
	var respBalancerById BalancerById
	var respBalancerHistVolume BalancerHistVolumeQuery

	var respUniswapTicker UniswapTickerQuery // Used in Balancer to look up Uniswap IDs of 'ETH' etc
	var respUniswapHist UniswapHistQuery

	var BalancerFilteredPoolList []string      // Pairs - IDS - 0x124145
	var BalancerFilteredPoolListPairs []string // Pairs - Tokens ETH/DAI
	var BalancerFilteredTokenList []string     // Tokens - ETH, DAI

	var Histrecord HistoricalCurrencyData

	if err := clientBalancer.Run(ctx, reqBalancerListOfPools, &respBalancerPoolList); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(respBalancerPoolList.Pools); i++ {
		fmt.Print("i: ")
		fmt.Print(i)
		fmt.Print(" | ")
		fmt.Print(respBalancerPoolList.Pools[i].Tokens[0].Symbol)
		fmt.Print(" | ")
		fmt.Println(respBalancerPoolList.Pools[i].Tokens[1].Symbol)
	}

	// Process received list of pools (PAIRS)
	for i := 0; i < len(respBalancerPoolList.Pools); i++ {
		if len(respBalancerPoolList.Pools[i].Tokens) > 1 {
			token0symbol := respBalancerPoolList.Pools[i].Tokens[0].Symbol
			token1symbol := respBalancerPoolList.Pools[i].Tokens[1].Symbol

			if isPoolPartOfFilter(token0symbol, token1symbol) {

				fmt.Print("t0: ")
				fmt.Print(token0symbol)
				fmt.Print("| t1: ")
				fmt.Println(token1symbol)

				// Filter pools to allowed components (WETH, DAI, USDC, USDT)
				BalancerFilteredPoolList = append(BalancerFilteredPoolList, respBalancerPoolList.Pools[i].ID)
				BalancerFilteredPoolListPairs = append(BalancerFilteredPoolListPairs, token0symbol+"/"+token1symbol)

				var tokenqueue []string

				// Split list of pairs into single tokens
				if !stringInSlice(token0symbol, BalancerFilteredTokenList) {
					BalancerFilteredTokenList = append(BalancerFilteredTokenList, token0symbol)
					tokenqueue = append(tokenqueue, token0symbol)
				}
				if !stringInSlice(token1symbol, BalancerFilteredTokenList) {
					BalancerFilteredTokenList = append(BalancerFilteredTokenList, token1symbol)
					tokenqueue = append(tokenqueue, token1symbol)
				}

				for j := 0; j < len(tokenqueue); j++ {
					// Check if database already has historical data
					if !isHistDataAlreadyDownloadedDatabase(convBalancerToken(tokenqueue[j])) {
						// Get Uniswap Ids of these tokens
						uniswapreqdata.reqUniswapIDFromTokenTicker.Var("ticker", convBalancerToken(tokenqueue[j]))
						if err := uniswapreqdata.clientUniswap.Run(ctx, uniswapreqdata.reqUniswapIDFromTokenTicker, &respUniswapTicker); err != nil {
							log.Fatal(err)
						}
						// Download historical data for each token for which data is missing
						if len(respUniswapTicker.IDsforticker) >= 1 {
							// request data from uniswap using this queried ticker
							uniswapreqdata.reqUniswapHist.Var("tokenid", setUniswapQueryIDForToken(tokenqueue[j], respUniswapTicker.IDsforticker[0].ID))

							fmt.Print("Querying historical (in GETBALANCER) data from UNISWAP for: ")
							fmt.Print(tokenqueue[j])
							if err := uniswapreqdata.clientUniswap.Run(ctx, uniswapreqdata.reqUniswapHist, &respUniswapHist); err != nil {
								log.Fatal(err)
							}

							fmt.Print("| returned days: ")
							fmt.Println(len(respUniswapHist.DailyTimeSeries))

							// if returned data - append it to database
							if len(respUniswapHist.DailyTimeSeries) > 0 {
								Histrecord = NewHistoricalCurrencyDataFromRaw(tokenqueue[j], respUniswapHist.DailyTimeSeries)
								appendDataForTokensFromDatabase(Histrecord)
							}
						} // if managed to find some IDs for this TOKEN
					} // if historical data needs updating
				} // tokenqueue loop ends

				// if historical data is in order - get current data
				reqBalancerByPoolID.Var("poolid", respBalancerPoolList.Pools[i].ID)

				if err := clientBalancer.Run(ctx, reqBalancerByPoolID, &respBalancerById); err != nil {
					log.Fatal(err)
				}

				currentSize, _ := strconv.ParseFloat(respBalancerById.Pool.Liquidity, 32)         // TO DO: Size
				currentVolume, _ := strconv.ParseFloat(respBalancerById.Pool.TotalSwapVolume, 32) // No historical for now

				// should we move this to historical section + add to database?
				// request from blockchain directly?
				fmt.Println("requesting data for id: ")
				fmt.Print(respBalancerPoolList.Pools[i].ID)

				reqBalancerHistVolume.Var("pairid", respBalancerPoolList.Pools[i].ID)

				if err := clientBalancer.Run(ctx, reqBalancerHistVolume, &respBalancerHistVolume); err != nil {
					log.Fatal(err)
				}

				fmt.Print("Queried historical volume from BALANCER - number of items: ")
				fmt.Println(len(respBalancerHistVolume.Pool.Swaps))

				future_daily_volume_est, future_pool_sz_est := estimate_future_balancer_volume_and_pool_sz(respBalancerHistVolume)
				historical_pool_sz_avg, historical_pool_daily_volume_avg := future_pool_sz_est, future_daily_volume_est

				currentInterestrate := float32(0.00)                                                 // Zero for liquidity pool
				BalancerRewardPercentage, _ := strconv.ParseFloat(respBalancerById.Pool.SwapFee, 32) // TO DO
				volatility := calculatehistoricalvolatility(retrieveDataForTokensFromDatabase2(token0symbol, token1symbol), 30)

				imp_loss_hist := estimate_impermanent_loss_hist(volatility, 1, "Balancer")
				px_return_hist := calculate_price_return_x_days(Histrecord, 30)

				ROI_raw_est := calculateROI_raw_est(currentInterestrate, float32(BalancerRewardPercentage), float32(future_pool_sz_est), float32(future_daily_volume_est), imp_loss_hist)      // + imp
				ROI_vol_adj_est := calculateROI_vol_adj(ROI_raw_est, volatility)                                                                                                               // Sharpe ratio
				ROI_hist := calculateROI_hist(currentInterestrate, float32(BalancerRewardPercentage), historical_pool_sz_avg, historical_pool_daily_volume_avg, imp_loss_hist, px_return_hist) // + imp + hist

				var recordalreadyexists bool
				recordalreadyexists = false

				// CHECK IF NOT DUPLICATING RECORD - IF ALREADY EXISTS - UPDATE NOT APPEND
				for k := 0; k < len(database.currencyinputdata); k++ {
					// Means record already exists - UPDATE IT, DO NOT APPEND
					if database.currencyinputdata[k].Pair == token0symbol+"/"+token1symbol && database.currencyinputdata[k].Pool == "Balancer" {
						recordalreadyexists = true
						database.currencyinputdata[k].PoolSize = float32(currentSize)
						database.currencyinputdata[k].PoolVolume = float32(currentVolume)

						database.currencyinputdata[k].ROI_raw_est = ROI_raw_est
						database.currencyinputdata[k].ROI_vol_adj_est = ROI_vol_adj_est
						database.currencyinputdata[k].ROI_hist = ROI_hist

						database.currencyinputdata[k].Volatility = volatility
						database.currencyinputdata[k].Yield = currentInterestrate
					}
				}

				// APPEND IF NEW
				if !recordalreadyexists {
					database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{token0symbol + "/" + token1symbol, float32(currentSize),
						float32(currentVolume), currentInterestrate, "Balancer", volatility, ROI_raw_est, 0.0, 0.0})
				}
				// fmt.Println("APPENDED BALANCER DATA")
			} // if pool is within pre filtered list ends
		} // if pool has some tokens ends
	} // balancer pair loop closes

	fmt.Println("BALANCER COMPLETED!!!!!")

} // balancer get data close
