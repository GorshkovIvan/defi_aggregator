package main

import (
	//"context"
	"fmt"
	"log"
	"math/big"
	//"time"

	//"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	curveRegistry "./curveRegistry"
)

func main(){

	client, err := ethclient.Dial("http://localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	// 1) Define pool specific parameters
	var curveRegistryAddress = common.HexToAddress("0x7D86446dDb609eD0F5f8684AcF30380a356b2B4c")
	provider, err := curveRegistry.NewMain(curveRegistryAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	number_of_pools, err := provider.PoolCount(&bind.CallOpts{})
	fmt.Println(number_of_pools)

	for i := 0; i < int32(number_of_pools); i++{
		fmt.Println(provider.pool_list(i))
	}




}