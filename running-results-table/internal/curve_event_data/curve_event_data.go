package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	curveRegistry "./curveRegistry"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CurvePoolData struct {

	poolAddress common.Address 
	poolCurrentBalances [8]*big.Int
	assetAddresses [8]common.Address
	assetDecimals [8]*big.Int
	assetIndices []*big.Int
	assetBalances [8]*big.Int
	volumes *[8]*big.Int 
	fees *[8]*big.Float
	returns []*big.Float

	

}

func main() {

	// Connecting to client 
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	//client, err := ethclient.Dial("http://localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	// Creaitng a contract instance 
	var curveRegistryAddress = common.HexToAddress("0x7D86446dDb609eD0F5f8684AcF30380a356b2B4c")
	provider, err := curveRegistry.NewMain(curveRegistryAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	// Gettign the number of pools
	number_of_pools, err := provider.PoolCount(&bind.CallOpts{})
	fmt.Println(number_of_pools)

	var one = big.NewInt(1)
	start := big.NewInt(1)
	end := big.NewInt(0).Sub(number_of_pools, big.NewInt(1))
	oldest_block := getOldestBlock(client)

	// Getting data from pools 

	var pools []CurvePoolData
	count_pools := 0

	// Getting data for the first pool
	pool_address, err := provider.PoolList(&bind.CallOpts{}, big.NewInt(0))

	if err != nil {
		log.Fatal(err)
	}


	//fmt.Println(pool_address)

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
	volumes, fees := curveGetPoolVolume(pool_address, oldest_block, client)
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

	// Appending a list of pool data structs
	pools = append(pools, CurvePoolData{poolAddress: pool_address, assetAddresses: coin_addresses, 
					volumes: volumes, fees:fees, assetDecimals: coin_decimals })

	fmt.Println("pool address:")				
	fmt.Println(pool_address)
	fmt.Println("Fees collected:")
	fmt.Println(pools[count_pools].fees)
	fmt.Println("Addresses of coins in the pool:")
	fmt.Println(pools[count_pools].assetAddresses)
	fmt.Println("Decimals for coins in the pool:")
	fmt.Println(pools[count_pools].assetDecimals)
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

		// Get indicies for underlying coins


		// Getting volumes and fees

		volumes, fees := curveGetPoolVolume(pool_address, oldest_block, client)

		pools = append(pools, CurvePoolData{poolAddress: pool_address, assetAddresses: coin_addresses, 
			volumes: volumes, fees: fees, assetDecimals: coin_decimals})

		count_pools++
		fmt.Println("pool address:")				
		fmt.Println(pool_address)
		fmt.Println("Fees collected:")
		fmt.Println(pools[count_pools].fees)
		fmt.Println("Addresses of coins in the pool:")
		fmt.Println(pools[count_pools].assetAddresses)
		fmt.Println("Decimals for coins in the pool:")
		fmt.Println(pools[count_pools].assetDecimals)
		fmt.Println("Normalised volumes:")
		
		for i := 0; i < 8; i++{
	
			normalisedVolume := new(big.Float).SetInt(pools[count_pools].volumes[i])
			
			normalisedVolume = negPow(normalisedVolume, pools[count_pools].assetDecimals[i].Int64())

			fmt.Println("Orginal volume")
			fmt.Println(pools[count_pools].volumes[i])
			fmt.Println("Normalised volume")
			fmt.Println(normalisedVolume)
			
		}

	}

}

func getOldestBlock(client *ethclient.Client) *big.Int {

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
	timeonemonthago := uint64((now.AddDate(0, 0, -1)).Unix())

	var j int64
	j = 0

	for {
		j -= 10
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

func curveGetPoolVolume(pool_address common.Address, oldest_block *big.Int, client *ethclient.Client) (*[8]*big.Int, *[8]*big.Float) {

	poolTopics := []string{"0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140"/* "0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b"*/}

	//3)  Query between oldest and current block for Balancer-specific addresses

	query := ethereum.FilterQuery{

		FromBlock: oldest_block,
		ToBlock:   nil, // = latest block
		Addresses: []common.Address{pool_address},
	}

	//fmt.Println("Querying FilterLogs..")

	logsX, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//4)  Loop through received data and filter it again
	// For each transaction in logsX - check if it matches lookup criteria - add volume if does:
	var fees = new([8]*big.Float)
	var swap_volumes = new([8]*big.Int)

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

		// add to volume
		asset0_index, asset0_volume, asset1_index, asset1_volume := getTradingVolumeFromTxLog(txlog.Logs, poolTopics)
		swap_volumes[asset0_index].Add(swap_volumes[asset0_index], asset0_volume)
		swap_volumes[asset1_index].Add(swap_volumes[asset1_index], asset1_volume)
		volume_float := new(big.Float).SetInt(asset1_volume)
		fees[asset1_index].Add(fees[asset1_index], volume_float.Mul(volume_float, big.NewFloat(0.02)))

	}

	return swap_volumes, fees

}

func decodeBytes(log *types.Log) (int, *big.Int, int, *big.Int) {

	asset0_index, _ := strconv.Atoi((new(big.Int).SetBytes(log.Data[0:32])).String())
	asset0_volume := new(big.Int).SetBytes(log.Data[32:64])

	asset1_index, _ := strconv.Atoi((new(big.Int).SetBytes(log.Data[64:96])).String())
	asset1_volume := new(big.Int).SetBytes(log.Data[96:128])

	return asset0_index, asset0_volume, asset1_index, asset1_volume
}

func getTradingVolumeFromTxLog(logs []*types.Log, pooltopics []string) (int, *big.Int, int, *big.Int) {

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
	asset0_index, asset0_volume, asset1_index, asset1_volume := decodeBytes(firstLog)

	return asset0_index, asset0_volume, asset1_index, asset1_volume
}

func negPow(a *big.Float, e int64) *big.Float {
    result := Zero().Copy(a)
	divTen := big.NewFloat(0.1)
    for i:=int64(0); i<e-1; i++ {
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