package main

import (
    "fmt"
    "log"
 //   "math"
 //   "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    stakingrewards "./StakingRewards" // for demo
)

func main() {
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/7ee001f2b684469faff12e0485f3f977")
    if err != nil {
        log.Fatal(err)
    }

    // USDT Uniswap Pool Address
    tokenAddress := common.HexToAddress("0x6C3e4cb2E96B01F4b866965A91ed4437839A121a")
    instance, err := stakingrewards.NewStakingrewards(tokenAddress, client)
    if err != nil {
        log.Fatal(err)
    }
    
    reward, err := instance.RewardPerToken(&bind.CallOpts{})
    if err != nil {
    	log.Fatal(err)
    	}
    	fmt.Printf("Reward: %s\n", reward)  
    	
/*
    address := common.HexToAddress("0x28e71d0b7f7f29106a1be2a5b289cab331e7b56f")
    bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
    if err != nil {
        log.Fatal(err)
    } 

    name, err := instance.Name(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("name: %s\n", name)  

    symbol, err := instance.Symbol(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }

    decimals, err := instance.Decimals(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("name: %s\n", name)         // "name: Golem Network"
    fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
    fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

    fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

    fbal := new(big.Float)
    fbal.SetString(bal.String())
    value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

    fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
    */
}

			