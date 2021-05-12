package db

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"encoding/hex"
	aaveDataProvider "pusher/defi_aggregator/running-results-table/internal/db/aave_protocol_data_provider"
	"pusher/defi_aggregator/running-results-table/internal/db/token"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AavePoolData struct {
	assetAddress     string
	assetName        string
	interest_rates   []*big.Int
	volumes          []*big.Int
	currentBalance   *big.Int
	rate_types       []int
	flashLoanVolumes []*big.Int
	flashLoanFees    []*big.Int
	fees             []*big.Float
	decimals         int64
	timestamp        []int64
}

type AaveDailyData struct {
	assetAddress            string
	assetName               string
	weightedAverageInterest *big.Float

	volumes []*big.Int
}

func getAave2Data(database *Database, uniswapreqdata UniswapInputStruct) {

	var token_check []string
	token_check = append(token_check, "DAI")
	oldest_available_record := time.Unix(get_newest_timestamp_from_db("Aave2", token_check), 0)

	fmt.Print("XXX: ")
	fmt.Print(time.Since(oldest_available_record).Hours())

	data_is_old := true

	if time.Since(oldest_available_record).Hours() < 35 {
		data_is_old = false
	}

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	aave_pool_address := common.HexToAddress("0x057835Ad21a177dbdd3090bB1CAE03EaCF78Fc6d")

	aave2_data_provider, err := aaveDataProvider.NewStore(aave_pool_address, client)

	if err != nil {
		log.Fatal(err)
	}

	var aave_daily_data []AavePoolData

	days_needed := 2

	for i := days_needed; i > 1 && data_is_old; i-- {

		fmt.Print("Day: ")
		fmt.Println(i)
		aave_daily_data = getAave2DataDaily(client, aave_daily_data, i, aave2_data_provider)

	}


	// Getting current balances for aave2
	fmt.Print("len(aave_daily_data): ")
	fmt.Print(len(aave_daily_data))

	for i := 0; i < len(aave_daily_data); i++ {

		pool_address := common.HexToAddress(aave_daily_data[i].assetAddress)
		reserveData, err := aave2_data_provider.GetReserveData(&bind.CallOpts{}, pool_address)

		fmt.Println(aave_daily_data[i].assetName)

		if err != nil {
			log.Fatal(err)
		}

		returnsSum := big.NewFloat(0)

		for j := 29; j > 29-(days_needed-1); j-- {
			fmt.Println("Day: ")
			fmt.Print(j)

			sum := big.NewInt(0)
			availableLiquidity := reserveData.AvailableLiquidity
			totalStableDebt := reserveData.TotalStableDebt
			totalVariableDebt := reserveData.TotalVariableDebt
			sum.Add(sum, availableLiquidity)
			sum.Add(sum, totalStableDebt)
			sum.Add(sum, totalVariableDebt)
			aave_daily_data[i].currentBalance = sum

			zero := Zero()
			fmt.Print("COMP ====== ")
			cmp := 0
			if len(aave_daily_data[i].fees) >= i { // ok
				cmp = aave_daily_data[i].fees[j].Cmp(zero)
			}
			fmt.Print(cmp)

			if cmp == 0 {
				

				var token_ []string
				token_ = append(token_, aave_daily_data[i].assetName)
				timestamp := aave_daily_data[i].timestamp[j]
				fmt.Print("Appending TOKEN TO AAVE: ")
				fmt.Print(token_[0])
				append_record_to_database("Aave2", token_, timestamp, int64(0), int64(0), int64(0), float64(0), float64(0))

				zeroInt := big.NewInt(0)
				zeroFloat := new(big.Float).SetInt(zeroInt)
				returnsSum.Add(returnsSum, zeroFloat)

				continue
			}

			balanceFloat := negPow(new(big.Float).SetInt(aave_daily_data[i].currentBalance), aave_daily_data[i].decimals)
			volumesFloat := negPow(new(big.Float).SetInt(aave_daily_data[i].volumes[j]), aave_daily_data[i].decimals)
			weightedAverageInterest := Div(aave_daily_data[i].fees[j], volumesFloat)
			totalBorrowed := big.NewInt(0)
			totalBorrowed.Add(totalBorrowed, totalStableDebt)
			totalBorrowed.Add(totalBorrowed, totalVariableDebt)
			totalBorrowedFloat := negPow(new(big.Float).SetInt(totalBorrowed), aave_daily_data[i].decimals)
			utilisationRate := Div(totalBorrowedFloat, balanceFloat)
			returns := Mul(utilisationRate, weightedAverageInterest)
			returnsSum.Add(returnsSum, returns)
			timestamp := aave_daily_data[i].timestamp[j]
			fmt.Print("Timestamp: ")
			fmt.Println(timestamp)
			fmt.Print("Returns: ")
			fmt.Println(returns)

			var token_ []string

			token_ = append(token_, aave_daily_data[i].assetName)
			w_int, _ := weightedAverageInterest.Float64()
			uti, _ := utilisationRate.Float64()

			fmt.Print("Appending TOKEN TO AAVE: ")
			fmt.Print(token_[0])

			append_record_to_database("Aave2", token_, timestamp, int64(0), int64(0), int64(0), w_int, uti)
		} // daily record -- finished 1 pool
		// roi stuff goes here

		currentInterestrate := float32(0)
		if !data_is_old { // else: data is not old
			var tokens []string
			tokens = append(tokens, aave_daily_data[i].assetName)
			
			_, _, _, _, interest, utilization := retrieve_hist_pool_sizes_volumes_fees_ir("Aave2", tokens)

			for i := 0; i < len(interest); i++ {
				currentInterestrate += float32(interest[i]) * float32(utilization[i])
			}

		} else {

			returnsSumFloat32, _ := returnsSum.Float32()
			currentInterestrate = returnsSumFloat32 / float32(days_needed) // XXX Interest rate w avg * utilization - AVG L30D
		}

		future_daily_volume_est, future_pool_sz_est := 0, 1
		historical_pool_sz_avg, historical_pool_daily_volume_avg := future_pool_sz_est, future_daily_volume_est
		AaveRewardPercentage := 0.0 
		// if token data in database - get actual volatility
		// if not available - set volatility to 0

		volatility := float32(0.0)
		px_return_hist := float32(0.0)
		Histrecord := retrieveDataForTokensFromDatabase2(aave_daily_data[i].assetName, "USD") // returns blank object if no hist record
		fmt.Print("aave_daily_data[i].assetName: ")
		fmt.Print(aave_daily_data[i].assetName)
		fmt.Print(len(Histrecord.Date))
		if len(Histrecord.Date) > 0 {
			volatility = calculatehistoricalvolatility(retrieveDataForTokensFromDatabase2(aave_daily_data[i].assetName, "USD"), 30)
			px_return_hist = calculate_price_return_x_days(Histrecord, 30)
		}

		imp_loss_hist := 0.0

		ROI_raw_est := calculateROI_raw_est(currentInterestrate, float32(AaveRewardPercentage), float32(future_pool_sz_est), float32(future_daily_volume_est), float32(imp_loss_hist))                        // + imp
		ROI_vol_adj_est := calculateROI_vol_adj(ROI_raw_est, volatility)                                                                                                                                      // Sharpe ratio
		ROI_hist := calculateROI_hist(currentInterestrate, float32(AaveRewardPercentage), float32(historical_pool_sz_avg), float32(historical_pool_daily_volume_avg), float32(imp_loss_hist), px_return_hist) // + imp + hist

		fmt.Print("| ROI_raw_est: ")
		fmt.Print(ROI_raw_est)
		fmt.Print("| ROI_vol_adj_est: ")
		fmt.Print(ROI_vol_adj_est)
		fmt.Print("| ROI_hist: ")
		fmt.Print(ROI_hist)

		var recordalreadyexists bool
		recordalreadyexists = false
		token0symbol := aave_daily_data[i].assetName
		token1symbol := aave_daily_data[i].assetName
		// CHECK IF NOT DUPLICATING RECORD - IF ALREADY EXISTS - UPDATE NOT APPEND
		for k := 0; k < len(database.currencyinputdata); k++ {
			// Means record already exists - UPDATE IT, DO NOT APPEND
			if database.currencyinputdata[k].Pair == token0symbol+"/"+token1symbol && database.currencyinputdata[k].Pool == "Aave2" {
				recordalreadyexists = true
				database.currencyinputdata[k].PoolSize = float32(0.0)
				database.currencyinputdata[k].PoolVolume = float32(0.0)

				database.currencyinputdata[k].ROI_raw_est = ROI_raw_est
				database.currencyinputdata[k].ROI_vol_adj_est = ROI_vol_adj_est
				database.currencyinputdata[k].ROI_hist = ROI_hist

				database.currencyinputdata[k].Volatility = volatility
				database.currencyinputdata[k].Yield = currentInterestrate
			}
		}

		// APPEND IF NEW
		if !recordalreadyexists {
			database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{token0symbol + "/" + token1symbol, float32(0.0),
				float32(0.0), currentInterestrate, "Aave2", volatility, ROI_raw_est, ROI_vol_adj_est, ROI_hist})

		}

	} // done all pools

}

