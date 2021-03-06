package db


import (
	"testing"

	"github.com/machinebox/graphql"
	
)


func TestGetAaveTickers(t *testing.T) {
	tickers := GetAaveTickers()

	if len(tickers) < 1 {
        t.Fatalf(`TestGetAaveTickers() = %q, error, there should be some tickers`, tickers)
    }
}

func TesttickersToString(t *testing.T) {

	var dai AaveSymbol
	var weth AaveSymbol
	dai.Symbol = "DAI"
	weth.Symbol = "WETH"
	var tickers AaveSymbols
	tickers.Symbols[0] = dai
	tickers.Symbols[1] = weth
	tickers_to_string := tickersToString(tickers)
	//want := regexp.MustCompile([TUSD YFI BAT MANA REP UNI WBTC REN BUSD LINK sUSD DAI AAVE LEND MKR USDC SNX USDT KNC ZRX ETH ENJ])

	if tickers_to_string[0] != "DAI" || tickers_to_string[1] != "WETH" {
        t.Fatalf(`tickersToString(...) = %q, error, should return DAI and WETH strings in an array`, tickers_to_string)
    }
}

func TestgetAaveCurrentData(t *testing.T){

	symbol, size, volume, interest, volatility := getAaveCurrentData()

	if symbol == "" || size < 0 || volume < 0|| interest < 0 || volatility < 0{
		t.Fatalf(`One of the values is missing`)
    }
	
}

func TestgetAaveData(t *testing.T){
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

	U := UniswapInputStruct{clientUniswap, reqUniswapIDFromTokenTicker, reqUniswapHist}

	getAaveData(&database, U)

	if len(database.historicalcurrencydata) == 0 {
		t.Errorf("fail!")
	}

	// Test 2
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

	// Test 3

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
		if len(database.historicalcurrencydata[0].Ticker) == 0 {
			t.Errorf("Ticker ")
		}
	}

	// Test 4

	
	
}
