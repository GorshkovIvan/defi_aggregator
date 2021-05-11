package db

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"math"
	"time"

	"pusher/defi_aggregator/running-results-table/internal/db/curveRegistry"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

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

	pools = getCurveDataI(client, provider, pools, 3, true)
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
		*/
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
		volumes, fees := curveGetPoolVolume(pool_address, oldest_block, latest_block, client)
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
		*/

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

			volumes, fees := curveGetPoolVolume(pool_address, oldest_block, latest_block, client)
			pools[count_pools].volumes = append(pools[count_pools].volumes, volumes)
			pools[count_pools].fees = append(pools[count_pools].fees, fees)
			count_pools++

		}

	}

	return pools
}

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

func curveGetPoolVolume(pool_address common.Address, oldest_block *big.Int, latest_block *big.Int, client *ethclient.Client) ([]int64, *[8]*big.Int, *[8]*big.Float, []int64) {

	poolTopics := []string{"0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140" /* "0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b"*/}

	//3)  Query between oldest and current block for Curve-specific addresses

	query := ethereum.FilterQuery{

		FromBlock: oldest_block,
		ToBlock:   latest_block, // = latest block
		Addresses: []common.Address{pool_address},
	}

	logsX, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//4)  Loop through received data and filter it again
	// For each transaction in logsX - check if it matches lookup criteria - add volume if does:
	var fees = new([8]*big.Float)
	var swap_volumes = new([8]*big.Int)
	var tradingvolumes []int64
	var dates []int64
	var poolsizes []int64

	for i := range swap_volumes {
		swap_volumes[i] = big.NewInt(0)
	}

	for i := range fees {
		fees[i] = big.NewFloat(0.0)
	}

	for i := 0; i < len(logsX); i++ {

		if logsX[i].Topics[0] != common.HexToHash(poolTopics[0]) /*&& logsX[i].Topics[0] != common.HexToHash(poolTopics[1])*/ {
			continue
		}

		txlog, err := client.TransactionReceipt(context.Background(), logsX[i].TxHash)

		if err != nil {
			log.Fatal(err)
		}

		// Here we have to add summing them up by day - not just total

		t_prev := uint64(0)
		t_new := uint64(0)

		t_prev = t_new       // uint
		t_new = block.Time() // uint
		
		cumulative_for_day := int64(0)

		if t_prev == 0 || (t_new-uint64(math.Mod(float64(t_new), 86400)))/86400 != 
		(t_prev-uint64(math.Mod(float64(t_prev), 86400)))/86400 { // 1 day
			dates = append(dates, int64(BoD(time.Unix(int64(t_new), 0)).Unix()))
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
		// add to volume
		asset0_index, asset0_volume, asset1_index, asset1_volume := getTradingVolumeFromTxLogCurve(txlog.Logs, poolTopics)
		swap_volumes[asset0_index].Add(swap_volumes[asset0_index], asset0_volume)
		swap_volumes[asset1_index].Add(swap_volumes[asset1_index], asset1_volume)
		volume_float := new(big.Float).SetInt(asset1_volume)
		fees[asset1_index].Add(fees[asset1_index], volume_float.Mul(volume_float, big.NewFloat(0.02)))

	}

	return dates, swap_volumes, fees, poolsizes
}

func decodeBytesCurve(log *types.Log) (int, *big.Int, int, *big.Int) {

	asset0_index, _ := strconv.Atoi((new(big.Int).SetBytes(log.Data[0:32])).String())
	asset0_volume := new(big.Int).SetBytes(log.Data[32:64])

	asset1_index, _ := strconv.Atoi((new(big.Int).SetBytes(log.Data[64:96])).String())
	asset1_volume := new(big.Int).SetBytes(log.Data[96:128])

	return asset0_index, asset0_volume, asset1_index, asset1_volume
}

func getTradingVolumeFromTxLogCurve(logs []*types.Log, pooltopics []string) (int, *big.Int, int, *big.Int) {

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
		return -1, common.Big0, -1, common.Big0
	}
	asset0_index, asset0_volume, asset1_index, asset1_volume := decodeBytesCurve(firstLog)

	return asset0_index, asset0_volume, asset1_index, asset1_volume
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
