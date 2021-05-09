package db

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"pusher/defi_aggregator/running-results-table/internal/db/token"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	//	"github.com/ethereum/go-ethereum/ethclient"
)

func getbalancerdata_from_blockchain() {
	// 0) Connect to client
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	oldest_available_record := time.Now() // XX - GET IT USING AARON's func
	// func get_oldest_timestamp_from_db(pool string, token0 string, token1 string) uint64
	oldest_available_record = oldest_available_record.AddDate(0, 0, -2)
	oldest_lookup_time := time.Now() //.Unix()
	days_ago := 3

	if (time.Since(oldest_available_record).Hours()) > 24 {
		fmt.Println("DATA IS OLD!!! NEED TO UPDATE!!!!!!!!!")
		oldest_lookup_time = oldest_lookup_time.AddDate(0, 0, -days_ago) // oldest_available_record.Unix()
		// math.Max(now.AddDate(0, 0, -days_ago)
	}

	fmt.Print("Now: ")
	fmt.Println(time.Now())
	fmt.Print("BoD for today: ")
	fmt.Println(BoD(time.Now()))
	fmt.Print("diff: ")
	t := (time.Since(oldest_available_record).Hours())
	fmt.Print(t)

	// 1) If data is old and need to update it - Define pool specific parameters
	var BalancerpoolAddress = common.HexToAddress("0x1eff8af5d577060ba4ac8a29a13525bb0ee2a3d5") //pass in the pool id...
	balancertopic := "0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378"
	fmt.Println("  CONNECTED TO INFURA SUCCESSFULLY!!!!!!")

	tokenAddress := common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599")

	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := instance.BalanceOf(&bind.CallOpts{}, BalancerpoolAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("bal: %s\n", bal)

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

	var j int64
	j = 0 // compute block id [30] days away from now
	for {
		j -= 25
		oldest_block.Add(oldest_block, big.NewInt(j))

		block, err := client.BlockByNumber(context.Background(), oldest_block)
		if err != nil {
			log.Fatal(err)
		}

		if block.Time() <= uint64(oldest_lookup_time.Unix()) {
			fmt.Print("nd oldest lookup block: ") // 5671744
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
	fmt.Println(" ...Looping through each retrieved record..")

	// Count volumes by day
	var dates []int64
	var tradingvolumes []int64
	var poolsizes []int64

	cumulative_for_day := int64(0)

	t_prev := uint64(0) // oldest time logsX[0].
	t_new := uint64(0)  // oldest time

	//4)  Loop through received data and filter it again
	// For each transaction in logsX - check if it matches lookup criteria - add volume if does:
	for i := 0; i < len(logsX); i++ {
		if logsX[i].Topics[0] != common.HexToHash(balancertopic) {
			continue
		}

		fmt.Print(i)
		fmt.Print(" | th: ")
		fmt.Print(logsX[i].TxHash)
		fmt.Print(" | blk#: ")
		fmt.Print(logsX[i].BlockNumber)
		fmt.Print(" | ")

		// Get date from block number
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(logsX[i].BlockNumber)))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(" | t: ")
		fmt.Print(block.Time())

		txlog, err := client.TransactionReceipt(context.Background(), logsX[i].TxHash)
		if err != nil {
			log.Fatal(err)
		}

		t_prev = t_new       // uint
		t_new = block.Time() // uint

		if t_prev == 0 || (t_new-uint64(math.Mod(float64(t_new), 86400)))/86400 != (t_prev-uint64(math.Mod(float64(t_prev), 86400)))/86400 { // 1 day
			dates = append(dates, int64(BoD(time.Unix(int64(t_new), 0)).Unix()))
			tradingvolumes = append(tradingvolumes, cumulative_for_day)
			poolsizes = append(poolsizes, bal.Int64())
			cumulative_for_day = 0

			fmt.Println(" t new:")
			fmt.Print(t_new)
			fmt.Print(" | prev: ")
			fmt.Print(t_prev)
			fmt.Print("day crossed: ")
			fmt.Print(int64(BoD(time.Unix(int64(t_new), 0)).Unix()))
			fmt.Print("..cumulative: ")
			fmt.Println(cumulative_for_day)
		} else {
			token0AD := "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599" // BTC
			token1AD := "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2" // WETH
			dec := 18
			tkn1 := getTradingVolumeFromTxLog2(txlog.Logs, balancertopic, token0AD, token1AD, dec)
			// convert to usd
			cumulative_for_day += tkn1
			fmt.Print(" | tkn1: ")
			fmt.Print(tkn1)
			fmt.Print(" | cumulative: ")
			fmt.Println(cumulative_for_day)
			// + size of pool
		}
	} // loop through log finished

	fmt.Println("-----------------SUMMARY DAILY: -----------------------------------")
	for i := 0; i < len(dates); i++ {
		fmt.Print("i: ")
		fmt.Print(i)
		fmt.Print("| t: ")
		fmt.Print(dates[i])
		fmt.Print("| volumes: ")
		fmt.Print(tradingvolumes[i])
		// append to db(dates[i],tradingvolumes[i],poolsizes[i])
	}

}

//Saved Codes:
/*
	if len(logs) >= 2 {
		if len(logs[1].Topics) >= 3 {
			if logs[1].Topics[2] == common.HexToHash(token0AD) {
				//fmt.Print("The log was: ")
				//fmt.Print(logs)
				fmt.Print("J = 4 !!! - case 1")
				//tokenaddr := logs[2].Topics[2]
				//fmt.Print("FOUND TOKEN AT: ")
				//fmt.Print(tokenaddr)
				j = 4 // 5
			}
		}
	}

	if len(logs) >= 3 {
		if len(logs[2].Topics) >= 3 {
			if logs[2].Topics[2] == common.HexToHash(token0AD) {
				fmt.Print("J = 4 !!! - case 1.5")
				j = 4 // 5
			}
		}
	}

	//	fmt.Println("checkpoint 1")
	//	fmt.Println(len(logs))

	if len(logs) >= 4 {
		//fmt.Println(len(logs[2].Topics))
		if len(logs[3].Topics) >= 3 {
			if logs[3].Topics[2] == common.HexToHash(token1AD) {
				fmt.Print("J = 0 !!!! - case 2")
				//tokenaddr := logs[3].Topics[2]
				//fmt.Print("FOUND TOKEN AT: ")
				//fmt.Print(tokenaddr)
				j = 0 // 1 3 4
			}
		}
	}

	fmt.Println("checkpoint 2 ")
	fmt.Println(len(logs))

	if len(logs) >= 3 {
		fmt.Println(len(logs[2].Topics))
		if len(logs[2].Topics) >= 3 {
			//fmt.Println(len(logs[2].Topics))
			fmt.Print(logs[2].Topics[2])
			fmt.Println(common.HexToHash(token1AD))
			if logs[2].Topics[2] == common.HexToHash(token1AD) {
				fmt.Print("J = 0 !!!! - case 3")
				j = 0 // 2 3
			}
		}
	}
*/

//fmt.Print("checkpoint 3")

/*
	fmt.Print(" | Trying to match to known ids: ")
	fmt.Print(common.HexToHash(token0AD))
	fmt.Print(" | ")
	fmt.Print(common.HexToHash(token1AD))
*/
//	if tokenaddr == common.HexToHash(token0AD) {
//		j = 0
//	}
//	if tokenaddr == common.HexToHash(token1AD) {
//		j = 3
//	}

/*
	fmt.Print(" | len: ")
	fmt.Print(len(logs))
	fmt.Print(" | logs: ")
	fmt.Print(logs)
*/

/*
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
*/

//	fmt.Println(decimals)
//	fmt.Printf("symbol: %s\n", symbol)

// print token0 id
// print token1 id
// match decimals we have from graphql
