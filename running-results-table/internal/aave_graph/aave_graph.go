package aave_graph

import (
  "context"
  "log"
	"strconv"
	"github.com/machinebox/graphql"
)

type AaveQuery struct {
	Reserve AaveData `json:"reserve"`
}

type AaveData struct {
	ID                 string `json:"id"`
	Symbol             string `json:"symbol"`
	LiquidityRate      string `json:"liquidityRate"`
	StableBorrowRate   string `json:"stableBorrowRate"`
	VariableBorrowRate string `json:"variableBorrowRate"`
	TotalBorrows       string `json:"totalBorrows"`
}

func GetData() (string, float32, float32){

  clientAave := graphql.NewClient("https://api.thegraph.com/subgraphs/name/aave/protocol")
  ctx := context.Background()
  reqAave := graphql.NewRequest(`
    query($address: ID!)
    {
        reserve(id: $address){
        id
        symbol
        liquidityRate
        stableBorrowRate
        variableBorrowRate
        totalBorrows
      }
      }
        `)

  reqAave.Var("address", "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
  reqAave.Header.Set("Cache-Control", "no-cache")

  var respAave AaveQuery

  if err := clientAave.Run(ctx, reqAave, &respAave); err != nil {
    log.Fatal(err)
  }

  value, _ := strconv.ParseFloat(respAave.Reserve.TotalBorrows, 32)
  float := float32(value)
  float2 := float32(-0.0125)

  return respAave.Reserve.Symbol, float, float2
}
