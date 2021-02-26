package db

import (
	"context"
	"log"
	"strconv"
	"fmt"
	"github.com/machinebox/graphql"
)

type AaveSymbols struct {
	Symbols []AaveSymbol `json:"reserves"`
}

type AaveSymbol struct {
	Symbol string `json:"symbol"`
}

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

func GetAaveData() (string, float32, float32, float32, float32) {

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

	return respAave.Reserve.Symbol, float, float2, float32(0.01), float32(0.001)
}

func GetAaveTickers ()(AaveSymbols){

	clientAave := graphql.NewClient("https://api.thegraph.com/subgraphs/name/aave/protocol")
	ctx := context.Background()

	reqAave := graphql.NewRequest(`
		query
		{
			reserves{
		
				symbol

			}
		}
        `)

	reqAave.Header.Set("Cache-Control", "no-cache")

	var respAave AaveSymbols
	
	if err := clientAave.Run(ctx, reqAave, &respAave); err != nil {
		log.Fatal(err)
				}
	
	return respAave

}

func getPriceDataFromUniswap(required_tickers []string, available_tickers AaveSymbols){

	var AaveFilteredTokenList []string
	for i := 0; i < len(required_tickers); i++{
		if(isTickerInAave(required_tickers[i], available_tickers)){
			AaveFilteredTokenList = append(AaveFilteredTokenList, required_tickers[i])

		}

	}

	fmt.Println(AaveFilteredTokenList)

}

func isTickerInAave(ticker string, available_tickers AaveSymbols) (bool){

	for i := 0; i < len(available_tickers.Symbols); i++ {
		if (available_tickers.Symbols[i].Symbol == ticker){
			return true
		}
	}

	return false

}
/*
func main(){

	symbols := GetAaveTickers()
	var required_tickers []string
	required_tickers = append(required_tickers, "USDT")
	required_tickers = append(required_tickers, "USDC")
	getPriceDataFromUniswap(required_tickers, symbols)

}
*/