func sumVolumes(volumes []*big.Int) *big.Int {

	sum := big.NewInt(0)
	for i := 0; i < len(volumes); i++ {
		sum.Add(sum, volumes[i])
	}

	return sum
}

func sumFees(fees []*big.Float) *big.Float {

	sum := big.NewFloat(0)
	for i := 0; i < len(fees); i++ {
		sum.Add(sum, fees[i])
	}

	return sum
}

func getAave2DataDaily(client *ethclient.Client, aave_daily_data []AavePoolData, daysAgo int, data_provider *aaveDataProvider.Store) []AavePoolData {
	fmt.Println("getAave2DataDaily")
	oldest_block, latest_block := getOldestBlock(client, daysAgo)

	old_block, err := client.BlockByNumber(context.Background(), oldest_block)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Oldest block: ")
	fmt.Println(old_block.Time())

	latest_time := int64(BoD(time.Unix(int64(old_block.Time()), 0)).Unix())

	pool_address := common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")

	volumes_data := aaveGetPoolVolume(pool_address, oldest_block, latest_block, client, data_provider)

	for i := 0; i < len(volumes_data); i++ {
		pool_added := false
		for j := 0; j < len(aave_daily_data); j++ {
			if volumes_data[i].assetName == aave_daily_data[j].assetName {
				pool_added = true
			}
		}
		if !pool_added {
			newPool := AavePoolData{assetName: volumes_data[i].assetName, assetAddress: volumes_data[i].assetAddress, decimals: volumes_data[i].decimals}
			for j := 0; j < 31; j++ {
				newPool.volumes = append(newPool.volumes, big.NewInt(0))
				newPool.fees = append(newPool.fees, big.NewFloat(0.0))
				newPool.timestamp = append(newPool.timestamp, int64(0))
			}
			aave_daily_data = append(aave_daily_data, newPool)
		}
	}

	for i := 0; i < len(aave_daily_data); i++ {

		for j := 0; j < len(volumes_data); j++ {
			if volumes_data[j].assetName == aave_daily_data[i].assetName {

				total_volumes := sumVolumes(volumes_data[j].volumes)
				aave_daily_data[i].volumes[31-daysAgo] = total_volumes
				total_fees := sumFees(volumes_data[j].fees)
				aave_daily_data[i].fees[31-daysAgo] = total_fees

			}
			aave_daily_data[i].timestamp[31-daysAgo] = int64(latest_time)
		}

	}
	//volumes_data = aaveGetFlashLoansVolume(pool_address, oldest_block, client, volumes_data)

	return aave_daily_data

}

