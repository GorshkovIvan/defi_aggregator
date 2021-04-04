package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
    //"strings"
	
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
    //"github.com/ethereum/go-ethereum"
    //"github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    //"github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)


//type Event interface{}

func main(){

	client, err := ethclient.Dial("http://localhost:8888")
    if err != nil {
        log.Fatal(err)
    }

	//var poolAddress = common.HexToAddress("0x8b6e6e7b5b3801fed2cafd4b22b8a16c2f2db21a")

	//balancerEvents := NewBalancerEvents(poolAddress, client)
	hash := common.HexToHash("0x45f6ed12044e324fc8fd492aa8fb52aa54abfdedcf984121cbe3126e65512f0c")
	txlog, err := client.TransactionReceipt(context.Background(), hash)

	if err != nil {
        log.Fatal(err)
    }

	

	actualIn, actualOut := decodeActualInAndActualOut(txlog.Logs)


	fmt.Println(actualIn)
	fmt.Println(actualOut)

}

func decodeSingleSwap(log *types.Log) (*big.Int, *big.Int, *big.Int, *big.Int) {
	fmt.Println(new(big.Int).SetBytes(log.Data))
	amount0In := new(big.Int).SetBytes(log.Data[0:32])
	amount1In := new(big.Int).SetBytes(log.Data[32:64])
	amount0Out := new(big.Int).SetBytes(log.Data[64:96])
	amount1Out := new(big.Int).SetBytes(log.Data[96:128])
	return amount0In, amount1In, amount0Out, amount1Out
 }
 func decodeActualInAndActualOut(logs []*types.Log) (actualIn *big.Int, actualOut *big.Int) {
	var firstLog *types.Log
	var lastLog *types.Log
	for _, log := range logs {
	   if log.Topics[0] != common.HexToHash("0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378") {
		  continue
	   }
	   if firstLog == nil {
		  firstLog = log
	   }
	   lastLog = log
	}
	// could not find any valid swaps, thus the transaction failed
	if firstLog == nil {
	   return common.Big0, common.Big0
	}
	asset0In, asset1In, _, _ := decodeSingleSwap(firstLog)
	if asset0In.Cmp(common.Big0) > 0 && asset1In.Cmp(common.Big0) == 0 {
	   actualIn = asset0In
	} else if asset0In.Cmp(common.Big0) == 0 && asset1In.Cmp(common.Big0) > 0 {
	   actualIn = asset1In
	} else {
	   // something is wrong, we can not have 0 inputs for a successful swap
	   //panic(fmt.Sprintf("Could not decode transaction %s", logs[0].TxHash.Hex()))
	   return common.Big0, common.Big0
	}
	_, _, asset0Out, asset1Out := decodeSingleSwap(lastLog)
	if asset0Out.Cmp(common.Big0) > 0 && asset1Out.Cmp(common.Big0) == 0 {
	   actualOut = asset0Out
	} else if asset1Out.Cmp(common.Big0) > 0 && asset0Out.Cmp(common.Big0) == 0 {
	   actualOut = asset1Out
	} else {
	   // something is wrong, we can not have 0 outputs for a successful swap
	   //panic(fmt.Sprintf("Could not decode transaction %s", logs[0].TxHash.Hex()))
	   return common.Big0, common.Big0
	}
	return actualIn, actualOut
 }