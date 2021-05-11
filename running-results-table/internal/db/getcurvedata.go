package db

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"time"

	"pusher/defi_aggregator/running-results-table/internal/db/curveRegistry"
	"pusher/defi_aggregator/running-results-table/internal/db/token"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CurvePoolData struct {
	poolAddress common.Address
	//poolCurrentBalances [8]*big.Int
	assetAddresses     [8]common.Address
	assetDecimals      [8]*big.Int
	volumes            []*[8]*big.Int
	fees               []*[8]*big.Float
	balances           []*[8]*big.Int
	normalisedBalances []*big.Float
}

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
	fmt.Println("GETTING CURVE DATA!!!")
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

		fmt.Print("Pool Address: ")
		fmt.Println(pool_address)

		//Getting token Adresses:
		coin_addresses, err := provider.GetCoins(&bind.CallOpts{}, pool_address)
		if err != nil {
			log.Fatal(err)
		}

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
				fmt.Println(coin_addresses[j])
				log.Fatal(err)
			}
			name = conv_curve_token_to_uniswap(name)

			if !isCoinPartOfFilter(name) {
				// coin not part of our filter - skip this whole pool
				skip_pool = true
				fmt.Print("..not part of filter..breaking out of loop..")
				break
			}
			tokenqueue = append(tokenqueue, name)

			current_coin_balances, err := provider.GetBalances(&bind.CallOpts{}, pool_address)

			for j := 0; j < len(decimals); j++ {
				balance := new(big.Float).SetInt(current_coin_balances[j])
				fmt.Print("i: ")
				fmt.Print(i)
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

		fmt.Print("Loop thru coins done!! ")
		fmt.Print("sz of tokenqueue: ")
		fmt.Print(len(tokenqueue))

		for jj := 0; jj < len(tokenqueue); jj++ {
			fmt.Print(tokenqueue[jj])
			fmt.Print(" | ")
		}

		fmt.Print(" | number of tokenqueue items: ")
		fmt.Println(len(tokenqueue))

		// if the next pool flag is true, go on to the next pool
		if skip_pool {
			fmt.Println("Skipping pool! Found tokens not in our permitted list")
			continue
		}

		if len(tokenqueue) > 1 && len(tokenqueue) < 5 {
			days_ago := 3
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
				dates, tradingvolumes, poolsizes, fees, interest = retrieve_hist_pool_sizes_volumes_fees_ir("Curve", tokenqueue)
			}

			// ROI calculation
			currentSize := 0.0
			if len(poolsizes) > 0 {
				currentSize = float64(poolsizes[len(poolsizes)-1])
			}

			currentVolume := 0.0
			if len(tradingvolumes) > 0 {
				currentVolume = float64(tradingvolumes[len(tradingvolumes)-1])
			}

			// Use balancer function - also applicable for curve
			future_daily_volume_est, future_pool_sz_est := estimate_future_balancer_volume_and_pool_sz(dates, tradingvolumes, poolsizes)
			historical_pool_sz_avg, historical_pool_daily_volume_avg := future_pool_sz_est, future_daily_volume_est
			currentInterestrate := float32(0.00) // POPULATE

			CurveRewardPercentage := 0.0
			volatility := calculatehistoricalvolatility(retrieveDataForTokensFromDatabase2(tokenqueue[0], tokenqueue[1]), 30)

			imp_loss_hist := estimate_impermanent_loss_hist(volatility, 1, "Curve")
			px_return_hist := calculate_price_return_x_days(Histrecord, 30)

			ROI_raw_est := calculateROI_raw_est(currentInterestrate, float32(CurveRewardPercentage), float32(future_pool_sz_est), float32(future_daily_volume_est), imp_loss_hist)      // + imp
			ROI_vol_adj_est := calculateROI_vol_adj(ROI_raw_est, volatility)                                                                                                            // Sharpe ratio
			ROI_hist := calculateROI_hist(currentInterestrate, float32(CurveRewardPercentage), historical_pool_sz_avg, historical_pool_daily_volume_avg, imp_loss_hist, px_return_hist) // + imp + hist

			fmt.Print("| ROI_raw_est: ")
			fmt.Print(ROI_raw_est)
			fmt.Print("| ROI_vol_adj_est: ")
			fmt.Print(ROI_vol_adj_est)
			fmt.Print("| ROI_hist: ")
			fmt.Print(ROI_hist)

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

func curveGetPoolVolume(pool_address common.Address, client *ethclient.Client, balances []int64, tokenqueue []string) ([]int64, []int64, []int64, []int64, []float64) {

	poolTopics := []string{"0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140" /* "0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b"*/}

	// block calcs - YA
	var current_block *big.Int
	var oldest_block *big.Int
	current_block = big.NewInt(0)

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	current_block = header.Number

	fmt.Print("current block: ")
	fmt.Print(current_block)

	oldest_block = new(big.Int).Set(current_block)
	j := int64(0) // compute block id [days_ago] days away from now
	for {
		j -= 2000
		oldest_block.Add(oldest_block, big.NewInt(j))

		block, err := client.BlockByNumber(context.Background(), oldest_block)
		if err != nil {
			log.Fatal(err)
		}

		oldest_lookup_time := time.Now() //.Unix()

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

	//3)  Query between oldest and current block for Curve-specific addresses

	query := ethereum.FilterQuery{
		FromBlock: oldest_block,
		ToBlock:   current_block, // = latest block
		Addresses: []common.Address{pool_address},
	}

	logsX, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//4)  Loop through received data and filter it again
	// For each transaction in logsX - check if it matches lookup criteria - add volume if does:
	var fees []int64
	var tradingvolumes []int64
	var dates []int64
	var poolsizes []int64
	var interest []float64
	cumulative_for_day_fees := int64(0)
	cumulative_for_day_volume := int64(0)
	t_prev := uint64(0)
	t_new := uint64(0)

	// which symbols do we need to get here??
	Histrecord_2 := retrieveDataForTokensFromDatabase2(tokenqueue[0], tokenqueue[1])

	// loop through whole log
	for i := 0; i < len(logsX); i++ {
		if logsX[i].Topics[0] != common.HexToHash(poolTopics[0]) /*&& logsX[i].Topics[0] != common.HexToHash(poolTopics[1])*/ {
			continue
		}

		// Get date from block number
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(logsX[i].BlockNumber)))
		if err != nil {
			log.Fatal(err)
		}

		txlog, err := client.TransactionReceipt(context.Background(), logsX[i].TxHash)

		if err != nil {
			log.Fatal(err)
		}

		// Here we have to add summing them up by day - not just total
		t_prev = t_new       // uint
		t_new = block.Time() // uint

		if t_prev == 0 || (t_new-uint64(math.Mod(float64(t_new), 86400)))/86400 !=
			(t_prev-uint64(math.Mod(float64(t_prev), 86400)))/86400 { // 1 day
			dates = append(dates, int64(BoD(time.Unix(int64(t_new), 0)).Unix()))
			fees = append(fees, cumulative_for_day_fees)
			tradingvolumes = append(tradingvolumes, cumulative_for_day_volume)

			// Get balances of each token
			//			bal_int, _ := bal_float.Int64()
			//			bal_intT2, _ := bal_floatT2.Int64()
			if len(Histrecord_2.Price) >= 1 {
				//				bal_int = bal_int                                                            //
				//				bal_intT2 = bal_intT2 * int64(Histrecord_2.Price[len(Histrecord_2.Price)-1]) //--------------------------------------------
			} // convert to token1

			var totalBal = 0
			for balLen := 0; balLen < len(balances); balLen++ {
				totalBal += int(balances[balLen])
			}

			poolsizes = append(poolsizes, int64(totalBal)) // bal.Int64()
			cumulative_for_day_fees = 0
			cumulative_for_day_volume = 0 // reset days tally if day threshold crossed
		} else { //-------------
			// convert to usd
			// Get the data
			asset0_index, asset0_volume, asset1_index, asset1_volume := getTradingVolumeFromTxLogCurve(txlog.Logs, poolTopics)

			fmt.Print(asset0_index)
			fmt.Print(asset1_index)

			// which tickers?
			exch0 := float64(1.0)
			exch1 := float64(1.0)

			// if isTokenStableCoin(xxx) {exch0 = float64(1.0)}
			// if isTokenStableCoin(xxx) {exch1 = float64(1.0)}

			sz_0 := int64(float64(asset0_volume) * exch0)
			sz_1 := int64(float64(asset1_volume) * exch1)

			pool_fee := 0.02
			// get actual fee

			f0 := int64(float64(sz_0) * pool_fee)
			f1 := int64(float64(sz_1) * pool_fee)

			// Add it to tally for that day
			cumulative_for_day_fees += (f0 + f1)
			cumulative_for_day_volume += (sz_0 + sz_1)
		}
		// add to volume

	} // loop through logs ends

	return dates, tradingvolumes, poolsizes, fees, interest
}

