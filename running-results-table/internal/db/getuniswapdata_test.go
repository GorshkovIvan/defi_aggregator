package db

import (
	"testing"

	"github.com/machinebox/graphql"
)

func TestUniswapDataDownload(t *testing.T) {
	database := New()
	clientUniswap := graphql.NewClient("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2")

	reqUniswapHist := graphql.NewRequest(`
				query ($tokenid:String!){
						tokenDayDatas(first: 30 orderBy: date, orderDirection: asc,
						 where: {
						   token:$tokenid
						 }
						) {
						   date
						   priceUSD
						   token{
							   id
							   symbol
						   }
						}
				  }
			`)

	reqUniswapIDFromTokenTicker := graphql.NewRequest(`
						query ($ticker:String!){
							tokens(where:{symbol:$ticker})
							{
								id
								symbol
							}
						}
			`)

	reqUniswapIDFromTokenTicker.Header.Set("Cache-Control", "no-cache")
	reqUniswapHist.Header.Set("Cache-Control", "no-cache")

	// Run whole thing
	U := UniswapInputStruct{clientUniswap, reqUniswapIDFromTokenTicker, reqUniswapHist}
	getUniswapData(&database, U)

	// fmt.Println(len(database.historicalcurrencydata))
	// fmt.Println(len(database.historicalcurrencydata[0].Ticker))

	// Test 1
	if len(database.historicalcurrencydata) == 0 {
		t.Errorf("Unsufficient data downloaded!")
	}

	// Test 2

	// Test 3
	var teststringarray []string

	for i := 0; i < len(database.historicalcurrencydata); i++ {
		teststringarray = append(teststringarray, database.historicalcurrencydata[i].Ticker)
		if len(database.historicalcurrencydata[i].Ticker) == 0 {
			t.Errorf("No data pulled")
		}
		if len(database.historicalcurrencydata[i].Date) == 0 {
			t.Errorf("No data pulled")
		}
		if len(database.historicalcurrencydata[i].Price) == 0 {
			t.Errorf("No data pulled")
		}
		if len(database.historicalcurrencydata[i].Date) > 100 {
			t.Errorf("Too much data pulled")
		}
		if len(database.historicalcurrencydata[i].Price) > 100 {
			t.Errorf("Too much data pulled")
		}

	}

	if !stringInSlice("ETH", teststringarray) && !stringInSlice("WETH", teststringarray) {
		t.Errorf("ETH data missing!")
	}
	if !stringInSlice("DAI", teststringarray) {
		t.Errorf("DAI data missing!")
	}
	if !stringInSlice("USDC", teststringarray) {
		t.Errorf("USDC data missing!")
	}
	if !stringInSlice("USDT", teststringarray) {
		t.Errorf("USDT data missing!")
	}
	if !stringInSlice("WBTC", teststringarray) {
		t.Errorf("WBTC data missing!")
	}

}
