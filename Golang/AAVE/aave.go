package main

import (
    "fmt"
    "log"

    //"math"
    //"math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    variable_interest "./interest" // for demo
)

func main() {
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/7ee001f2b684469faff12e0485f3f977")
    if err != nil {
        log.Fatal(err)
    }

    // USDT Uniswap Pool Address
    tokenAddress := common.HexToAddress("0x057835Ad21a177dbdd3090bB1CAE03EaCF78Fc6d")
    instance, err := variable_interest.NewStore(tokenAddress, client)

    if err != nil {
        log.Fatal(err)
    }

    //tokenAddress = common.HexToAddress("0xfC1E690f61EFd961294b3e1Ce3313fBD8aa4f85d")
    //_,_,_,_,interest,_,_,_,_,_,_,_ := instance.GetReserveData(&bind.CallOpts{},tokenAddress)
    //interest, err:= instance.GetReserveData(&bind.CallOpts{},tokenAddress)
    tokens, err := instance.GetAllATokens(&bind.CallOpts{})

    if err != nil {
    	log.Fatal(err)
    	}

      fmt.Printf("Reward: %s\n", tokens[0])
    	fmt.Printf("Reward: %s\n", tokens[1])

      tokenAddress = common.HexToAddress("0x9ff58f4fFB29fA2266Ab25e75e2A8b3503311656")

      //reserve_data, err := json.Marshal(instance.GetReserveData(&bind.CallOpts{}, tokenAddress))
      /*
      reserve_data := new(struct {
    		AvailableLiquidity      *big.Int
    		TotalStableDebt         *big.Int
    		TotalVariableDebt       *big.Int
    		LiquidityRate           *big.Int
    		VariableBorrowRate      *big.Int
    		StableBorrowRate        *big.Int
    		AverageStableBorrowRate *big.Int
    		LiquidityIndex          *big.Int
    		VariableBorrowIndex     *big.Int
    		LastUpdateTimestamp     *big.Int
    	})
*/
      reserve_data, err := instance.GetReserveData(&bind.CallOpts{}, tokenAddress)

      if err != nil {
      	log.Fatal(err)
      	}

      fmt.Printf("%+v\n", reserve_data)
      //fmt.Printf(reserve_data)
      //fmt.Printf("Reward: %s\n", reserve_data)

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