func BoD(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func getOldestBlock(client *ethclient.Client, daysAgo int) (*big.Int, *big.Int) {

	fmt.Println("Getting oldest block")
	var current_block *big.Int
	var oldest_block *big.Int
	var latest_block *big.Int
	current_block = big.NewInt(0)

	// Get current block
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	current_block = header.Number

	//2)  Find oldest block in our lookup date range
	oldest_block = new(big.Int).Set(current_block)
	latest_block = new(big.Int).Set(current_block)

	now := time.Now()

	time_needed := uint64(now.Unix()) - 24*60*60*uint64(daysAgo)
	time_for_latest_block := uint64(now.Unix()) - 24*60*60*uint64(daysAgo-1)

	var j int64
	j = 0
	latest_block_found := false
	for {
		j -= 50
		oldest_block.Add(oldest_block, big.NewInt(j))

		if !latest_block_found {
			latest_block.Add(latest_block, big.NewInt(j))
		}

		block, err := client.BlockByNumber(context.Background(), oldest_block)
		if err != nil {
			log.Fatal(err)
		}

		if block.Time() < time_for_latest_block {
			latest_block_found = true
		}

		if block.Time() < time_needed {

			break
		}
	}
	fmt.Println("Got to the end of the oldest block")
	return oldest_block, latest_block
}

func aaveGetPoolVolume(pool_address common.Address, oldest_block *big.Int, latest_block *big.Int, client *ethclient.Client, aaveDataProvider *aaveDataProvider.Store) []AavePoolData {
	fmt.Println("aaveGetPoolVolume")
	var pools []AavePoolData

	poolTopics := []string{"0xc6a898309e823ee50bac64e45ca8adba6690e99e7841c45d754e2a38e9019d9b"}

	query := ethereum.FilterQuery{

		FromBlock: oldest_block,
		ToBlock:   latest_block, // = latest block
		Addresses: []common.Address{pool_address},
	}

	logsX, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	allocated := false

	for i := 0; i < len(logsX); i++ {

		if logsX[i].Topics[0] != common.HexToHash(poolTopics[0]) {
			continue
		}

		txlog, err := client.TransactionReceipt(context.Background(), logsX[i].TxHash)

		if err != nil {
			log.Fatal(err)
		}

		// add to volume
		amount, rate_type, interest_rate, assetAddress := getTradingVolumeFromTxLog(txlog.Logs, poolTopics)
		allocated = false

		for i := range pools {
			if pools[i].assetAddress == assetAddress {

				allocated = true
				
				pools[i].volumes = append(pools[i].volumes, amount)
				pools[i].interest_rates = append(pools[i].interest_rates, interest_rate)
				pools[i].rate_types = append(pools[i].rate_types, rate_type)
				pools[i].fees = append(pools[i].fees, calculateFee(amount, interest_rate, pools[i].decimals))

			}
		}
		var name string
		if !allocated {

			if assetAddress != "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee" {

				tokenAddress := common.HexToAddress(assetAddress)
				instance, err := token.NewToken(tokenAddress, client)
				if err != nil {
					log.Fatal(err)
				}
				name, err = instance.Symbol(&bind.CallOpts{})

				if err != nil {
					name = "Unknown"
					
				}
			} else {
				name = "Eth"
			}

			interest_rates := []*big.Int{interest_rate}
			volumes := []*big.Int{amount}
			rate_types := []int{rate_type}
			// Calculating decimals


			if err != nil {
				log.Fatal(err)
			}
			tokenAddress := common.HexToAddress(assetAddress)

			reserve_data, err := aaveDataProvider.GetReserveConfigurationData(&bind.CallOpts{}, tokenAddress)

			if err != nil {
				log.Fatal(err)
			}

			decimals := (reserve_data.Decimals).Int64()
			fee := calculateFee(amount, interest_rate, decimals)
			fees := []*big.Float{fee}
			pools = append(pools, AavePoolData{assetAddress: assetAddress, interest_rates: interest_rates, volumes: volumes, rate_types: rate_types,
				assetName: name, fees: fees, decimals: decimals})

		}

	}
	return pools
}

func decodeBytes(log *types.Log) (*big.Int, int, *big.Int) {

	amount := new(big.Int).SetBytes(log.Data[32:64])
	rate_type, _ := strconv.Atoi((new(big.Int).SetBytes(log.Data[64:96])).String())
	interest_rate := new(big.Int).SetBytes(log.Data[96:128])

	return amount, rate_type, interest_rate

}


func HashToReserveAddress(hash common.Hash) string {
	var value []string
	value = append(value, "0", "x")
	value = append(value, hex.EncodeToString(hash[12:32]))
	valueStr := strings.Join(value, "")

	return valueStr
}

func getTradingVolumeFromTxLog(logs []*types.Log, pooltopics []string) (*big.Int, int, *big.Int, string) {

	var firstLog *types.Log
	var assetAddress string
	

	for _, log := range logs {
		if log.Topics[0] != common.HexToHash(pooltopics[0]) {
			continue
		}
		if firstLog == nil {

			firstLog = log
			address := log.Topics[1]
			assetAddress = HashToReserveAddress(address)

		}
		
	}

	if firstLog == nil { // could not find any valid swaps, thus the transaction failed
		return common.Big0, -1, common.Big0, "none"
	}
	amount, rate_type, interest_rate := decodeBytes(firstLog)

	return amount, rate_type, interest_rate, assetAddress
}

/* Getting flash loans data */

func aaveGetFlashLoansVolume(pool_address common.Address, oldest_block *big.Int, client *ethclient.Client, pools []AavePoolData) []AavePoolData {

	poolTopics := []string{"0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55"}

	query := ethereum.FilterQuery{

		FromBlock: oldest_block,
		ToBlock:   nil, // = latest block
		Addresses: []common.Address{pool_address},
	}

	fmt.Println(query.FromBlock)
	fmt.Println(query.ToBlock)

	logsX, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(logsX))

	allocated := false

	for i := 0; i < len(logsX); i++ {

		if logsX[i].Topics[0] != common.HexToHash(poolTopics[0]) {
			continue
		}

		txlog, err := client.TransactionReceipt(context.Background(), logsX[i].TxHash)

		if err != nil {
			log.Fatal(err)
		}

		// add to volume
		amount, deposit_fee, assetAddress := getFlashLoansVolumeFromTxLog(txlog.Logs, poolTopics)
		allocated = false

		for i := range pools {
			if pools[i].assetAddress == assetAddress {

				allocated = true
				//fmt.Println("Appended %s", assetAddress)
				pools[i].flashLoanVolumes = append(pools[i].flashLoanVolumes, amount)
				pools[i].flashLoanFees = append(pools[i].flashLoanFees, deposit_fee)

			}
		}

		var name string
		if !allocated {
			if assetAddress != "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee" {
				fmt.Println("Bug")
				fmt.Println(assetAddress)
				tokenAddress := common.HexToAddress(assetAddress)
				instance, err := token.NewToken(tokenAddress, client)
				if err != nil {
					log.Fatal(err)
				}
				name, err = instance.Name(&bind.CallOpts{})
				if err != nil {
					log.Fatal(err)
				}
			} else {
				fmt.Println("assetAddress in else statement:")
				fmt.Println(assetAddress)
				name = "Eth"
			}

			flashLoanFees := []*big.Int{deposit_fee}
			flashLoanVolumes := []*big.Int{amount}

			pools = append(pools, AavePoolData{assetAddress: assetAddress, flashLoanFees: flashLoanFees, flashLoanVolumes: flashLoanVolumes,
				assetName: name})
			fmt.Println("pool added:")
			fmt.Println(assetAddress)
		}

	}

	return pools

}

