package db

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
func getCurveData0() {

	// Connecting to client
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	// Creaitng a contract instance
	var curveRegistryAddress = common.HexToAddress("0x7D86446dDb609eD0F5f8684AcF30380a356b2B4c")
	provider, err := curveRegistry.NewMain(curveRegistryAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	var pools []CurvePoolData

	//pools = getCurveDataI(client, provider, pools, 3, true)
	fmt.Print("chk998")
	fmt.Println("Got day one")

	for i := 0; i < 31; i++ {

		fmt.Println("pool address:")
		fmt.Println(pools[i].poolAddress)
		fmt.Println("Fees collected:")
		for j := 0; j < 2; j++ {
			fmt.Println(pools[i].fees[j])
		}
		fmt.Println("Addresses of coins in the pool:")
		fmt.Println(pools[i].assetAddresses)
		fmt.Println("Decimals for coins in the pool:")
		fmt.Println(pools[i].assetDecimals)
		/*
			fmt.Println("Normalised volumes:")

			for j := 0; j < 8; j++{

				normalisedVolume := new(big.Float).SetInt(pools[i].volumes[j])
				normalisedVolume = negPow(normalisedVolume, pools[i].assetDecimals[j].Int64())

			}

	}

	zero := big.NewFloat(0)

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

	}
}
*/
/*
func getCurveDataI(client *ethclient.Client, provider *curveRegistry.Main, pools []CurvePoolData, daysAgo int, first_turn bool) []CurvePoolData {
	fmt.Print("In CurvedataI")
	number_of_pools := big.NewInt(32)
	oldest_block := getOldestBlock(client, daysAgo)
	latest_block := getOldestBlock(client, daysAgo-1)
	count_pools := 0
	var one = big.NewInt(1)
	start := big.NewInt(1)
	end := big.NewInt(0).Sub(number_of_pools, big.NewInt(1))

	if first_turn {
		// Getting data from pools		// Getting data for the first pool
		pool_address, err := provider.PoolList(&bind.CallOpts{}, big.NewInt(0))

		if err != nil {
			log.Fatal(err)
		}

		// Addresses of underlying coins in the pool
		coin_addresses, err := provider.GetCoins(&bind.CallOpts{}, pool_address)
		if err != nil {
			log.Fatal(err)
		}

		// Getting the number of decimal spaces for undelying coins in the pool
		coin_decimals, err := provider.GetDecimals(&bind.CallOpts{}, pool_address)
		if err != nil {
			log.Fatal(err)
		}

		// Getting current pool balances

		// Getting swap volumes and fees and balances
		volumes, fees := curveGetPoolVolume(pool_address, client)
		var volumes_array []*[8]*big.Int
		volumes_array = append(volumes_array, volumes)
		var fees_array []*[8]*big.Float
		fees_array = append(fees_array, fees)

		// Appending a list of pool data structs
		pools = append(pools, CurvePoolData{poolAddress: pool_address, assetAddresses: coin_addresses,
			volumes: volumes_array, fees: fees_array, assetDecimals: coin_decimals})

		count_pools++
		/*
			fmt.Println("pool address:")
			fmt.Println(pool_address)
			fmt.Println("Fees collected:")
			fmt.Println(pools[count_pools].fees)
			fmt.Println("Addresses of coins in the pool:")
			fmt.Println(pools[count_pools].assetAddresses)
			fmt.Println("Decimals for coins in the pool:")
			fmt.Println(pools[count_pools].assetDecimals)


		// Getting data for the rest of the pools

		// i must be a new int so that it does not overwrite start
		for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, one) {

			pool_address, err = provider.PoolList(&bind.CallOpts{}, i)
			fmt.Println(pool_address)
			if err != nil {
				log.Fatal(err)
			}

			coin_addresses, err := provider.GetCoins(&bind.CallOpts{}, pool_address)
			if err != nil {
				log.Fatal(err)
			}

			// Get decimals for underlying tokens

			coin_decimals, err := provider.GetDecimals(&bind.CallOpts{}, pool_address)

			if err != nil {
				log.Fatal(err)
			}

			// Getting volumes and fees

			volumes, fees := curveGetPoolVolume(pool_address, oldest_block, latest_block, client)

			var volumes_array []*[8]*big.Int
			volumes_array = append(volumes_array, volumes)
			var fees_array []*[8]*big.Float
			fees_array = append(fees_array, fees)

			pools = append(pools, CurvePoolData{poolAddress: pool_address, assetAddresses: coin_addresses,
				volumes: volumes_array, fees: fees_array, assetDecimals: coin_decimals})

			count_pools++

		}

	} else {

		fmt.Println("Got into else")

		pool_address, err := provider.PoolList(&bind.CallOpts{}, big.NewInt(0))

		if err != nil {
			log.Fatal(err)
		}

		volumes, fees := curveGetPoolVolume(pool_address, oldest_block, latest_block, client)
		fmt.Println("Got first fees")
		pools[count_pools].volumes = append(pools[count_pools].volumes, volumes)
		pools[count_pools].fees = append(pools[count_pools].fees, fees)
		count_pools++
		fmt.Println("Upended first pool:")
		fmt.Println(pools[count_pools].volumes)

		for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, one) {
			fmt.Println("looping:")
			fmt.Println(i)
			pool_address, err = provider.PoolList(&bind.CallOpts{}, i)
			if err != nil {
				log.Fatal(err)
			}

			volumes, fees := curveGetPoolVolume(pool_address, client)
			pools[count_pools].volumes = append(pools[count_pools].volumes, volumes)
			pools[count_pools].fees = append(pools[count_pools].fees, fees)
			count_pools++

		}

	}

	return pools
}
*/
/*
func getOldestBlock(client *ethclient.Client, daysAgo int) *big.Int {

	var current_block *big.Int
	var oldest_block *big.Int
	current_block = big.NewInt(0)

	// Get current block
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	current_block = header.Number

	//2)  Find oldest block in our lookup date range
	oldest_block = new(big.Int).Set(current_block)

	now := time.Now()

	//timeonehourago := uint64(now.Add(-2*time.Hour).Unix())
	//timeonemonthago := uint64((now.AddDate(0, 0, -1)).Unix())
	timeonemonthago := uint64(now.Unix()) - 24*60*60*uint64(daysAgo)

	var j int64
	j = 0

	for {
		j -= 500
		oldest_block.Add(oldest_block, big.NewInt(j))

		block, err := client.BlockByNumber(context.Background(), oldest_block)
		if err != nil {
			log.Fatal(err)
		}

		if block.Time() < timeonemonthago {

			break
		}
	}

	return oldest_block
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

			sz_0 := int64(float64(asset0_volume) * exch0)
			sz_1 := int64(float64(asset1_volume) * exch1)

			pool_fee := 0.02
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

//Function to convert a big Float to a BigInt
/*
func FloatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	// Set precision if required.
	// bigval.SetPrec(64)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}
*/

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

/*

/*
	current_coin_balances, err := provider.GetUnderlyingBalances(pool_address)

	if err != nil {
		log.Fatal(err)
	}

	// Add decimal spaces to volumes and fees and balances

	for i := 0; i < 8; i++{
		normvolumes[i] = negPow(volumes[i], coin_decimals.Int64())
		normfees[i] = negPow(fees[i], coin_decimals.Int64())
		current_coin_balances[i] = negPow(current_coin_balances[i], coin_decimals.Int64())
	}

	// Calculate returns

	var returns []*big.Float

	for i := 0; i < 8; i++{
		returns = append(returns, Quo(normfees[i], current_coin_balances[i]) )
	}
*/

/*
	/*
	fmt.Println("Returns:")
	fmt.Println(pools[count_pools].returns)

	for i := 0; i < 8; i++{

		normalisedVolume := new(big.Float).SetInt(pools[count_pools].volumes[i])

		normalisedVolume = negPow(normalisedVolume, pools[count_pools].assetDecimals[i].Int64())

		fmt.Println("Orginal volume")
		fmt.Println(pools[count_pools].volumes[i])
		fmt.Println("Normalised volume")
		fmt.Println(normalisedVolume)



	}

*/
