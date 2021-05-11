package db

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	//	"time"

	"pusher/defi_aggregator/running-results-table/internal/db/curveRegistry"
	"pusher/defi_aggregator/running-results-table/internal/db/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func convCurveToken(token string) string {
	assetName := " "

	if token == "Eth" {
		assetName = "ETH"
	} else if token == "Republic Token" {
		assetName = "REN"
	} else if token == "Synthetix Network Token" {
		assetName = "SNX"
	} else if token == "yearn.finance" {
		assetName = "YFI"
	} else if token == "Wrapped BTC" {
		assetName = "WBTC"
	} else if token == "Wrapped Ether" {
		assetName = "WETH"
	} else if token == "Uniswap" {
		assetName = "UNI"
	}
	return assetName
}

/*
func estimate_future_curve_volume_and_pool_sz(dates []int64, tradingvolumes []int64, poolsizes []int64) (float32, float32) {

	return 0.0, 0.0
}
*/
func conv_curve_token_to_uniswap(curve_token string) string {
	assetName := curve_token

	if curve_token == "yDAI" {
		assetName = "DAI"
	} else if curve_token == "cDAI" {
		assetName = "DAI"
	} else if curve_token == "ycDAI" {
		assetName = "DAI"
	} else if curve_token == "oBTC" {
		assetName = "WBTC"
	} else if curve_token == "HBTC" {
		assetName = "WBTC"
	} else if curve_token == "oBTC" {
		assetName = "WBTC"
	} else if curve_token == "TBTC" {
		assetName = "WBTC"
	} else if curve_token == "BBTC" {
		assetName = "WBTC"
	} else if curve_token == "pBTC" {
		assetName = "WBTC"
	} else if curve_token == "oBTC" {
		assetName = "WBTC"
	} else if curve_token == "renBTC" {
		assetName = "WBTC"
	} else if curve_token == "GUSD" {
		assetName = "USDC"
	} else if curve_token == "HUSD" {
		assetName = "USDC"
	} else if curve_token == "mUSD" {
		assetName = "USDC"
	} else if curve_token == "sUSD" {
		assetName = "USDC"
	} else if curve_token == "cUSDC" {
		assetName = "USDC"
	} else if curve_token == "yUSDC" {
		assetName = "USDC"
	}

	return assetName
}

