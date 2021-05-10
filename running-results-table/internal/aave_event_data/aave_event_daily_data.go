package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"
	token "./erc20Interface"
	//aaveDataProvider "./aave_protocol_data_provider"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"encoding/hex"
	"strings"
)

type AavePoolData struct {

	assetAddress string 
	assetName string
	interest_rates []*big.Int 
	volumes []*big.Int 
	rate_types []int
	flashLoanVolumes []*big.Int
	flashLoanFees []*big.Int
	
	

}

type AaveDailyData struct {

	//assetAddress string 
	assetName string
	volumes []*big.Int 

}


func main(){

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	var aave_daily_data []AaveDailyData

	aave_daily_data = getAaveData(client, aave_daily_data, 2)
	fmt.Println("Day 1 done")
	aave_daily_data = getAaveData(client, aave_daily_data, 3)
	fmt.Println("Day 2 done")


	for i := 0; i < len(aave_daily_data); i++ {
		fmt.Println("Name")
		fmt.Println(aave_daily_data[i].assetName)
		fmt.Println("Volumes")
		fmt.Println(aave_daily_data[i].volumes)

	}



}

func sumVolumes(volumes []*big.Int) *big.Int {

	sum := big.NewInt(0)
	for i := 0; i < len(volumes); i++{
		sum.Add(sum, volumes[0])
	}

	return sum
}

func getAaveData(client *ethclient.Client, aave_daily_data []AaveDailyData, daysAgo int) []AaveDailyData {

	oldest_block := getOldestBlock(client, daysAgo)
	latest_block := getOldestBlock(client, daysAgo - 1)
	pool_address := common.HexToAddress("0x398ec7346dcd622edc5ae82352f02be94c62d119")


	volumes_data := aaveGetPoolVolume(pool_address, oldest_block, latest_block, client)
	
	for i := 0; i < len(volumes_data); i++ {
		pool_added := false
		for j := 0; j < len(aave_daily_data); j++ {
			if(volumes_data[i].assetName == aave_daily_data[j].assetName){
				pool_added = true
			}
		}
		if(!pool_added){
			newPool := AaveDailyData{assetName : volumes_data[i].assetName}
			aave_daily_data = append(aave_daily_data, newPool)
		}
	}

	for i := 0; i < len(aave_daily_data); i++ {

		for j := 0; j < len(volumes_data); j++{
			if(volumes_data[j].assetName == aave_daily_data[i].assetName){
				
				total_volumes := sumVolumes(volumes_data[j].volumes)
				aave_daily_data[i].volumes = append(aave_daily_data[i].volumes, total_volumes)
			}
		}

	}
	//volumes_data = aaveGetFlashLoansVolume(pool_address, oldest_block, client, volumes_data)

	/*
	// Adding decimal spaces 
	for i := range volumes_data{ 

		assetAddress := volumes_data[i].assetAddress

		if(assetAddress != "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"){
					
			tokenAddress := common.HexToAddress(assetAddress)
			instance, err := token.NewToken(tokenAddress, client)
			if err != nil {
				log.Fatal(err)
			}
			decimals, err = instance.Decimals(&bind.CallOpts{})
			if err != nil {
				log.Fatal(err)
			}

		}else{
			decimals = big.NewInt(18)
		}

		// add decimals 


	
	}
	*/
	return aave_daily_data

}

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
	timeonemonthago := uint64(now.Unix()) - 24 * 60 * 60 * uint64(daysAgo)
	
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




