package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "strings"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"

    uniswapPairs "./uniswapPairs"
)

func main() {
    client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }

    var contractAddress = common.HexToAddress("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc")
	uniswap_pair, err := uniswapPairs.NewUniswapPairs(contractAddress, client)

    query := ethereum.FilterQuery{
        FromBlock: big.NewInt(2394201),
        ToBlock:   big.NewInt(2394220),
        Addresses: []common.Address{
            contractAddress,
        },
    }

    logs, err := client.FilterLogs(context.Background(), query)
    if err != nil {
        log.Fatal(err)
    }

    contractAbi, err := abi.JSON(strings.NewReader(string(uniswapPairs.UniswapPairsABI)))
    if err != nil {
        log.Fatal(err)
    }

    for _, vLog := range logs {
        fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
        fmt.Println(vLog.BlockNumber)     // 2394201
        fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

        //event := new(uniswapPairs.UniswapPairsSwap)
        event, err := uniswap_pair.contract.UnpackLogs("Swap", vLog.Data)
		fmt.Println(event[0])
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(event.Amount0In)   // foo
        // bar

        var topics [4]string
        for i := range vLog.Topics {
            topics[i] = vLog.Topics[i].Hex()
        }

        fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
    }

    eventSignature := []byte("ItemSet(bytes32,bytes32)")
    hash := crypto.Keccak256Hash(eventSignature)
    fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}