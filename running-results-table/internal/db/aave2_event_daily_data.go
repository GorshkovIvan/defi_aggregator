package db

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"
	token "./token"
	aaveDataProvider "./aave_protocol_data_provider"
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
	currentBalance *big.Int
	rate_types []int
	flashLoanVolumes []*big.Int
	flashLoanFees []*big.Int
	fees []*big.Float
	decimals int64
	timestamp []int64
	
}

type AaveDailyData struct {

	assetAddress string 
	assetName string
	weightedAverageInterest *big.Float

	volumes []*big.Int 

}





func getAave2Data(){

	var token_check []string
	token_check = append(token_check, "DAI")
	oldest_available_record := get_newest_timestamp_from_db("Aave2", token_check)

	if time.Since(oldest_available_record).Hours() > 35 {
		return
	}	

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	aave_pool_address := common.HexToAddress("0x057835Ad21a177dbdd3090bB1CAE03EaCF78Fc6d")

	
	aave2_data_provider, err := aaveDataProvider.NewStore(aave_pool_address, client)

	if err != nil {
		log.Fatal(err)
	}

	var aave_daily_data []AavePoolData
	days_needed := 1

	for i := 31; i > (31 - days_needed); i-- {

		fmt.Print("Day: ")
		fmt.Println(i)
		aave_daily_data = getAave2DataDaily(client, aave_daily_data, i, aave2_data_provider)
		
	}

	//aave_daily_data = getAave2DataDaily(client, aave_daily_data, 2, aave2_data_provider)
	//fmt.Println("Day 2 done")

	/*	
	for i := 0; i < len(aave_daily_data); i++ {
		fmt.Println("Name")
		fmt.Println(aave_daily_data[i].assetName)
		fmt.Println("Volumes")
		fmt.Println(aave_daily_data[i].volumes)
		fmt.Println("Fees")
		fmt.Println(aave_daily_data[i].fees)

	}*/

	// Getting current blanances for aave2 

	
	for i := 0; i < len(aave_daily_data); i++ {


		pool_address := common.HexToAddress(aave_daily_data[i].assetAddress)
		reserveData, err := aave2_data_provider.GetReserveData(&bind.CallOpts{}, pool_address)

		fmt.Println(aave_daily_data[i].assetName)

		if err != nil {
			log.Fatal(err)
		}

	
		for j := 0; j < days_needed; j++ {
			fmt.Println("Day: ")
			fmt.Print(j)
			
			sum := big.NewInt(0)
			availableLiquidity := reserveData.AvailableLiquidity
			totalStableDebt := reserveData.TotalStableDebt
			totalVariableDebt := reserveData.TotalVariableDebt
			sum.Add(sum, availableLiquidity)
			sum.Add(sum, totalStableDebt)
			sum.Add(sum, totalVariableDebt)
			aave_daily_data[i].currentBalance = sum
			/*
			fmt.Println("Decimals")
			fmt.Println(aave_daily_data[i].decimals)
			fmt.Println("currentBalance")
			fmt.Println(aave_daily_data[i].currentBalance)
			fmt.Println("currentBalance float")
			totalBalanceFloat := negPow(new(big.Float).SetInt(aave_daily_data[i].currentBalance), aave_daily_data[i].decimals)
			fmt.Println(totalBalanceFloat)
			*/
			
			
			
			zero := Zero()
			if aave_daily_data[i].fees[j].Cmp(zero) == 0 {

				fmt.Print("Returns iteration ")
				fmt.Print(j)
				fmt.Println(" : ")
				
				fmt.Println(0.0)
				
				var token_ []string
				token_ = append(token_, aave_daily_data[i].assetName)
				
				
				append_record_to_database("Aave2", token_, timestamp, 0, 0, 0, 0.0, 0.0)
				
				continue
			}
			
			balanceFloat := negPow(new(big.Float).SetInt(aave_daily_data[i].currentBalance), aave_daily_data[i].decimals)
			volumesFloat := negPow(new(big.Float).SetInt(aave_daily_data[i].volumes[j]), aave_daily_data[i].decimals)
			weightedAverageInterest := Div(aave_daily_data[i].fees[j], volumesFloat)
			totalBorrowed := big.NewInt(0)
			totalBorrowed.Add(totalBorrowed, totalStableDebt)
			totalBorrowed.Add(totalBorrowed, totalVariableDebt)
			totalBorrowedFloat := negPow(new(big.Float).SetInt(totalBorrowed), aave_daily_data[i].decimals)
			utilisationRate := Div(totalBorrowedFloat, balanceFloat)
			returns := Mul(utilisationRate, weightedAverageInterest)

			timestamp := aave_daily_data[i].timestamp[j]
			fmt.Print("Timestamp: ")
			fmt.Println(timestamp)
			fmt.Print("Returns: ")
			fmt.Println(returns)
			
			var token_ []string
			token_ = append(token_, aave_daily_data[i].assetName)
			
			append_record_to_database("Aave2", token_, timestamp, 0, 0, 0, weightedAverageInterest.float64(), utilisationRate.float64())
			
		}
		

	}

	//getUsdFromVolumeAave2(aave_daily_data[0])
 	
}

