package db

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

func BoD(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func getbalancerdata_from_blockchain() {
	// 0) Connect to client
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	//Quering the db for oldest available Record
	oldest_available_record := time.Now() // get oldest timestamp from DB
	oldest_available_record = oldest_available_record.AddDate(0, 0, -2)

	oldest_lookup_time := time.Now() //.Unix()

	days_ago := 3

	if (time.Now()).Sub(oldest_available_record).Hours() > 24 {
		fmt.Println("DATA IS OLD!!! NEED TO UPDATE!!!!!!!!!")
		oldest_lookup_time = oldest_lookup_time.AddDate(0, 0, -days_ago) // oldest_available_record.Unix()
		// math.Max(now.AddDate(0, 0, -days_ago)
	}

	// oldest_lookup_time := uint64((now.AddDate(0, 0, -days_ago)).Unix())
	fmt.Print("Now: ")
	fmt.Println(time.Now())
	fmt.Print("BoD for today: ")
	fmt.Println(BoD(time.Now()))
	fmt.Print("diff: ")
	t := (time.Now()).Sub(oldest_available_record).Hours()
	fmt.Print(t)

	// 1) If data is old and need to update it - Define pool specific parameters
	var BalancerpoolAddress = common.HexToAddress("0x1eff8af5d577060ba4ac8a29a13525bb0ee2a3d5") //pass in the pool id...
	balancertopic := "0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378"
	fmt.Println("CONNECTED TO INFURA SUCCESSFULLY!!!!!!")

	/*
		var xxx = common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

		tokenAddress := common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599")
		instance, err := token.NewToken(tokenAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		bal, err := instance.BalanceOf(&bind.CallOpts{}, xxx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

		balance, err := client.BalanceAt(context.Background(), BalancerpoolAddress, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("POOL SIZE test: ")
		fmt.Print(balance)
	*/

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
	// now := time.Now()

	// func get_oldest_timestamp_from_db(pool string, token0 string, token1 string) uint64
	// func append_record_to_database(pool string, token0 string, token1 string, date uint64, trading_volume_usd int64, pool_sz_usd int64)
	// func create_new_hist_volume_poolsz_entry(pool string, token0 string, token1 string)

	var j int64
	j = 0
	// compute block id [30] days away from now
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
		// BlockHash *common.Hash, - add more parameters to filter
		FromBlock: oldest_block,
		ToBlock:   nil, // = latest block
		Addresses: []common.Address{BalancerpoolAddress},
	}
	fmt.Println("Querying FilterLogs..")

	logsX, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//var total_volume_token0 *big.Int // cumulative
	//var total_volume_token1 *big.Int // cumulative

	//total_volume_token0 = big.NewInt(0)
	//total_volume_token1 = big.NewInt(0)

	fmt.Print("Number of block logs: ")
	fmt.Print(len(logsX))
	fmt.Println(" ...Looping through each retrieved record..")

	// Count volumes by day
	var dates []int64
	var tradingvolumes []int64

	cumulative_for_day := int64(0)

	t_prev := uint64(0) // oldest time logsX[0].
	t_new := uint64(0)  // oldest time

	//4)  Loop through received data and filter it again
	// For each transaction in logsX - check if it matches lookup criteria - add volume if does:
	for i := 0; i < len(logsX); i++ {
		if logsX[i].Topics[0] != common.HexToHash(balancertopic) {
			continue
		}

		//if logsX[i].TxHash == common.HexToHash("0x7eb7275cd5cd317c4805e9eae203a0e2cddbf45bef3a72108506701f6d6a5fff") {
		fmt.Print(i)
		fmt.Print(" | th: ")
		fmt.Print(logsX[i].TxHash)
		fmt.Print(" | blk#: ")
		fmt.Print(logsX[i].BlockNumber)
		fmt.Print(" | ")

		//Get date from block number
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

		/*
			DAILY COUNT LOOP GOES HERE
		*/
		t_prev = t_new       // uint
		t_new = block.Time() // uint

		if t_prev == 0 || t_new-t_prev > 24*60*60 { // 1 day
			fmt.Println("DATA OLD!!!! APPENDING NEW DAY!!!!")
			dates = append(dates, int64(BoD(time.Unix(int64(t_new), 0)).Unix()))
			tradingvolumes = append(tradingvolumes, cumulative_for_day)
			cumulative_for_day = 0
		} else {
			x0, x1 := getTradingVolumeFromTxLog2(txlog.Logs, balancertopic)
			// convert to usd
			// decimals
			//f := new(big.Float).SetInt(x0)
			fmt.Print(x1)
			var smallnum, _ = new(big.Int).SetString(x0.Text(16), 10)
			num := smallnum.Uint64()
			cumulative_for_day += int64(num)
		}

		/*
			for i := 0; i < len(dates); i++ {
				// append record to database
				// also need to get volume of pool
			}

		*/

		// add to volume
		x0, x1 := getTradingVolumeFromTxLog2(txlog.Logs, balancertopic)
		//total_volume_token0.Add(total_volume_token0, x0)
		// total_volume_token1.Add(total_volume_token1, x1)

		fmt.Print(" | x0: ")
		fmt.Print(x0)
		fmt.Print(" | x1: ")
		fmt.Println(x1)
		//fmt.Print(" | cum0: ")
		//fmt.Println(total_volume_token0)
		//} // if txhash

	} // loop through log finished

}