func decodeFlashLoanBytes(log *types.Log) (*big.Int, *big.Int) {

	amount := new(big.Int).SetBytes(log.Data[0:32])
	total_fee := new(big.Int).SetBytes(log.Data[32:64])
	protocol_fee := new(big.Int).SetBytes(log.Data[64:96])
	deposit_fee := big.NewInt(0).Sub(total_fee, protocol_fee)

	return amount, deposit_fee
}

func getFlashLoansVolumeFromTxLog(logs []*types.Log, pooltopics []string) (*big.Int, *big.Int, string) {

	var firstLog *types.Log
	var assetAddress string
	

	for _, log := range logs {
		if log.Topics[0] != common.HexToHash(pooltopics[0]) {
			continue
		}
		if firstLog == nil {

			firstLog = log
			address := log.Topics[2]
			assetAddress = HashToReserveAddress(address)

		}
		
	}

	if firstLog == nil { // could not find any valid swaps, thus the transaction failed
		return common.Big0, common.Big0, "none"
	}
	amount, deposit_rate := decodeFlashLoanBytes(firstLog)

	return amount, deposit_rate, assetAddress
}

func calculateFee(volume *big.Int, interest *big.Int, decimal int64) *big.Float {

	volume_float := negPow(new(big.Float).SetInt(volume), decimal)
	volume_interest := negPow(new(big.Float).SetInt(interest), 27)

	return Mul(volume_float, volume_interest)
}

func negPow(a *big.Float, e int64) *big.Float {
	result := Zero().Copy(a)
	divTen := big.NewFloat(0.1)
	for i := int64(0); i < e; i++ {
		result = Mul(result, divTen)
	}
	return result
}

func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}


func Div(a, b *big.Float) *big.Float {
	return Zero().Quo(a, b)
}
