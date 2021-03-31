package db

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/machinebox/graphql"
)

func getBalancerData(database *Database, uniswapreqdata UniswapInputStruct) {
	clientBalancer := graphql.NewClient("https://api.thegraph.com/subgraphs/name/balancer-labs/balancer")

	// fmt.Println(clientBalancer.Log)
	// 2 - declare queries
	reqBalancerListOfPools := graphql.NewRequest(`
	query {
        pools{
        id
			tokens {
			id
			symbol
			}
      	}
	 }
	`)

	reqBalancerByPoolID := graphql.NewRequest(`
	query ($poolid:String!){
		pool(id:$poolid) {
			id
			finalized
			publicSwap
			swapFee
			totalSwapVolume
			totalWeight
			tokensList
			tokens {
				id
				address
				balance
				decimals
				symbol
				denormWeight
			}
		}
	  }
`)

	reqBalancerListOfPools.Var("key", "value")
	reqBalancerListOfPools.Header.Set("Cache-Control", "no-cache")
	reqBalancerByPoolID.Header.Set("Cache-Control", "no-cache")
	ctx := context.Background()

	var respBalancerPoolList BalancerPoolList
	var respBalancerById BalancerById

	var respUniswapTicker UniswapTickerQuery // Used in Balancer to look up Uniswap IDs of 'ETH' etc
	var respUniswapHist UniswapHistQuery
	//	var respUniswapById UniswapCurrentQuery

	var BalancerFilteredPoolList []string      // Pairs - IDS - 0x124145
	var BalancerFilteredPoolListPairs []string // Pairs - Tokens ETH/DAI
	var BalancerFilteredTokenList []string     // Tokens - ETH, DAI

	if err := clientBalancer.Run(ctx, reqBalancerListOfPools, &respBalancerPoolList); err != nil {
		log.Fatal(err)
	}

	// Process received list of pools (PAIRS)
	for i := 0; i < len(respBalancerPoolList.Pools); i++ {
		if len(respBalancerPoolList.Pools[i].Tokens) > 1 {
			token0symbol := respBalancerPoolList.Pools[i].Tokens[0].Symbol
			token1symbol := respBalancerPoolList.Pools[i].Tokens[1].Symbol

			if isPoolPartOfFilter(token0symbol, token1symbol) {
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
					if !isHistDataAlreadyDownloaded(convBalancerToken(tokenqueue[j]), database) {
						// Get Uniswap Ids of these tokens
						uniswapreqdata.reqUniswapIDFromTokenTicker.Var("ticker", convBalancerToken(tokenqueue[j]))
						if err := uniswapreqdata.clientUniswap.Run(ctx, uniswapreqdata.reqUniswapIDFromTokenTicker, &respUniswapTicker); err != nil {
							log.Fatal(err)
						}
						// Download historical data for each token for which data is missing
						if len(respUniswapTicker.IDsforticker) >= 1 {
							// request data from uniswap using this queried ticker
							uniswapreqdata.reqUniswapHist.Var("tokenid", setUniswapQueryIDForToken(tokenqueue[j], respUniswapTicker.IDsforticker[0].ID))

							fmt.Print("Querying historical data for: ")
							fmt.Print(tokenqueue[j])
							if err := uniswapreqdata.clientUniswap.Run(ctx, uniswapreqdata.reqUniswapHist, &respUniswapHist); err != nil {
								log.Fatal(err)
							}

							fmt.Print("| returned days: ")
							fmt.Println(len(respUniswapHist.DailyTimeSeries))

							// if returned data - append it to database
							if len(respUniswapHist.DailyTimeSeries) > 0 {
								// Append to database
								database.historicalcurrencydata = append(database.historicalcurrencydata, NewHistoricalCurrencyDataFromRaw(tokenqueue[j], respUniswapHist.DailyTimeSeries))
							}
						} // if managed to find some IDs for this TOKEN
					} // if historical data needs updating
				} // tokenqueue loop ends

				// if historical data is in order - get current data
				reqBalancerByPoolID.Var("poolid", respBalancerPoolList.Pools[i].ID)

				if err := clientBalancer.Run(ctx, reqBalancerByPoolID, &respBalancerById); err != nil {
					log.Fatal(err)
				}

				currentSize := float32(10000.0)                                              // TO DO: Size
				currentVolume, _ := strconv.ParseFloat(respBalancerById.TotalSwapVolume, 32) // No historical for now
				currentInterestrate := float32(0.00)                                         // Zero for liquidity pool
				BalancerRewardPercentage := float32(0.003)                                   // TO DO

				volatility := calculatehistoricalvolatility(retrieveDataForTokensFromDatabase(token0symbol, token1symbol, database), 30)
				ROI_raw_est := calculateROI_raw_est(currentInterestrate, BalancerRewardPercentage, float32(currentVolume), volatility)

				var recordalreadyexists bool
				recordalreadyexists = false

				// CHECK IF NOT DUPLICATING RECORD - IF ALREADY EXISTS - UPDATE NOT APPEND
				for k := 0; k < len(database.currencyinputdata); k++ {
					// Means record already exists - UPDATE IT, DO NOT APPEND
					if database.currencyinputdata[k].Pair == token0symbol+"/"+token1symbol && database.currencyinputdata[k].Pool == "Balancer" {
						recordalreadyexists = true
						database.currencyinputdata[k].PoolSize = currentSize
						database.currencyinputdata[k].PoolVolume = float32(currentVolume)

						database.currencyinputdata[k].ROI_raw_est = ROI_raw_est
						database.currencyinputdata[k].ROI_vol_adj_est = 0
						database.currencyinputdata[k].ROI_hist = 0					
						
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
}
