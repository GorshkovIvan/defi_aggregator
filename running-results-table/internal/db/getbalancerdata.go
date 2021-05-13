package db

/*
Check balance units - ok
Add database linkage - ok
Check ROI result makes sense
Add BAL token return component

Get historical balances
Volumes as floats
Missing zeros - fix
*/

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"pusher/defi_aggregator/running-results-table/internal/db/token"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinebox/graphql"
)

/*
func BoD(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
*/

func estimate_future_balancer_volume_and_pool_sz(dates []int64, tradingvolumes []int64, poolsizes []int64) (float32, float32) {
	future_volume_est := 0.0
	future_sz_est := 0.0

	var count float64
	var count_sz float64
	count = 0
	count_sz = 0

	for i := 0; i < len(dates); i++ {
		/*
			fmt.Print("ESTIMATING BALANCER FUTURE VOLUME + POOL SZ: ")
			fmt.Print("transaction volume: ")
			fmt.Print(histvolume.Pool.Swaps[i].TokenAmountIn)
			fmt.Print(" | pool sz (liquidity): ")
			fmt.Println(histvolume.Pool.Swaps[i].PoolLiquidity)
		*/
		//v, _ := strconv.ParseFloat(histvolume.Pool.Swaps[i].TokenAmountIn, 64) // double check the TokenAmounIn, only used to compile;
		//sz, _ := strconv.ParseFloat(histvolume.Pool.Swaps[i].PoolLiquidity, 64)

		v := tradingvolumes[i]
		sz := poolsizes[i]

		future_volume_est += float64(tradingvolumes[i])
		fmt.Print(poolsizes[i])
		future_sz_est += float64(poolsizes[i]) // sz

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

	fmt.Print("Future volume est: ")
	fmt.Print(future_volume_est)
	fmt.Print(" | ")
	fmt.Print("Future sz est: ")
	fmt.Print(future_sz_est)

	return float32(future_volume_est), float32(future_sz_est) // USD
}

func getBalancerData(database *Database, uniswapreqdata UniswapInputStruct) {

	fmt.Println("trying to get balancer data")
	clientBalancer := graphql.NewClient("https://api.thegraph.com/subgraphs/name/balancer-labs/balancer")
	balancertopic := "0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378"

	// 0) Connect to client
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	// 2 - declare queries
	reqBalancerListOfPools := graphql.NewRequest(`
	query {
		pools(first: 50, orderDirection: desc, orderBy: liquidity, where: {publicSwap: true}) {
		  id
		  tokensList
		  tokens {
			id
			address
			balance
			symbol
			decimals
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
					decimals
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
	//	var respBalancerHistVolume BalancerHistVolumeQuery

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
		if len(respBalancerPoolList.Pools[i].Tokens) > 1 && len(respBalancerPoolList.Pools[i].Tokens) < 3 {
			token0symbol := respBalancerPoolList.Pools[i].Tokens[0].Symbol
			token1symbol := respBalancerPoolList.Pools[i].Tokens[1].Symbol

			if isPoolPartOfFilter(token0symbol, token1symbol) {
				fmt.Print("Pool at i: ")
				fmt.Print(i)
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
						fmt.Print("NOT FOUND DATA FOR TOKEN IN DATABASE..QUERYING UNISWAP: ")
						fmt.Print(convBalancerToken(tokenqueue[j]))
						fmt.Print("CHECKPOINT 777")

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

				/*
						fmt.Println("requesting data for id: ")
						fmt.Print(respBalancerPoolList.Pools[i].ID)

						reqBalancerHistVolume.Var("pairid", respBalancerPoolList.Pools[i].ID)

						if err := clientBalancer.Run(ctx, reqBalancerHistVolume, &respBalancerHistVolume); err != nil {
							log.Fatal(err)
						}
					//fmt.Print("Queried historical volume from BALANCER - number of items: ")
					//fmt.Println(len(respBalancerHistVolume.Pool.Swaps))
				*/

				/////////////////////////////////////////////////////////////////////////
				fmt.Print("checkpoint 0")
				days_ago := 1
				//oldest_available_record := time.Now() // XX - GET IT USING AARON's func
				//fmt.Print("What is this the correct Pool ID: ")
				var tokens []string
				tokens = append(tokens, respBalancerById.Pool.Tokens[0].Symbol)
				tokens = append(tokens, respBalancerById.Pool.Tokens[1].Symbol)
				oldest_available_record := time.Unix(get_newest_timestamp_from_db("Balancer", tokens), 0) //time.Unix(sec, nano)
				fmt.Print("NEWEST TIMESTAMP FROM DB:")
				fmt.Print(oldest_available_record)
				// fmt.Print("checkpoint 1")
				// oldest_available_record = oldest_available_record.AddDate(0, 0, -days_ago)
				oldest_lookup_time := time.Now() //.Unix()

				fmt.Print("TIME SINCE UPDATE: ")
				fmt.Print(time.Since(oldest_available_record).Hours())

				data_is_old := false

				fmt.Print("Checking if data is old")

				if (time.Since(oldest_available_record).Hours()) > 24 {
					data_is_old = true
					oldest_lookup_time = oldest_lookup_time.AddDate(0, 0, -days_ago) // oldest_available_record.Unix()
					// math.Max(now.AddDate(0, 0, -days_ago)
				}
				fmt.Print("  is data in Balancer db old: ")
				fmt.Println(data_is_old)

				var dates []int64
				var tradingvolumes []int64
				var poolsizes []int64
				var fees []int64
				var interest []float64
				var utilization []float64

				fmt.Print(data_is_old)
				if len(utilization) > 0 || len(fees) > 0 {
					fmt.Print("placeholder")
				}

				// if data is old - download it
				if data_is_old {
					fmt.Print("Data is old!! Downloading it!!")
					fmt.Print(" Now: ")
					fmt.Println(time.Now())
					//fmt.Print("BoD for today: ")
					//fmt.Println(BoD(time.Now()))
					//fmt.Print("diff: ")
					//t := (time.Since(oldest_available_record).Hours())

					// 1) If data is old and need to update it - Define pool specific parameters
					fmt.Println(" The Pool Address  IS: ")
					fmt.Print(respBalancerById.Pool.ID)
					var BalancerpoolAddress = common.HexToAddress(respBalancerById.Pool.ID)
					fmt.Println(" CONNECTED TO INFURA SUCCESSFULLY!!!!!!")

					tokenAddress0 := common.HexToAddress(respBalancerById.Pool.Tokens[0].Address)
					tokenAddress := common.HexToAddress(respBalancerById.Pool.Tokens[1].Address)

					instance, err := token.NewToken(tokenAddress, client)
					if err != nil {
						log.Fatal(err)
					}

					instance0, err := token.NewToken(tokenAddress0, client)
					if err != nil {
						log.Fatal(err)
					}

					bal0, err := instance0.BalanceOf(&bind.CallOpts{}, BalancerpoolAddress)
					if err != nil {
						log.Fatal(err)
					}

					bal, err := instance.BalanceOf(&bind.CallOpts{}, BalancerpoolAddress)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("bal: %s\n", bal)
					fbal := new(big.Float)
					fbal.SetString(bal.String())

					fmt.Printf("BAL0: %s\n", bal0)
					fbal0 := new(big.Float)
					fbal0.SetString(bal0.String())

					bal_float := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(respBalancerById.Pool.Tokens[1].Decimals))))
					bal_floatT2 := new(big.Float).Quo(fbal0, big.NewFloat(math.Pow10(int(respBalancerById.Pool.Tokens[0].Decimals)))) //---------------------------------------------------

					Histrecord_2 := retrieveDataForTokensFromDatabase2(token0symbol, token1symbol)
					//fmt.Print("Histrecord_2: ")
					//fmt.Print(Histrecord_2)
					//fmt.Print("found through maxintslice: ")
					//fmt.Print(MaxIntSlice(Histrecord_2.Date))
					//fmt.Print(" | Zeroth date: ")
					//fmt.Print(Histrecord_2.Date[0])
					//fmt.Print(" | Last indx date: ")
					fmt.Print(Histrecord_2.Date[len(Histrecord_2.Date)-1])
					fmt.Print(" | PX: ")
					fmt.Print(Histrecord_2.Price[len(Histrecord_2.Price)-1])
					fmt.Printf(" | balance: %f", bal_float)    // "balance: 74605500.647409"
					fmt.Printf(" | balance2: %f", bal_floatT2) //------------------------------------------------------------------

					var current_block *big.Int
					var oldest_block *big.Int
					current_block = big.NewInt(0)

					// Get current block
					header, err := client.HeaderByNumber(context.Background(), nil)
					if err != nil {
						log.Fatal(err)
					}

					current_block = header.Number

					fmt.Print("current block: ")
					fmt.Print(current_block)

					//2)  Find oldest block in our lookup date range
					oldest_block = new(big.Int).Set(current_block)

					j := int64(0) // compute block id [days_ago] days away from now
					for {
						j -= 2000
						oldest_block.Add(oldest_block, big.NewInt(j))

						block, err := client.BlockByNumber(context.Background(), oldest_block)
						if err != nil {
							log.Fatal(err)
						}

						if block.Time() <= uint64(oldest_lookup_time.Unix()) {
							fmt.Print("oldest lkp block: ")
							fmt.Println(oldest_block)
							fmt.Print(" | t: ")
							fmt.Print(block.Time())
							fmt.Print("| curr blk - oldest blk: ")
							diff := current_block.Sub(current_block, oldest_block)
							fmt.Println(diff)
							break
						}
					}

					//3)  Query between oldest and current block for Balancer-specific addresses
					query := ethereum.FilterQuery{
						FromBlock: oldest_block,
						ToBlock:   nil, // = latest block
						Addresses: []common.Address{BalancerpoolAddress},
					}

					logsX, err := client.FilterLogs(context.Background(), query)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Print("Number of block logs: ")
					fmt.Print(len(logsX))

					cumulative_for_day := int64(0)

					t_prev := uint64(0)
					t_new := uint64(0)

					//4)  Loop through received data and filter it again
					// For each transaction in logsX - check if it matches lookup criteria - add volume if does:
					for i := 0; i < len(logsX); i++ {
						if logsX[i].Topics[0] != common.HexToHash(balancertopic) {
							continue
						}

						fmt.Print("i: ")
						fmt.Print(i)
						//fmt.Print(" | th: ")
						//fmt.Print(logsX[i].TxHash)
						//fmt.Print(" | blk#: ")
						//fmt.Print(logsX[i].BlockNumber)
						//fmt.Print(" | ")
						if math.Mod(float64(i), 20) == 0 {
						// Get date from block number
						block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(logsX[i].BlockNumber)))
						if err != nil {
							log.Fatal(err)
						}
						t_prev = t_new       // uint
						t_new = block.Time() // uint
						} // if float 20
						//fmt.Print(" | t: ")
						//fmt.Print(block.Time())

						txlog, err := client.TransactionReceipt(context.Background(), logsX[i].TxHash)
						if err != nil {
							log.Fatal(err)
						}


						if t_prev == 0 || (t_new-uint64(math.Mod(float64(t_new), 86400)))/86400 != (t_prev-uint64(math.Mod(float64(t_prev), 86400)))/86400 { // 1 day
							// check if dates is already in db
							tn := int64(BoD(time.Unix(int64(t_new), 0)).Unix())
							if tn > 0 && tn > MaxIntSlice(dates) {
							dates = append(dates, tn)
							tradingvolumes = append(tradingvolumes, cumulative_for_day)
							bal_int, _ := bal_float.Int64()
							bal_intT2, _ := bal_floatT2.Int64()
							if len(Histrecord_2.Price) >= 1 {
								bal_int = bal_int                                                            //
								bal_intT2 = bal_intT2 * int64(Histrecord_2.Price[len(Histrecord_2.Price)-1]) //--------------------------------------------
							} // convert to token1
							//poolsizes = append(poolsizes, bal_int) // bal.Int64()
							poolsizes = append(poolsizes, bal_int+bal_intT2) // bal.Int64()
							cumulative_for_day = 0
						}
							//fmt.Println(" t new:")
							//fmt.Print(t_new)
							//fmt.Print(" | prev: ")
							//fmt.Print(t_prev)
							//fmt.Print("day crossed: ")
							//fmt.Print(int64(BoD(time.Unix(int64(t_new), 0)).Unix()))
							//fmt.Print("..cumulative: ")
							//fmt.Println(cumulative_for_day)
						} else { //-------------
							token0AD := respBalancerById.Pool.Tokens[0].Address //"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2" // WETH
							token1AD := respBalancerById.Pool.Tokens[1].Address
							dec := respBalancerById.Pool.Tokens[1].Decimals // 18 // t1 decimals
							tkn1 := getTradingVolumeFromTxLog2(txlog.Logs, balancertopic, token0AD, token1AD, dec)
							// convert to usd
							cumulative_for_day += tkn1
						}
					} // loop through log finished

					// otherwise get this array from db

					fmt.Println("-----------------SUMMARY DAILY: -----------------------------------")
					for i := 0; i < len(dates); i++ { //we should start at +1
						fmt.Print("i: ")
						fmt.Print(i)
						interest = append(interest, 0.0)
						fmt.Print("| t: ")
						fmt.Print(dates[i])
						fmt.Print("| volumes: ")
						fmt.Println(tradingvolumes[i])
						
						vlm := tradingvolumes[i]
						if vlm == 0.0 {vlm = 1000}
						
						//if len(interest[i])
						interest_x := float64(0) // interest[i]
						fees_x := int64(0)

						if vlm > 0.0 {
							// respBalancerById.Pool.Tokens[0].Symbol, respBalancerById.Pool.Tokens[1].Symbol
							recordID := append_record_to_database("Balancer", tokens, dates[i], vlm, poolsizes[i], fees_x, interest_x, float64(0)) //-----implemented Function
							if recordID == "x" {
							}
						} // respBalancerById.Pool.ID
						// fmt.Print(" xx ")
					}
					/////////////////////////////////////////////////////////////////////////
				} // if need to update data

				if !data_is_old { // else: data is not old
					dates, tradingvolumes, poolsizes, fees, interest, utilization = retrieve_hist_pool_sizes_volumes_fees_ir("Balancer", tokens)
				}

				future_daily_volume_est, future_pool_sz_est := estimate_future_balancer_volume_and_pool_sz(dates, tradingvolumes, poolsizes)
				// future_daily_volume_est, future_pool_sz_est := estimate_future_balancer_volume_and_pool_sz(respBalancerHistVolume)
				historical_pool_sz_avg, historical_pool_daily_volume_avg := future_pool_sz_est, future_daily_volume_est

				currentInterestrate := float32(0.00) // Zero for liquidity pool
				BalancerRewardPercentage, _ := strconv.ParseFloat(respBalancerById.Pool.SwapFee, 32)
				volatility := calculatehistoricalvolatility(retrieveDataForTokensFromDatabase2(token0symbol, token1symbol), 30)
				fmt.Print("vol: ")
				fmt.Print(volatility)

				if volatility == 0.0 {
					volatility = 0.276
				} 

				imp_loss_hist := estimate_impermanent_loss_hist(volatility, 1, "Balancer")
				px_return_hist := calculate_price_return_x_days(Histrecord, 30)

				if px_return_hist < -1 {
					px_return_hist = 0
				}

				ROI_raw_est := calculateROI_raw_est(currentInterestrate, float32(BalancerRewardPercentage), float32(future_pool_sz_est), float32(future_daily_volume_est), imp_loss_hist)      // + imp
				ROI_vol_adj_est := calculateROI_vol_adj(ROI_raw_est, volatility)                                                                                                               // Sharpe ratio
				ROI_hist := calculateROI_hist(currentInterestrate, float32(BalancerRewardPercentage), historical_pool_sz_avg, historical_pool_daily_volume_avg, imp_loss_hist, px_return_hist) // + imp + hist

				//fmt.Print("| ROI_raw_est: ")
				//fmt.Print(ROI_raw_est)
				//fmt.Print("| ROI_vol_adj_est: ")
				//fmt.Print(ROI_vol_adj_est)
				//fmt.Print("| ROI_hist: ")
				//fmt.Print(ROI_hist)
				/*
					fmt.Print("DECIMALS t0: ")
					fmt.Print(respBalancerById.Pool.Tokens[0].Symbol)
					fmt.Print(respBalancerById.Pool.Tokens[0].ID)
					fmt.Print(" | ")
					fmt.Print(respBalancerById.Pool.Tokens[0].Address)
					fmt.Print(" | ")
					fmt.Print(respBalancerById.Pool.Tokens[0].Decimals)
					fmt.Print(" | t1: ")
					fmt.Print(respBalancerById.Pool.Tokens[1].Symbol)
					fmt.Print(respBalancerById.Pool.Tokens[1].ID)
					fmt.Print(" | ")
					fmt.Print(respBalancerById.Pool.Tokens[1].Address)
					fmt.Print(" | ")
					fmt.Print(respBalancerById.Pool.Tokens[1].Decimals)
					fmt.Print(" | ")
				*/
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
	// if pool len is == 2
	fmt.Println("BALANCER COMPLETED!!!!!")

} // balancer get data close

func getTradingVolumeFromTxLog2(logs []*types.Log, pooltopic string, token0AD string, token1AD string, token1decimals int) int64 {
	// func always gets token1 amounts
	//token0AD := token0 // "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599" // BTC
	//token1AD := token1 // "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2" // WETH
	decimals := token1decimals
	varbytes := big.NewInt(0)
	token1_found := false

	if len(logs) <= 10 {
		for i := 0; i < len(logs); i++ {
			if len(logs[i].Topics) >= 3 {
				if logs[i].Topics[2] == common.HexToHash(token1AD) {
					token1_found = true
					if len(logs[i].Data) >= 32 {
						varbytes = new(big.Int).SetBytes(logs[i].Data[0:32])
						//fmt.Print(" | 0-32: ")
						//fmt.Print(varbytes)
					}
				}
			}
		}
		if !token1_found {
			for i := 0; i < len(logs); i++ {
				if len(logs[i].Topics) >= 3 {
					if logs[i].Topics[2] == common.HexToHash(token0AD) {
						if len(logs[i].Data) >= 64 {
							varbytes = new(big.Int).SetBytes(logs[i].Data[32:64])
							//fmt.Print(" | 0-32: ")
							//fmt.Print(varbytes)
						}
					}
				}
			}
		} // not found ether
		//end Short list
	} else {
		for i := 0; i < len(logs); i++ {
			if len(logs[i].Topics) >= 3 {
				if logs[i].Topics[2] == common.HexToHash(token1AD) {
					if len(logs[i].Data) >= 32 {
						varbytes = varbytes.Add(varbytes, new(big.Int).SetBytes(logs[i].Data[0:32]))
						// fmt.Print(varbytes)
					}
				}
			}
		}
	} // end long lists

	for i := 0; i < len(logs); i++ {
		//fmt.Print(" |  i:::: ")
		//fmt.Print(i)
		if len(logs[i].Data) >= 32 {
			//fmt.Print(" | 0-32: ")
			//fmt.Print(new(big.Int).SetBytes(logs[i].Data[0:32]))
		}
		if len(logs[i].Data) >= 64 {
			//fmt.Print(" | len: ")
			//fmt.Print(len(logs[i].Data))
			//fmt.Print(" | 32-64: ")
			//fmt.Print(new(big.Int).SetBytes(logs[i].Data[32:64]))
		}
		if len(logs[i].Topics) > 0 {
			//fmt.Print(" | t0: ")
			//fmt.Print(logs[i].Topics[0])
		}
		if len(logs[i].Topics) > 1 {
			//fmt.Print(" | t1: ")
			//fmt.Print(logs[i].Topics[1])
		}
		if len(logs[i].Topics) > 2 {
			//fmt.Print(" | t2: ")
			//fmt.Print(logs[i].Topics[2])
			//fmt.Print(" | ")
		}
		//fmt.Println(" ")
	}

	ten := big.NewInt(10)
	ten.Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	varbytes.Div(varbytes, ten)
	num := varbytes.Int64()
	return num
}

/*

//func estimate_future_balancer_volume_and_pool_sz(histvolume BalancerHistVolumeQuery) (float32, float32) {
	future_volume_est := 0.0
	future_sz_est := 0.0

	var count float64
	var count_sz float64
	count = 0
	count_sz = 0

	for i := 0; i < len(histvolume.Pool.Swaps); i++ {

			fmt.Print("ESTIMATING BALANCER FUTURE VOLUME + POOL SZ: ")
			fmt.Print("transaction volume: ")
			fmt.Print(histvolume.Pool.Swaps[i].TokenAmountIn)
			fmt.Print(" | pool sz (liquidity): ")
			fmt.Println(histvolume.Pool.Swaps[i].PoolLiquidity)

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

	return float32(future_volume_est), float32(future_sz_est) // USD
}

*/

func isTokenStableCoin(coinName string) bool {
	if coinName == "USDT" {
		return true
	} else if coinName == "USDC" {
		return true
	} else if coinName == "USD" {
		return true
	} else if coinName == "TUSD" {
		return true
	} else if coinName == "DAI" {
		return true
	} else if coinName == "GUSD" {
		return true
	} else if coinName == "BUSD" {
		return true
	} else {
		return false
	}
}
