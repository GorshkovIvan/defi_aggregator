package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	token "./token"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main(){

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	
	if err != nil {
		log.Fatal(err)
	}

	// Creaitng a contract instance 
	var dai = common.HexToAddress("0x2dded6Da1BF5DBdF597C45fcFaa3194e53EcfeAF")
	provider, err := curveRegistry.NewMain(curveRegistryAddress, client)

	if err != nil {
		log.Fatal(err)
	}

}