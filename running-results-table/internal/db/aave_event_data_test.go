package db

import (
	"context"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

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

func TestAaveGetPoolVolume(t *testing.T) {

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")

	if err != nil {
		log.Fatal(err)
	}

	oldest_block := getOldestBlock(client)
	pool_address := common.HexToAddress("0x398ec7346dcd622edc5ae82352f02be94c62d119")

	volumes := aaveGetPoolVolume(pool_address, oldest_block, client)

	if volumes == nil {
		t.Errorf("No volume retrieved")
	}
}

func TestGetTradingVolumeFromTxLog(t *testing.T) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	oldest_block := getOldestBlock(client)
	pool_address := common.HexToAddress("0x398ec7346dcd622edc5ae82352f02be94c62d119")

	query := ethereum.FilterQuery{

		FromBlock: oldest_block,
		ToBlock:   nil, // = latest block
		Addresses: []common.Address{pool_address},
	}

	poolTopics := []string{"0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b"}
	logsX, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	txlog, err := client.TransactionReceipt(context.Background(), logsX[0].TxHash)

	rate_type := -1
	assetAddress := ""

	amount, rate_type, interest_rate, assetAddress := getTradingVolumeFromTxLog(txlog.Logs, poolTopics)

	if amount == nil {
		t.Errorf("Index of amount not retrieved!")
	}

	// this fails, check functionality in aave_event_data line 219
	if rate_type == -1 {
		t.Errorf("Volume of rate_type not retrieved!")
	}

	if interest_rate == nil {
		t.Errorf("Index of interest_rate not retrieved!")
	}

	if assetAddress == "" {
		t.Errorf("Volume of assetAddress not retrieved!")
	}

}

func TestAaveGetFlashLoansVolume(t *testing.T) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	oldest_block := getOldestBlock(client)
	//pool_address := common.HexToAddress("0x398ec7346dcd622edc5ae82352f02be94c62d119")
	pool_address := common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")

	if err != nil {
		log.Fatal(err)
	}

	volumes_data := aaveGetPoolVolume(pool_address, oldest_block, client)

	volumes_data = aaveGetFlashLoansVolume(pool_address, oldest_block,
		client, volumes_data)

	if volumes_data == nil {
		t.Errorf("Flash Loans data not retrieved!")
	}
}

func TestGetFlashLoansVolumeFromTxLog(t *testing.T) {
	//client, err := ethclient.Dial("http://localhost:8888")
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	oldest_block := getOldestBlock(client)
	pool_address := common.HexToAddress("0x398ec7346dcd622edc5ae82352f02be94c62d119")

	if err != nil {
		log.Fatal(err)
	}

	poolTopics := []string{"0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55"}

	query := ethereum.FilterQuery{

		FromBlock: oldest_block,
		ToBlock:   nil, // = latest block
		Addresses: []common.Address{pool_address},
	}

	logsX, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(logsX); i++ {

		if logsX[i].Topics[0] != common.HexToHash(poolTopics[0]) {
			continue
		}

		txlog, err := client.TransactionReceipt(context.Background(), logsX[i].TxHash)

		if err != nil {
			log.Fatal(err)
		}

		assetAddress := ""

		amount, deposit_fee, assetAddress := getFlashLoansVolumeFromTxLog(txlog.Logs, poolTopics)

		if amount == nil {
			t.Errorf("Index of amount not retrieved!")
		}

		if deposit_fee == nil {
			t.Errorf("Volume of deposit_fee not retrieved!")
		}

		if assetAddress == "" {
			t.Errorf("Volume of assetAddress not retrieved!")
		}
	}
}

/*
// =========== !!! TO DO !!! ==========
// to do these tests from here
func TestDecodeBytes(log *types.Log) (*big.Int, int, *big.Int) {


	amount := new(big.Int).SetBytes(log.Data[0:32])
	rate_type, _ := strconv.Atoi((new(big.Int).SetBytes(log.Data[32:64])).String())
	interest_rate := new(big.Int).SetBytes(log.Data[64:96])

	return amount, rate_type, interest_rate
}

func HashToReserveAddress(hash common.Hash) string {
	var value []string
	value = append(value, "0", "x")
	value = append(value, hex.EncodeToString(hash[12:32]))
	valueStr := strings.Join(value, "")


	return valueStr
}

func TestDecodeFlashLoanBytes(t *testing.T) {


}*/