func sumVolumes(volumes []*big.Int) *big.Int {

	sum := big.NewInt(0)
	for i := 0; i < len(volumes); i++{
		sum.Add(sum, volumes[i])
	}

	return sum
}

func sumFees(fees []*big.Float) *big.Float {

	sum := big.NewFloat(0)
	for i := 0; i < len(fees); i++{
		sum.Add(sum, fees[i])
	}

	return sum
}

func getAave2DataDaily(client *ethclient.Client, aave_daily_data []AavePoolData, daysAgo int, data_provider *aaveDataProvider.Store) []AavePoolData {
	fmt.Println("getAave2DataDaily")
	oldest_block := getOldestBlock(client, daysAgo)
	latest_block := getOldestBlock(client, daysAgo - 1)

	
	old_block, err := client.BlockByNumber(context.Background(), oldest_block)
	if err != nil {
	log.Fatal(err)
	}

	fmt.Println("Oldest block: ")
	fmt.Println(old_block.Time()) 

	
	latest_time := int64(BoD(time.Unix(int64(old_block.Time()), 0)).Unix())
	
	/*
	late_block, err := client.BlockByNumber(context.Background(), latest_block)
	if err != nil {
	log.Fatal(err)
	}

	fmt.Println("Latest block:")
	fmt.Println(late_block.Time())
	latest_time := late_block.Time()
	*/
	
	pool_address := common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")

	volumes_data := aaveGetPoolVolume(pool_address, oldest_block, latest_block, client, data_provider)
	
	for i := 0; i < len(volumes_data); i++ {
		pool_added := false
		for j := 0; j < len(aave_daily_data); j++ {
			if(volumes_data[i].assetName == aave_daily_data[j].assetName){
				pool_added = true
			}
		}
		if(!pool_added){
			newPool := AavePoolData{assetName : volumes_data[i].assetName, assetAddress: volumes_data[i].assetAddress, decimals: volumes_data[i].decimals}
			for j := 0; j < 31; j++ {
				newPool.volumes = append(newPool.volumes, big.NewInt(0))
				newPool.fees = append(newPool.fees, big.NewFloat(0.0))
				newPool.timestamp = append(newPool.timestamp, int64(0))
			}
			aave_daily_data = append(aave_daily_data, newPool)
		}
	}

	for i := 0; i < len(aave_daily_data); i++ {

		for j := 0; j < len(volumes_data); j++{
			if(volumes_data[j].assetName == aave_daily_data[i].assetName){
				
				total_volumes := sumVolumes(volumes_data[j].volumes)
				aave_daily_data[i].volumes[31 - daysAgo] = total_volumes
				total_fees := sumFees(volumes_data[j].fees)
				aave_daily_data[i].fees[31 - daysAgo] = total_fees
				
			}
			aave_daily_data[i].timestamp[31 - daysAgo] = int64(latest_time)
		}

	}
	//volumes_data = aaveGetFlashLoansVolume(pool_address, oldest_block, client, volumes_data)


	return aave_daily_data

}

