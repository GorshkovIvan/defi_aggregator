package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
Schematic

Historical volumes
0) Connect to client
1) Define pool-specific lookup variabels
2) Find oldest block - iterate through block numbers, check their time - once find T-31 - this is out oldest
3) Query between oldest and current block for Balancer-specific addresses
4) Loop through received data - filter out by a) date (minor changes), b) pool, c) token(?) --> TOTAL VOLUMES
5) Get gas fees paid

Historical pool sz
0) For each historical snapshot (How?): Get wallet list from pool
1) For each wallet get balance - add them up

Questions remaining:
- Can we avoid querying all blocks first?
- Querying a lot of days is very slow (30 days is 200k blocks..)
- How do we get rid of localhost? - need for real code
- Should we add more parameters to filterquery? - BlockHash + CommonHash?
- Numbers decoder function gets seem off - are these in Gwei? (check with TheGraph)
- How to get historical wallets in pool?
- How to get other parameters - tokens in pool,
- What are topic[1] etc?
*/

func main() {
	// 0) Connect to client
	client, err := ethclient.Dial("http://localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	// 1) Define pool specific parameters
	var BalancerpoolAddress = common.HexToAddress("0x8b6e6e7b5b3801fed2cafd4b22b8a16c2f2db21a")
	balancertopic := "0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378"

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
	timeonemonthago := uint64((now.AddDate(0, 0, -1)).Unix())
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
	query := ethereum.FilterQuery{
		// BlockHash *common.Hash, - add more parameters to filter
		FromBlock: oldest_block,
		ToBlock:   nil, // = latest block
		Addresses: []common.Address{BalancerpoolAddress},
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

		if logsX[i].Topics[0] != common.HexToHash(balancertopic) {
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
		x0, x1 := getTradingVolumeFromTxLog(txlog.Logs, balancertopic)
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
	var0to32 := new(big.Int).SetBytes(log.Data[0:32])
	var32to64 := new(big.Int).SetBytes(log.Data[32:64])
	/*
		fmt.Print(" | in dcb - x0: ")
		fmt.Print(var0to32)
		fmt.Print(" | x1: ")
		fmt.Print(var32to64)
	*/
	return var0to32, var32to64
}

func getTradingVolumeFromTxLog(logs []*types.Log, pooltopic string) (actualIn *big.Int, actualOut *big.Int) {
	var firstLog *types.Log
	var lastLog *types.Log

	for _, log := range logs {
		if log.Topics[0] != common.HexToHash(pooltopic) {
			continue
		}
		if firstLog == nil {
			firstLog = log
		}
		lastLog = log
	}

	/* Print routines - for debugging
	lenZ := len(logs)
	fmt.Print("| len of log: ")
	fmt.Print(lenZ)

	for i := 0; i < lenZ; i++ {
		x0 := new(big.Int).SetBytes(txlog.Logs[i].Data[0:32])
		x1 := len(txlog.Logs[i].Data)
		x2 := len(txlog.Logs[i].Topics)
		var x3 common.Hash
		//	var x4 common.Hash
		var x5 *big.Int //big.Rat
		if x2 >= 1 {
			x3 = txlog.Logs[i].Topics[0]
		}
		//if x2 >= 2 {
		//		x4 = txlog.Logs[i].Topics[1]
		//	}
		if x1 > 32 {
			x5 = new(big.Int).SetBytes(txlog.Logs[i].Data[32:64])
		}

		fmt.Print(i)
		fmt.Print(" | x0: ")
		fmt.Print(x0)
		fmt.Print(" | x1: ")
		fmt.Print(x1)
		fmt.Print(" | x2: ")
		fmt.Print(x2)
		fmt.Print(" | x3: ")
		fmt.Print(x3)
		fmt.Print(" | ")
		fmt.Print("x4: ")
		fmt.Print(x4)
		fmt.Print(" | x5: ")
		fmt.Println(x5)
	}
	*/

	if firstLog == nil { // could not find any valid swaps, thus the transaction failed
		return common.Big0, common.Big0
	}
	asset0In, asset1In := decodeBytes(firstLog)
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

	/*
		fmt.Print("| RETURNING..in: ")
		fmt.Print(actualIn)
		fmt.Print("| out: ")
		fmt.Print(actualOut)
	*/
	return actualIn, actualOut
}
