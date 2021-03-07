package db

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/machinebox/graphql"
)

// Holds a list of Aave tickers
type AaveSymbols struct {
	Symbols []AaveSymbol `json:"reserves"`
}

type AaveSymbol struct {
	Symbol string `json:"symbol"`
}

// Holds Curent Aave data
type AaveQuery struct {
	Reserve AaveData `json:"reserve"`
}

// Current Aave data
type AaveData struct {
	ID                 string `json:"id"`
	Symbol             string `json:"symbol"`
	LiquidityRate      string `json:"liquidityRate"`
	StableBorrowRate   string `json:"stableBorrowRate"`
	VariableBorrowRate string `json:"variableBorrowRate"`
	TotalBorrows       string `json:"totalBorrows"`
}

// Returns Currrent Data from Aave protocol
func getAaveCurrentData() (string, float32, float32, float32, float32) {

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

	totalBorrows, _ := strconv.ParseFloat(respAave.Reserve.TotalBorrows, 32)
	volume := float32(totalBorrows)
	stableBorrowRate, _ := strconv.ParseFloat(respAave.Reserve.StableBorrowRate, 32)
	interest := float32(stableBorrowRate)
	size := float32(69)
	volatility := float32(420)

	return respAave.Reserve.Symbol, size, volume, interest, volatility
}

// Returns all tickers from Aave
func GetAaveTickers() []string {

	clientAave := graphql.NewClient("https://api.thegraph.com/subgraphs/name/aave/protocol")
	ctx := context.Background()

	reqAaveListOfPools := graphql.NewRequest(`
		query
		{
			reserves{
		
				symbol

			}
		}
        `)

	reqAaveListOfPools.Header.Set("Cache-Control", "no-cache")

	var respAavePoolList AaveSymbols

	if err := clientAave.Run(ctx, reqAaveListOfPools, &respAavePoolList); err != nil {
		log.Fatal(err)
	}

	return tickersToString(respAavePoolList)

}

// Converts all tickers to strings
func tickersToString(tickers AaveSymbols) []string {

	var stringTokenList []string

	for i := 0; i < len(tickers.Symbols); i++ {

		stringTokenList = append(stringTokenList, tickers.Symbols[i].Symbol)

	}

	return stringTokenList

}

// Updates database with Uniswap historical data and current Aave data
func getAaveData(database *Database, uniswapreqdata UniswapInputStruct) {
	ctx := context.Background()
	var respUniswapTicker UniswapTickerQuery
	var respUniswapHist UniswapHistQuery
	AavePoolList := GetAaveTickers()
	var AaveFilteredTokenList []string
	

	//Updating historical data

	// Process received list token
	for i := 0; i < len(AavePoolList); i++ {

		if len(AavePoolList) > 1 {
			token0symbol := AavePoolList[i]
			token1symbol := token0symbol

			if isPoolPartOfFilter(token0symbol, token1symbol) {
				// Filter pools to allowed components (WETH, DAI, USDC, USDT)
				//var tokenqueue []string
				AaveFilteredTokenList = append(AaveFilteredTokenList, token0symbol)
				//tokenqueue = append(tokenqueue, token1symbol)

				//for j := 0; j < len(tokenqueue); j++ {
				// Check if database already has historical data
				if !isHistDataAlreadyDownloaded(token0symbol, database) { // tokenqueue[j]
					
					
					uniswapreqdata.reqUniswapIDFromTokenTicker.Var("ticker", token0symbol) // tokenqueue[j]

					if err := uniswapreqdata.clientUniswap.Run(ctx, uniswapreqdata.reqUniswapIDFromTokenTicker, &respUniswapTicker); err != nil {
						log.Fatal(err)
					}
					// Download historical data for each token for which data is missing
					if len(respUniswapTicker.IDsforticker) >= 1 {
					// request data from uniswap using this queried ticker
						uniswapreqdata.reqUniswapHist.Var("tokenid", setUniswapQueryIDForToken(token0symbol, respUniswapTicker.IDsforticker[0].ID)) // tokenqueue[j]

						fmt.Print("Querying historical data for: ")
						fmt.Print(token0symbol) // tokenqueue[j]
						if err := uniswapreqdata.clientUniswap.Run(ctx, uniswapreqdata.reqUniswapHist, &respUniswapHist); err != nil {
							log.Fatal(err)
						}

						fmt.Print("| returned days: ")
						fmt.Println(len(respUniswapHist.DailyTimeSeries))

						// if returned data - append it to database
						if len(respUniswapHist.DailyTimeSeries) > 0 {
						// Append to database
							database.historicalcurrencydata = append(database.historicalcurrencydata, NewHistoricalCurrencyDataFromRaw(token0symbol, respUniswapHist.DailyTimeSeries)) // tokenqueue[j]
						}
					}
					
				} // if historical data needs updating
				// } // tokenqueue loop ends

			}
		}
	}

	// Updating current data
	symbol, size, volume, interest, volatility := getAaveCurrentData()
	ROI := calculateROI(interest, 0, volume, volatility)
	database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{symbol, size, volume, interest, "Aave", volatility, ROI})

}


