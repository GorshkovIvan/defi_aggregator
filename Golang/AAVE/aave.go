package main

import (
    "fmt"
    "log"

    "math"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    data_provider "./protocol_data_provider"
)

func convert_number(n *big.Int) *big.Float{
  x := new(big.Float)
  x.SetString(n.String())
  value := new(big.Float).Quo(x, big.NewFloat(math.Pow10(6)))
  return value
}

func convert_rate(n *big.Int) *big.Float{
  x := new(big.Float)
  x.SetString(n.String())
  value := new(big.Float).Quo(x, big.NewFloat(math.Pow10(25)))
  return value
}

func main() {
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/7ee001f2b684469faff12e0485f3f977")
    if err != nil {
        log.Fatal(err)
    }

    // AAVE data contract address
    tokenAddress := common.HexToAddress("0x057835Ad21a177dbdd3090bB1CAE03EaCF78Fc6d")
    instance, err := data_provider.NewStore(tokenAddress, client)

    if err != nil {
        log.Fatal(err)
    }

    tokens, err := instance.GetAllReservesTokens(&bind.CallOpts{})

    if err != nil {
    	log.Fatal(err)
    	}


    s := fmt.Sprintf("%s", tokens[0])
    s = string(s)
    base_len := len(s)

    for i := 0; i < len(tokens) - 1; i++ {

      s := fmt.Sprintf("%s", tokens[i])
      s = string(s)

      if(len(s) < base_len){
        fmt.Println(s[1:4])
        s = s[7:len(s)-1]
      }else{
        fmt.Println(s[1:5])
        s = s[6:len(s)-1]
      }

      tokenAddress = common.HexToAddress(s)
      reserve_data, err := instance.GetReserveData(&bind.CallOpts{}, tokenAddress)

      if err != nil {
        log.Fatal(err)
        }

      fmt.Printf("Available Liquidity: %f \n", convert_number(reserve_data.AvailableLiquidity))
      fmt.Printf("Total Stable Debt: %f \n", convert_number(reserve_data.TotalStableDebt))
      fmt.Printf("Total Variable Debt: %f \n", convert_number(reserve_data.TotalVariableDebt))
      fmt.Printf("Liquidity Rate: %f \n", convert_rate(reserve_data.LiquidityRate))
      fmt.Printf("Variable Borrow Rate: %f \n", convert_rate(reserve_data.VariableBorrowRate))
      fmt.Printf("Stable Borrow Rate: %f \n", convert_rate(reserve_data.StableBorrowRate))
      fmt.Printf("Average Stable Borrow Rate: %f\n", convert_rate(reserve_data.AverageStableBorrowRate))
      fmt.Printf("Liquidity Index: %f\n", convert_rate(reserve_data.LiquidityIndex))
      fmt.Printf("Variable Borrow Index: %f\n", convert_rate(reserve_data.VariableBorrowIndex))
      fmt.Printf("\n")
      fmt.Printf("\n")

    }


}
