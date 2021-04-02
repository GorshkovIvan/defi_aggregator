package main

import (
    "context"
    "fmt"
    "log"
    //"math/big"
    //"strings"
	
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
    //"github.com/ethereum/go-ethereum"
    //"github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    //"github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/core/types"
	uniswapPairs "./uniswapPairs"
)

func main(){

	client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }

	var contractAddress = common.HexToAddress("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc")
	//uniswap_pair, err := uniswapPairs.NewUniswapPairs(contractAddress, client)

	if err != nil {
        log.Fatal(err)
    }

	query := &bind.FilterOpts{

		Start: uint64(2394201),
		End:   nil,
		Context: context.Background(),
	}

	//var contractFilterer bind.ContractFilterer
	swapFilterer, err := uniswapPairs.NewUniswapPairsFilterer(contractAddress, client)

	if err != nil {
        log.Fatal(err)
    }

	
	var addresses []common.Address
	//addresses = append(addresses, test_address)

	iterator, err := swapFilterer.FilterSwap(query, addresses, addresses)

	if err != nil {
        log.Fatal(err)
    }

	
	fmt.Println(iterator.Event)

	for  i := 1; i < 100000; i++{ 

		if i % 100 == 0{
			fmt.Println(i)
		}
		
		if(iterator.Event != nil){
			fmt.Println("Passes not nil")
			data, err := swapFilterer.ParseSwap(iterator.Event.Raw)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(data.Amount0In)
		}
		
	}
}