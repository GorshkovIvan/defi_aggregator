package main

import (
	"context"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestDecodeSingleSwap(t *testing.T) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
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

	firstLog := txlog.Logs[0]

	asset0In, asset1In, _, _ := decodeSingleSwap(firstLog)

	if(asset0In == nil) {
		t.Errorf("Asset0In is not decoded!")
	}

	if(asset1In == nil) {
		t.Errorf("Asset1In is not decoded!")
	}

}

func TestDecodeActualInAndActualOut(t *testing.T) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
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

	if(actualIn == nil) {
		t.Errorf("ActualIn is not decoded!")
	}

	if(actualOut == nil) {
		t.Errorf("ActualOut is not decoded!")
	}

}