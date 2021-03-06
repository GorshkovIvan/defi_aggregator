package db

import (
	"testing"

	"github.com/machinebox/graphql"
)

func TestUniswapDataOutput(t *testing.T) {
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

	// Test 1
	if len(database.historicalcurrencydata) == 0 {
		t.Errorf("Unsufficient data downloaded!")
	}

	// Test 2
	if len(database.historicalcurrencydata[0].Ticker) == 0 {
		t.Errorf("Ticker ")
	}
}
