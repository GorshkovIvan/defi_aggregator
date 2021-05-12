package db // curve_event_data

import (
	"context"
	"log"
	"math/big"
	"testing"

	//curveRegistry "./curveRegistry"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
func TestGetOldestBlock(t *testing.T) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")

	if err != nil {
		log.Fatal(err)
	}

	oldest_block := getOldestBlock(client)

	if oldest_block == nil {
		t.Errorf("No block retrieved")
	}


}
*/
func TestCurveGetPoolVolume(t *testing.T) {

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")

	if err != nil {
		log.Fatal(err)
	}

	oldest_block := getOldestBlock(client)
	pool_address, err := provider.PoolList(&bind.CallOpts{}, big.NewInt(0))

	volumes, fees := curveGetPoolVolume(pool_address, oldest_block, client)

	if volumes == nil {
		t.Errorf("No volume retrieved")
	}

	if fees == nil {
		t.Errorf("No fees retrieved")
	}

}

/*
func TestDecodeBytesCurve(t *testing.T) {
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
	coin_addresses, err := provider.GetUnderlyingCoins(&bind.CallOpts{}, pool_address)
	if err != nil {
		log.Fatal(err)
	}

	// Getting the number of decimal spaces for undelying coins in the pool
	coin_decimals, err := provider.GetUnderlyingDecimals(&bind.CallOpts{}, pool_address)
	if err != nil {
		log.Fatal(err)
	}

	// Getting current pool balances


	poolTopics := []string{/*"0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140", "0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b"}

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

		if logsX[i].Topics[0] != common.HexToHash(poolTopics[0]) /*&& logsX[i].Topics[0] != common.HexToHash(poolTopics[1]) {
			continue
		}

		txlog, err := client.TransactionReceipt(context.Background(), logsX[i].TxHash)

		if err != nil {
			log.Fatal(err)
		}

		var firstLog *types.Log
		//var lastLog *types.Log

		for _, log := range logs {
			if log.Topics[0] != common.HexToHash(pooltopics[0]) /*&& log.Topics[0] != common.HexToHash(pooltopics[1]) {
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

		if (asset0_index == nil) {
			t.Errorf("No asset0_index decoded")
		}
		if (asset0_volume == nil) {
			t.Errorf("No asset0_volume decoded")
		}
		if (asset1_index == nil) {
			t.Errorf("No asset1_index decoded")
		}
		if (asset1_volume == nil) {
			t.Errorf("No asset1_volume decoded")
		}
	}
}
*/
/*
func testDecodeBytesCurve(t *testing.T) {
	var firstLog *types.Log

	poolTopics := []string{ "0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b"}
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	query := ethereum.FilterQuery{

		FromBlock: oldest_block,
		ToBlock:   nil, // = latest block
		Addresses: []common.Address{pool_address},
	}
	logsX, err := client.FilterLogs(context.Background(), query)

	txlog, err := client.TransactionReceipt(context.Background(), logsX[0].TxHash)

	if err != nil {
		log.Fatal(err)
	}

	var firstLog *types.Log
	//var lastLog *types.Log

	for _, log := range logs {
		if log.Topics[0] != common.HexToHash(pooltopics[0]) /*&& log.Topics[0] != common.HexToHash(pooltopics[1]) {
			continue
		}
		if firstLog == nil {
			firstLog = log
		}
		//lastLog = log
	}


	asset0_index, asset0_volume, asset1_index, asset1_volume := getTradingVolumeFromTxLog(txlog.Logs, poolTopics)

	if (asset0_index == nil) {
		t.Errorf("No asset0_index decoded")
	}
	if (asset0_volume == nil) {
		t.Errorf("No asset0_volume decoded")
	}
	if (asset1_index == nil) {
		t.Errorf("No asset1_index decoded")
	}
	if (asset1_volume == nil) {
		t.Errorf("No asset1_volume decoded")
	}
}*/

func TestGetTradingVolumeFromTxLog(t *testing.T) {
	poolTopics := []string{"0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b"}
	logsX, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	txlog, err := client.TransactionReceipt(context.Background(), logsX[0].TxHash)

	asset0_index, asset0_volume, asset1_index, asset1_volume := getTradingVolumeFromTxLogCurve(txlog.Logs, poolTopics)

	if asset0_index == nil {
		t.Errorf("Index of asset0 not retrieved!")
	}

	if asset0_volume == nil {
		t.Errorf("Volume of asset0 not retrieved!")
	}

	if asset1_index == nil {
		t.Errorf("Index of asset1 not retrieved!")
	}

	if asset1_volume == nil {
		t.Errorf("Volume of asset1 not retrieved!")
	}

}

func TestNegPow(t *testing.T) {

	if negPow(big.NewFloat(1234567890), 10) != 0.123456789 {
		t.Errorf("Wrong value for negative power")
	}
}

func TestZero(t *testing.T) {
	if Zero() != 0.0 {
		t.Errorf("Zero is not returned")
	}
}

/*
func TestMul(t.*testing.T) {
	//to test: don't know what the function is doing lol
}
*/