/*
func decodeBytes(log *types.Log) (*big.Int, *big.Int) {
	fmt.Println("Decoding bytes!!!!! ")
	var0to32 := new(big.Int).SetBytes(log.Data[0:32])
	var32to64 := new(big.Int).SetBytes(log.Data[32:64])

	fmt.Print(" | in dcb - x0: ")
	fmt.Print(var0to32)
	fmt.Print(" | x1: ")
	fmt.Print(var32to64)

	return var0to32, var32to64
}
*/
func getTradingVolumeFromTxLog2(logs []*types.Log, pooltopic string) (actualIn *big.Int, actualOut *big.Int) {
	fmt.Println("Decoding bytes: ")

	var0to32 := big.NewInt(0)
	var32to64 := big.NewInt(0)

	j := 4

	if len(logs) >= j {
		if len(logs[j].Data) > 0 {
			var0to32 = new(big.Int).SetBytes(logs[j].Data[0:32])
		}
		if len(logs[j].Data) > 0 {
			var32to64 = new(big.Int).SetBytes(logs[j].Data[32:64])
		}
	}

	for i := 0; i < len(logs); i++ {
		fmt.Print(i)
		if len(logs[i].Data) > 0 {
			// var0to32 = new(big.Int).SetBytes(logs[i].Data[0:32])
			fmt.Print(" | 0-32: ")
			fmt.Print(var0to32)
		}
		if len(logs[i].Data) > 32 {
			// var32to64 = new(big.Int).SetBytes(logs[i].Data[32:64])
			fmt.Print(" | 32-64: ")
			fmt.Print(var32to64)
		}
		if len(logs[i].Topics) > 0 {
			fmt.Print(" | topic 0: ")
			fmt.Print(logs[i].Topics[0])
		}
		if len(logs[i].Topics) > 1 {
			fmt.Print(" | topic 1: ")
			fmt.Print(logs[i].Topics[1])
		}
		if len(logs[i].Topics) > 2 {
			fmt.Print(" | topic 2: ")
			fmt.Print(logs[i].Topics[2])
		}
		fmt.Println(" ")
	}

	//	if len(logs[0].Data) >= 64 && len(logs) >= 1 {
	return var0to32, var32to64
	//	} else {
	//		return big.NewInt(0), big.NewInt(0)
	//	}
}

// xxxx
/*
func getTradingVolumeFromTxLog(logs []*types.Log, pooltopic string) (actualIn *big.Int, actualOut *big.Int) {
	fmt.Println("Getting Trading Volume from logs!!! ")
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

	lenZ := len(logs)
	fmt.Print("| len of log: ")
	fmt.Print(lenZ)

	for i := 0; i < lenZ; i++ {
		x0 := new(big.Int).SetBytes(logs[i].Data[0:32])
		x1 := len(logs[i].Data)
		x2 := len(logs[i].Topics)
		var x3 common.Hash
		var x4 common.Hash

		var x5 *big.Int //big.Rat
		if x2 >= 1 {
			x3 = logs[i].Topics[0]
		}
		if x2 >= 2 {
			x4 = logs[i].Topics[1]
		}
		if x1 > 32 {
			x5 = new(big.Int).SetBytes(logs[i].Data[32:64])
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

	return actualIn, actualOut
}
*/