func decodeBytesCurve(log *types.Log) (int, *big.Int, int, *big.Int) {

	asset0_index, _ := strconv.Atoi((new(big.Int).SetBytes(log.Data[0:32])).String())
	asset0_volume := new(big.Int).SetBytes(log.Data[32:64])

	asset1_index, _ := strconv.Atoi((new(big.Int).SetBytes(log.Data[64:96])).String())
	asset1_volume := new(big.Int).SetBytes(log.Data[96:128])

	return asset0_index, asset0_volume, asset1_index, asset1_volume
}

func getTradingVolumeFromTxLogCurve(logs []*types.Log, pooltopics []string) (int64, int64, int64, int64) { //Should return an int not big int

	var firstLog *types.Log
	//var lastLog *types.Log

	for _, log := range logs {
		if log.Topics[0] != common.HexToHash(pooltopics[0]) /*&& log.Topics[0] != common.HexToHash(pooltopics[1])*/ {
			continue
		}
		if firstLog == nil {
			firstLog = log
		}
		//lastLog = log
	}

	if firstLog == nil { // could not find any valid swaps, thus the transaction failed
		return -1, 0, -1, 0
	}
	asset0_index, asset0_volume, asset1_index, asset1_volume := decodeBytesCurve(firstLog)

	return int64(asset0_index), asset0_volume.Int64(), int64(asset1_index), asset1_volume.Int64()
}

func negPow(a *big.Float, e int64) *big.Float {
	result := Zero().Copy(a)
	divTen := big.NewFloat(0.1)
	for i := int64(0); i < e-1; i++ {
		result = Mul(result, divTen)
	}
	return result
}

func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

func Mul(a, b *big.Float) *big.Float {
	return Zero().Mul(a, b)
}