func BoD(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
/*
func addDecimals(amount *big.Int, assetAddress string, client *ethclient.Client) *big.Float {

	var decimals64 int64

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
		decimals64 = int64(decimals)

	}else{
		decimals = big.NewInt(18)
		decimals64 = int64(decimals)
	}

	amountFloat := Big.NewFloat(amount)

	return negPow(amountFloat, decimals64)


}*/

func getOldestBlock(client *ethclient.Client, daysAgo int) *big.Int {
	fmt.Println("Getting oldest block")
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
	time_needed := uint64(now.Unix()) - 24 * 60 * 60 * uint64(daysAgo)
	
	var j int64
	j = 0

	for {
		j -= 50
		oldest_block.Add(oldest_block, big.NewInt(j))

		block, err := client.BlockByNumber(context.Background(), oldest_block)
		if err != nil {
			log.Fatal(err)
		}

		if block.Time() < time_needed {

			break
		}
	}
	fmt.Println("Got to the end of the oldest block")
	return oldest_block
}




func aaveGetPoolVolume(pool_address common.Address, oldest_block *big.Int, latest_block *big.Int, client *ethclient.Client, aaveDataProvider *aaveDataProvider.Store) []AavePoolData {
	fmt.Println("aaveGetPoolVolume")
	var pools []AavePoolData

	poolTopics := []string{"0xc6a898309e823ee50bac64e45ca8adba6690e99e7841c45d754e2a38e9019d9b"}
	

	query := ethereum.FilterQuery{

		FromBlock: oldest_block,
		ToBlock:   latest_block, // = latest block
		Addresses: []common.Address{pool_address},
	}

	logsX, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

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
			pools[i].fees = append(pools[i].fees, calculateFee(amount, interest_rate, pools[i].decimals))
			
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
				name, err = instance.Symbol(&bind.CallOpts{})

				if err != nil {
					name = "Unknown"
					//log.Fatal(err)
				}
			}else{
				name = "Eth"
			}
			

			interest_rates := []*big.Int{interest_rate}
			volumes := []*big.Int{amount}
			rate_types := []int{rate_type}
			// Calculating decimals 

			//aave_pool_address := common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")

			//aave2_data_provider, err := aaveDataProvider.NewStore(aave_pool_address, client)
		
			if err != nil {
				log.Fatal(err)
			}
			tokenAddress := common.HexToAddress(assetAddress)

			reserve_data, err := aaveDataProvider.GetReserveConfigurationData(&bind.CallOpts{}, tokenAddress)
			
			if err != nil {
				log.Fatal(err)
			} 

			decimals := (reserve_data.Decimals).Int64()
			fee := calculateFee(amount, interest_rate, decimals)
			fees := []*big.Float{fee}
			pools = append(pools, AavePoolData{assetAddress: assetAddress, interest_rates: interest_rates, volumes: volumes, rate_types: rate_types, 
			assetName: name, fees: fees, decimals: decimals})

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

	amount := new(big.Int).SetBytes(log.Data[32:64])
	rate_type, _ := strconv.Atoi((new(big.Int).SetBytes(log.Data[64:96])).String())
	interest_rate := new(big.Int).SetBytes(log.Data[96:128])
	
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

func calculateFee(volume *big.Int, interest *big.Int, decimal int64) *big.Float {
	//fmt.Println("volume_float and volume_interest BEFORE: ")
	//fmt.Println(volume)
	//fmt.Println(interest)
	volume_float := negPow(new(big.Float).SetInt(volume), decimal)
	volume_interest := negPow(new(big.Float).SetInt(interest), 27)
	//fmt.Println("volume_float and volume_interest AFTER: ")
	//fmt.Println(volume_float)
	//fmt.Println(volume_interest)
	return Mul(volume_float, volume_interest)
}

func negPow(a *big.Float, e int64) *big.Float {
    result := Zero().Copy(a)
	divTen := big.NewFloat(0.1)
    for i:=int64(0); i<e; i++ {
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

func Div(a, b *big.Float) *big.Float {
    return Zero().Quo(a, b)
}

