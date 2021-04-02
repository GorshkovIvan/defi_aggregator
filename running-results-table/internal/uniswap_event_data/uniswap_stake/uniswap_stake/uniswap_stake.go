package main

import (
    "context"
    "fmt"
    "log"
    "time"
    "crypto/ecdsa"
    "math"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    DAI "./DAI"
    uniswapRouter2 "./uniswapRouter2"
)

func main() {

    //client, err := ethclient.Dial("https://mainnet.infura.io/v3/7ee001f2b684469faff12e0485f3f977")
    client, err := ethclient.Dial("http://localhost:8888")
    if err != nil {
        log.Fatal(err)
    }

	var Router02Address = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	var DaiAddress  = common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
    var WETHAddress = common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	//var MyAddress = common.HexToAddress("0xD75cA8E887c227c2F87D376F76967416DC7E8911")
    // Init uniswap

    fmt.Println("Initialising uniswap")
	uniswap, err := uniswapRouter2.NewUniswapRouter2(Router02Address, client)
	if err != nil {
        fmt.Println("Contract failed to initialise")
        log.Fatal(err)
    }
	fmt.Println(uniswap.WETH(&bind.CallOpts{}))

    // Init DAI

    dai, err := DAI.NewDAI(DaiAddress, client)
	if err != nil {
        log.Fatal(err)
    }
    //unlockedAddres := common.HexToAddress("0x447a9652221f46471a2323B98B73911cda58FD8A")

    privateKey, err := crypto.HexToECDSA("4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d")
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()

    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    //gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.From = fromAddress
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(1000000000000000000)     // in wei
    auth.GasLimit = uint64(6721975) // in units
    auth.GasPrice = big.NewInt(20000000000) 
    deadline := big.NewInt(int64((time.Now().UTC().UnixNano() / 1e6)/1000) + 60 * 20)

    // SwapETHForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int)


    tokens := []common.Address{WETHAddress, DaiAddress}
    amountToExchange := big.NewInt(30)
    fmt.Println("Swaping Eth to DAI...")
    swap_transaction, err := uniswap.SwapExactETHForTokens(auth, amountToExchange, tokens, fromAddress, deadline)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Swaped! Transaction data:")
    fmt.Println(swap_transaction)
    // Liquidity Provision
    
    nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    auth.Nonce = big.NewInt(int64(nonce))
    
    // Balance after swap

    balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
    fbalance := new(big.Float)
    fbalance.SetString(balance.String())
    ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
    fmt.Println("Eth Balance after Swap:")
    fmt.Println(ethValue)
    // Dai Approval
    auth.Value = big.NewInt(0)
    amountTokenDesired := big.NewInt(1700)
    dai.Approve(auth, Router02Address, amountTokenDesired)

    // Staking
    fmt.Println("Staking...")
    nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(1000000000000000000)
    amountTokenMin := big.NewInt(1200)
    amountETHMin := big.NewInt(1)
    transaction, err := uniswap.AddLiquidityETH(auth, DaiAddress, amountTokenDesired, amountTokenMin, amountETHMin, fromAddress, deadline)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Invested! Transaction data:")
    fmt.Println(transaction)
    //opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int
}