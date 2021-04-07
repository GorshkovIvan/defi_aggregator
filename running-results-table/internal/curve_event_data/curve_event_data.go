package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	curveRegistry "./curveRegistry"
)

func main(){

	client, err := ethclient.Dial("http://localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	

	// Getting addresses of all pools 

	// 1) Define pool specific parameters
	var curveRegistryAddress = common.HexToAddress("0x7D86446dDb609eD0F5f8684AcF30380a356b2B4c")
	provider, err := curveRegistry.NewMain(curveRegistryAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	number_of_pools, err := provider.PoolCount(&bind.CallOpts{})
	fmt.Println(number_of_pools)

	var one = big.NewInt(1)
	start := big.NewInt(1)
    end := big.NewInt(0).Sub(number_of_pools, big.NewInt(1))
	pool_address, err := provider.PoolList(&bind.CallOpts{}, big.NewInt(0))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pool_address)

    // i must be a new int so that it does not overwrite start
    for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, one) {

		pool_address, err = provider.PoolList(&bind.CallOpts{}, i)

		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Println(pool_address)
    }

	poolTopics := []string{"0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140", "0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b"}

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
	//fmt.Print("Current block: ") // 5671744
	//fmt.Println(current_block)
	now := time.Now()
	//fmt.Println(now)
	timeonemonthago := uint64(now.Add(-2*time.Hour).Unix())
	//fmt.Print("1m ago: ")
	//fmt.Println(timeonemonthago)
	var j int64
	j = 0
	// compute block id [30] days away from now
	for {
		j -= 10
		oldest_block.Add(oldest_block, big.NewInt(j))
		//	fmt.Print("oldest block: ")
		//	fmt.Println(oldest_block)
		//	fmt.Print("current block: ")
		//	fmt.Println(current_block)

		block, err := client.BlockByNumber(context.Background(), oldest_block)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(block.Time())

		if block.Time() < timeonemonthago {
			fmt.Print(" | Oldest block: ") // 5671744
			fmt.Println(oldest_block)
			fmt.Print(" | time: ")
			fmt.Print(block.Time()) // 1527211625
			fmt.Print("| Diff: ")
			diff := current_block.Sub(current_block, oldest_block)
			fmt.Println(diff)
			break
		}
	}

	//3)  Query between oldest and current block for Balancer-specific addresses
	test_address, err := provider.PoolList(&bind.CallOpts{}, big.NewInt(0))

	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		// BlockHash *common.Hash, - add more parameters to filter
		FromBlock: oldest_block,
		ToBlock:   nil, // = latest block
		Addresses: []common.Address{test_address},
		// Topics [][]common.Hash, - add more parameters to filter
	}

	fmt.Println("Querying FilterLogs..")

	logsX, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	var total_volume_token0 *big.Int
	var total_volume_token1 *big.Int

	total_volume_token0 = big.NewInt(0)
	total_volume_token1 = big.NewInt(0)

	fmt.Print("Number of block logs: ")
	fmt.Println(len(logsX))
	fmt.Println("Looping through each retrieved record..")

	//4)  Loop through received data and filter it again
	// For each transaction in logsX - check if it matches lookup criteria - add volume if does:
	for i := 0; i < len(logsX); i++ {

		if logsX[i].Topics[0] != common.HexToHash(poolTopics[0]) && logsX[i].Topics[0] != common.HexToHash(poolTopics[1]){
			continue
		}

		fmt.Print(i)
		fmt.Print(" | tx hash: ")
		fmt.Print(logsX[i].TxHash)
		fmt.Print(" | block #: ")
		fmt.Print(logsX[i].BlockNumber)
		fmt.Print(" | ")

		// Get date from block number
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(logsX[i].BlockNumber)))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(" | time: ")
		fmt.Print(block.Time())
		// ADD - Get other pool characteristic - topics[0], index
		// ADD - Get tokens
		// ADD - If within criteria - get amounts from hash
		txlog, err := client.TransactionReceipt(context.Background(), logsX[i].TxHash)
		//transaction_raw_hash := "0x45f6ed12044e324fc8fd492aa8fb52aa54abfdedcf984121cbe3126e65512f0c"
		//hash := common.HexToHash(transaction_raw_hash)

		if err != nil {
			log.Fatal(err)
		}

		// add to volume
		x0, x1 := getTradingVolumeFromTxLog(txlog.Logs, poolTopics)
		total_volume_token0.Add(total_volume_token0, x0)
		total_volume_token1.Add(total_volume_token1, x1)
		
		fmt.Print(" | x0: ")
		fmt.Print(x0)
		fmt.Print(" | x1: ")
		fmt.Println(x1)
		fmt.Print(" | vlm0 totl: ")
		fmt.Println(total_volume_token0)
		
	}

}

func decodeBytes(log *types.Log) (*big.Int, *big.Int) {

	//fmt.Println("Byte data:")
	//fmt.Println(log.Data)
	//fmt.Println("Asset number:")
	//fmt.Println("Asset number:")
	var32to64 := new(big.Int).SetBytes(log.Data[32:64])
	//fmt.Println(var0to32)
	var96to128 := new(big.Int).SetBytes(log.Data[96:128])
	//fmt.Println(var32to64)

	return var32to64, var96to128
}


func getTradingVolumeFromTxLog(logs []*types.Log, pooltopics []string) (actualIn *big.Int, actualOut *big.Int) {

	var firstLog *types.Log
	//var lastLog *types.Log

	for _, log := range logs {
		if log.Topics[0] != common.HexToHash(pooltopics[0]) && log.Topics[0] != common.HexToHash(pooltopics[1]) {
			continue
		}
		if firstLog == nil {
			firstLog = log
		}
		//lastLog = log
	}


	if firstLog == nil { // could not find any valid swaps, thus the transaction failed
		return common.Big0, common.Big0
	}
	asset0, asset1 := decodeBytes(firstLog)
	/*
	if asset0In.Cmp(common.Big0) > 0 && asset1In.Cmp(common.Big0) == 0 {
		actualIn = asset0In
	} else if asset0In.Cmp(common.Big0) == 0 && asset1In.Cmp(common.Big0) > 0 {
		actualIn = asset1In
	} else if asset0In.Cmp(common.Big0) == 0 && asset1In.Cmp(common.Big0) == 0 {
		panic(fmt.Sprintf("PANIC 00 - Could not decode transaction %s", logs[0].TxHash.Hex()))
		return common.Big0, common.Big0
	} else {
		actualIn = asset0In
	}
	
	asset0Out, asset1Out := decodeBytes(lastLog)
	if asset0Out.Cmp(common.Big0) > 0 && asset1Out.Cmp(common.Big0) == 0 {
		actualOut = asset0Out
	} else if asset0Out.Cmp(common.Big0) == 0 && asset1Out.Cmp(common.Big0) > 0 {
		actualOut = asset1Out
	} else if asset0Out.Cmp(common.Big0) == 0 && asset1Out.Cmp(common.Big0) == 0 {
		panic(fmt.Sprintf("PANIC 01 - Could not decode transaction %s", logs[0].TxHash.Hex()))
		return common.Big0, common.Big0
	} else {
		actualOut = asset0Out
	}
	*/
	return asset0, asset1
}