func aaveGetPoolVolume(pool_address common.Address, oldest_block *big.Int, latest_block *big.Int, client *ethclient.Client) []AavePoolData {

	var pools []AavePoolData

	poolTopics := []string{"0x1e77446728e5558aa1b7e81e0cdab9cc1b075ba893b740600c76a315c2caa553"}
	

	query := ethereum.FilterQuery{

		FromBlock: oldest_block,
		ToBlock:   latest_block, // = latest block
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
		amount, rate_type, interest_rate, assetAddress := getTradingVolumeFromTxLog(txlog.Logs, poolTopics)
		allocated = false

		for i := range pools{
			if pools[i].assetAddress == assetAddress{
			
			allocated = true 
			//fmt.Println("Appended %s", assetAddress)
			pools[i].volumes = append(pools[i].volumes, amount)
			pools[i].interest_rates = append(pools[i].interest_rates, interest_rate)
			pools[i].rate_types = append(pools[i].rate_types, rate_type)

			}
		}
		var name string 
		if !allocated{
			if(assetAddress != "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"){
				
				tokenAddress := common.HexToAddress(assetAddress)
				instance, err := token.NewToken(tokenAddress, client)
				if err != nil {
					log.Fatal(err)
				}
				name, err = instance.Name(&bind.CallOpts{})
				if err != nil {
					log.Fatal(err)
				}
			}else{
				name = "Eth"
			}

			interest_rates := []*big.Int{interest_rate}
			volumes := []*big.Int{amount}
			rate_types := []int{rate_type}
			
			pools = append(pools, AavePoolData{assetAddress: assetAddress, interest_rates: interest_rates, volumes: volumes, rate_types: rate_types, 
			assetName: name })
			fmt.Println("pool added:")
			fmt.Println(assetAddress)
		}
		/*
		fmt.Println("amount:")
		fmt.Println(amount)
		fmt.Println("rate_type:")
		fmt.Println(rate_type)
		fmt.Println("interest_rate:")
		fmt.Println(interest_rate)
		*/
	}
	return pools
}

func decodeBytes(log *types.Log) (*big.Int, int, *big.Int) {


	amount := new(big.Int).SetBytes(log.Data[0:32])
	rate_type, _ := strconv.Atoi((new(big.Int).SetBytes(log.Data[32:64])).String())
	interest_rate := new(big.Int).SetBytes(log.Data[64:96])

	return amount, rate_type, interest_rate

}
/*
func decodeBytes(log *types.Log) (*big.Int, int, *big.Int) {


	amount := new(big.Int).SetBytes(log.Data[0:32])
	rate_type, _ := strconv.Atoi((new(big.Int).SetBytes(log.Data[32:64])).String())
	interest_rate := new(big.Int).SetBytes(log.Data[64:96])

	return amount, rate_type, interest_rate
}*/

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
	//var lastLog *types.Log

	for _, log := range logs {
		if log.Topics[0] != common.HexToHash(pooltopics[0]) {
			continue
		}
		if firstLog == nil {

			firstLog = log
			address := log.Topics[1]
			assetAddress = HashToReserveAddress(address)
			
		}
		//lastLog = log
	}

	if firstLog == nil { // could not find any valid swaps, thus the transaction failed
		return common.Big0, -1, common.Big0, "none"
	}
	amount, rate_type, interest_rate := decodeBytes(firstLog)

	return amount, rate_type, interest_rate, assetAddress
}

/* Getting flash loans data */

func aaveGetFlashLoansVolume(pool_address common.Address, oldest_block *big.Int, client *ethclient.Client, pools []AavePoolData) [] AavePoolData {

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


		for i := range pools{
			if pools[i].assetAddress == assetAddress{
			
			allocated = true 
			//fmt.Println("Appended %s", assetAddress)
			pools[i].flashLoanVolumes = append(pools[i].flashLoanVolumes, amount)
			pools[i].flashLoanFees = append(pools[i].flashLoanFees, deposit_fee)

			}
		}

		var name string 
		if !allocated{
			if(assetAddress != "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"){
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
			}else{
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
	//var lastLog *types.Log

	for _, log := range logs {
		if log.Topics[0] != common.HexToHash(pooltopics[0]) {
			continue
		}
		if firstLog == nil {

			firstLog = log
			address := log.Topics[2]
			assetAddress = HashToReserveAddress(address)
			
		}
		//lastLog = log
	}

	if firstLog == nil { // could not find any valid swaps, thus the transaction failed
		return common.Big0, common.Big0, "none"
	}
	amount, deposit_rate := decodeFlashLoanBytes(firstLog)

	return amount, deposit_rate, assetAddress
}