func getCurveData(database *Database, uniswapreqdata UniswapInputStruct) {
	fmt.Print("GETTING CURVE DATA!!!")
	// Connecting to client
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	var respUniswapTicker UniswapTickerQuery // Used to look up Uniswap IDs of 'ETH' etc
	var respUniswapHist UniswapHistQuery
	var Histrecord HistoricalCurrencyData
	// Creating a contract instance
	var curveRegistryAddress = common.HexToAddress("0x7D86446dDb609eD0F5f8684AcF30380a356b2B4c")
	provider, err := curveRegistry.NewMain(curveRegistryAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	pools_to_pull := int64(32)
	var pool_addresses []common.Address
	var i int64

	// Main loop
	for i = 0; i < pools_to_pull; i++ {
		pool_address, _ := provider.PoolList(&bind.CallOpts{}, big.NewInt(i))
		pool_addresses = append(pool_addresses, pool_address)

		// got list of pools

		// for each pool
		// 	get actual token names
		//  check if part of filter
		// 	check number of assets in pool
		//  	create tokenqueue from coinaddresses
		// 			for each coin address -- convert to TOKEN symbol -- from curve - to Uniswap
		// 			check if uniswap data is available
		// 				if not - download it + append to db
		//

		// 	check if db contains (Curve-Token0-Token1-Token2-Token3) 2-3-4 tokens
		// 	if >= 5 tokens --> skip
		// 	if data > 1 day old
		// 	calculate oldest block(oldest time available)

		// if data_is_old {
		//  query data
		//  sum it up by day
		//  append to db
		//  calculate roi
		//	}
		// loop ends

		//Getting token Adresses:
		coin_addresses, err := provider.GetCoins(&bind.CallOpts{}, pool_address)
		if err != nil {
			log.Fatal(err)
		}

		//	fmt.Print("Coin Addresses: ")
		//	fmt.Print(coin_addresses)
		var tokenqueue []string
		var decimals []int64
		var balances []int64

		var dates []int64
		var tradingvolumes []int64
		var fees []int64
		var poolsizes []int64
		var interest []float64

		//skips pool if token not in filter.
		skip_pool := false
		// get actual token names goes here (Getting the token name) //--------------------
		for j := 0; j < len(coin_addresses); j++ {
			fmt.Print("j: ")
			fmt.Print(j)
			fmt.Print(" | ")
			if coin_addresses[j] == common.HexToAddress("0x0000000000000000000000000000000000000000") {
				continue
			}
			if coin_addresses[j] == common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
				continue
			}

			instance, err := token.NewToken(coin_addresses[j], client)
			if err != nil {
				fmt.Println("failed getting instance")
				log.Fatal(err)
			}
			name, err := instance.Symbol(&bind.CallOpts{})
			if err != nil {
				fmt.Print("failed getting symbol for: ")
				fmt.Print(coin_addresses[j])
				log.Fatal(err)
			}
			name = conv_curve_token_to_uniswap(name)

			fmt.Print("calculated token name: ")
			fmt.Print(name)

			if !isCoinPartOfFilter(name) {
				// coin not part of our filter
				// skip this whole pool
				skip_pool = true
				fmt.Print("..not part of our filter!! breaking out..")
				break
			}
			tokenqueue = append(tokenqueue, name)
			fmt.Print("sz of tokenqueue is now: ")
			fmt.Print(len(tokenqueue))

			current_coin_balances, err := provider.GetBalances(&bind.CallOpts{}, pool_address)

			for j := 0; j < len(decimals); j++ {
				balance := new(big.Float).SetInt(current_coin_balances[j])
				fmt.Print("balance: ")
				fmt.Print(balance)
				balance = negPow(balance, decimals[j])
				b, _ := balance.Int64()
				balances = append(balances, b)
			}

			// uniswap check
			if !isHistDataAlreadyDownloadedDatabase(tokenqueue[j]) {
				fmt.Print("In Uniswap hist data check..")
				// Get Uniswap Ids of these tokens
				uniswapreqdata.reqUniswapIDFromTokenTicker.Var("ticker", tokenqueue[j])
				if err := uniswapreqdata.clientUniswap.Run(ctx, uniswapreqdata.reqUniswapIDFromTokenTicker, &respUniswapTicker); err != nil {
					log.Fatal(err)
				}
				// Download historical data for each token for which data is missing
				if len(respUniswapTicker.IDsforticker) >= 1 {
					// request data from uniswap using this queried ticker
					uniswapreqdata.reqUniswapHist.Var("tokenid", setUniswapQueryIDForToken(tokenqueue[j], respUniswapTicker.IDsforticker[0].ID))
					fmt.Print("Querying historical (in GETCURVE) data from UNISWAP for: ")
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

		} // j loop through coins

		fmt.Print("Loop thru coins done!!")

		for jj := 0; jj < len(tokenqueue); jj++ {
			fmt.Print(tokenqueue[jj])
			fmt.Print(" | ")
		}

		fmt.Print(" | number of tokenqueue items: ")
		fmt.Print(len(tokenqueue))

		// if the next pool flag is true, go on to the next pool
		if skip_pool {
			fmt.Print("Skipping pool!! Found tokens not in our pool")
			continue
		}

		if len(tokenqueue) > 1 && len(tokenqueue) < 5 {
			days_ago := 1
			fmt.Print("getting oldest record: ")
			oldest_available_record := time.Unix(get_newest_timestamp_from_db("Curve", tokenqueue), 0)
			fmt.Print(" OLDEST RECORD IS AT: ")
			fmt.Print(oldest_available_record)
			oldest_lookup_time := time.Now()
			data_is_old := false

			if (time.Since(oldest_available_record).Hours()) > 24 {
				data_is_old = true
				oldest_lookup_time = oldest_lookup_time.AddDate(0, 0, -days_ago)
			}
			fmt.Print("is data old: ")
			fmt.Print(data_is_old)

			if data_is_old { // download it
				// Getting the number of decimal spaces for underlying coins in the pool
				coin_decimals, err := provider.GetDecimals(&bind.CallOpts{}, pool_address)

				// copy decimals

				for jz := 0; jz < len(decimals); jz++ {
					decimals = append(decimals, coin_decimals[jz].Int64())
				}

				fmt.Print(len(coin_decimals))
				if err != nil {
					log.Fatal(err)
				}

				// Downloading actual data goes here
				// get balance of each of the [x] tokens
				// convert into normal values - init64

				//3)  Query between oldest and current block for Curve-specific addresses

				dates, tradingvolumes, poolsizes, fees, interest = curveGetPoolVolume(pool_address, client, balances, tokenqueue)
				for jjj := 0; jjj < len(dates); jjj++ {
					recordID := append_record_to_database("Curve", tokenqueue, dates[jjj], tradingvolumes[jjj], fees[jjj], poolsizes[jjj], interest[jjj])
					if len(recordID) == 0 {
						fmt.Print(recordID)
					}
				}

			} // if data is old

			if !data_is_old { // else: data is not old
				// also query fees?
				dates, tradingvolumes, poolsizes, fees, interest = retrieve_hist_pool_sizes_volumes_fees_ir("Curve", tokenqueue)
			}

			// ROI stuff goes here
			currentSize := 0.0
			currentVolume := 0.0

			future_daily_volume_est, future_pool_sz_est := estimate_future_balancer_volume_and_pool_sz(dates, tradingvolumes, poolsizes)
			historical_pool_sz_avg, historical_pool_daily_volume_avg := future_pool_sz_est, future_daily_volume_est
			currentInterestrate := float32(0.00) // POPULATE

			CurveRewardPercentage := 0.0 // strconv.ParseFloat(respBalancerById.Pool.SwapFee, 32)
			volatility := calculatehistoricalvolatility(retrieveDataForTokensFromDatabase2(tokenqueue[0], tokenqueue[1]), 30)

			imp_loss_hist := estimate_impermanent_loss_hist(volatility, 1, "Curve")
			px_return_hist := calculate_price_return_x_days(Histrecord, 30)

			ROI_raw_est := calculateROI_raw_est(currentInterestrate, float32(CurveRewardPercentage), float32(future_pool_sz_est), float32(future_daily_volume_est), imp_loss_hist)      // + imp
			ROI_vol_adj_est := calculateROI_vol_adj(ROI_raw_est, volatility)                                                                                                            // Sharpe ratio
			ROI_hist := calculateROI_hist(currentInterestrate, float32(CurveRewardPercentage), historical_pool_sz_avg, historical_pool_daily_volume_avg, imp_loss_hist, px_return_hist) // + imp + hist
			/*
				fmt.Print("| ROI_raw_est: ")
				fmt.Print(ROI_raw_est)
				fmt.Print("| ROI_vol_adj_est: ")
				fmt.Print(ROI_vol_adj_est)
				fmt.Print("| ROI_hist: ")
				fmt.Print(ROI_hist)

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

			str := tokenqueue[0] + "/" + tokenqueue[1]
			if len(tokenqueue) == 3 {
				str = str + "/" + tokenqueue[2]
			}
			if len(tokenqueue) == 4 {
				str = str + "/" + tokenqueue[3]
			}

			// CHECK IF NOT DUPLICATING RECORD - IF ALREADY EXISTS - UPDATE NOT APPEND
			for k := 0; k < len(database.currencyinputdata); k++ {
				// Means record already exists - UPDATE IT, DO NOT APPEND
				if database.currencyinputdata[k].Pair == str && database.currencyinputdata[k].Pool == "Curve" {
					recordalreadyexists = true
					database.currencyinputdata[k].PoolSize = float32(currentSize)
					database.currencyinputdata[k].PoolVolume = float32(currentVolume)

					database.currencyinputdata[k].ROI_raw_est = ROI_raw_est
					database.currencyinputdata[k].ROI_vol_adj_est = ROI_vol_adj_est
					database.currencyinputdata[k].ROI_hist = ROI_hist

					database.currencyinputdata[k].Volatility = volatility
					database.currencyinputdata[k].Yield = currentInterestrate
				}
			} // k loop ends

			// APPEND IF NEW
			if !recordalreadyexists {
				database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{str, float32(currentSize),
					float32(currentVolume), currentInterestrate, "Curve", volatility, ROI_raw_est, 0.0, 0.0})
			}
			fmt.Println("APPENDED CURVE DATA ITERATION")

		} // if len > 1
	} // pools to pull loop ends
	fmt.Println("CURVE COMPLETED!!!!!")
} // Get Curve Data closes

/*
			// Debugging print
	for i = 0; i < int64(len(pool_addresses)); i++ {

		//	fmt.Print("pool address:")
		//	fmt.Print(pool_addresses[i])
		//	fmt.Print(" | Fees collected:")
		//		for j := 0; j < 2; j++ {
		//			fmt.Println(pools[i].fees[j])
		//		}
		//		fmt.Println("Addresses of coins in the pool:")
		//		fmt.Println(pools[i].assetAddresses)
		//		fmt.Println("Decimals for coins in the pool:")
		//		fmt.Println(pools[i].assetDecimals)
	}

	//	pools := getCurveDataI(client, provider, pools, 3, true)
	//	zero := big.NewFloat(0)
	/*
		for i := 0; i < 4; i++ {
			var normalsied_fees []*big.Float
			var normalsied_balances []*big.Float
			current_coin_balances, err := provider.GetBalances(&bind.CallOpts{}, pools[i].poolAddress)

			if err != nil {
				log.Fatal(err)
			}

			for j := 0; j < 8; j++ {
				balance := new(big.Float).SetInt(current_coin_balances[j])
				balance = negPow(balance, pools[i].assetDecimals[j].Int64())
				normalsied_balances = append(normalsied_balances, balance)
			}

			for j := 0; j < 8; j++ {

				fee := negPow(pools[i].fees[0][j], pools[i].assetDecimals[j].Int64())
				normalsied_fees = append(normalsied_fees, fee)
			}

			pools[i].normalsiedBalances = normalsied_balances

			fmt.Println("Returns:")
			for j := 0; j < 8; j++ {
				if pools[i].normalsiedBalances[j].Cmp(zero) > 0 {
					returns := new(big.Float).Quo(normalsied_fees[j], pools[i].normalsiedBalances[j])
					fmt.Println(returns)
				}

			}

		} // loop 4

}
}
*/
