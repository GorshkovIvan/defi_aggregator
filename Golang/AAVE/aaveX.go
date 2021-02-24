package main

import (
    "fmt"
    "log"

    "math"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

     data_provider "pusher/defi_aggregator/Golang/AAVE/protocol_data_provider"
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


//func convert_rate_to_int(n *big.Int) int64{
//  var smallnum, _ = new(big.Int).SetString("2188824200011112223", 10)
//  num := smallnum.Int64()
//  return num
//}


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


    for i := 0; i < 2; i++ { // len(tokens) - 1
      s := fmt.Sprintf("%s", tokens[i])
      s = string(s)

      if(len(s) < base_len){
        //fmt.Println(s[1:4])
        s = s[7:len(s)-1]
      }else{
        //fmt.Println(s[1:5])
        s = s[6:len(s)-1]
      }

      str := fmt.Sprintf("%s", tokens[i]);
      str = string(str);
      str = str[1:5];
      // fmt.Printf(str);
      // fmt.Printf("\n");

      tokenAddress = common.HexToAddress(s)
      reserve_data, err := instance.GetReserveData(&bind.CallOpts{}, tokenAddress)
      if err != nil {log.Fatal(err)}

      var ROI_hist float32;
      ROI_est := convert_number(reserve_data.LiquidityRate);
      ROI_est_small,y := ROI_est.Float64()
      ROI_est_small = ROI_est_small / 1000000000000000000000.0
      ROI_hist = 0.0;
      if (y < 0){} // fmt.Printf("X",y);

      // if err2 != nil {log.Fatal(err2)}
      // ROI_est = 0.1;
      // ROI_est = 

      fmt.Printf(str);
      fmt.Printf(" ROI est: %f \n", ROI_est_small)
      fmt.Printf(" | ")
      fmt.Printf("ROI hist: ", ROI_hist);
      fmt.Printf("\n")

      // Lending ROI

      // Historical price data
      // Historical yield data
          // calculate historical yield volatility
          // apply it to current rate to estimate next 30 days
          // output lo-hi yields
      
      // Non-stable ROI

      // fmt.Printf("ROI historical: ", ROI_hist);


    }
}
