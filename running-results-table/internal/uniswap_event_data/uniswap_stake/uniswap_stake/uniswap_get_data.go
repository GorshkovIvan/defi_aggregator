package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
    //"strings"
	
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum"
    //"github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    //"github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
	uniswapPairs "./uniswapPairs"
)


//type Event interface{}

func main(){

	client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }

	var contractAddress = common.HexToAddress("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc")

/*
	query := &bind.FilterOpts{

  		Start: uint64(12035100),
  		End:   nil,
  		Context: context.Background(),
}
*/
	uniswap_pair, err := uniswapPairs.NewUniswapPairs(contractAddress, client)
	// NEW

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(12035100),
		ToBlock:   nil,
		Addresses: []common.Address{
		contractAddress,
		},
	}

	//l, err := client.FilterLogs(context.Background(), query)
	logs := make(chan types.Log)

	if err != nil {
		log.Fatal(err)
	}

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		log.Fatal(err)
	}

	//SwapIterator := uniswapPairs.UniswapPairsSwapIterator{contract: uniswap_pair, event: "Swap", logs: logs, sub: sub}
	var SwapIterator uniswapPairs.UniswapPairsSwapIterator
	SwapIterator.contract = uniswap_pair

// OLD

//var addresses []common.Address
//addresses = append(addresses, contractAddress)
//SwapIterator, err = uniswap_pair.FilterSwap(query, nil, nil)
/*
if err != nil {
	log.Fatal(err)
}
*/
var cont bool

for { cont = SwapIterator.Next()

	if(SwapIterator.Event != nil){
		fmt.Println(uniswap_pair.ParseSwap(SwapIterator.Event).Amount0In)
	}
	if(!cont){
		break
	}
	
}
/*
//contractAbi, err := abi.JSON(strings.NewReader(string(uniswapPairs.UniswapPairsABI)))
if err != nil {
	log.Fatal(err)
}

for _, vLog := range logs {
	fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
	fmt.Println(vLog.BlockNumber)     // 2394201
	fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

	var event *uniswapPairs.UniswapPairsSwap
	//var event map[string]interface{}
	//event, err := contractAbi.Unpack("Swap", vLog.Data)
	event, err := uniswap_pair.ParseSwap(vLog)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(event)
	fmt.Println(event.Amount0In)   // foo
	//fmt.Println(string(event.Value[:])) // bar

	var topics [4]string
	for i := range vLog.Topics {
		topics[i] = vLog.Topics[i].Hex()
	}

	fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}

	eventSignature := []byte("Swap")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex())
*/
}

