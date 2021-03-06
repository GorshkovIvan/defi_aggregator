package db

import "testing"

func TestGetAaveTickers(t *testing.T) {
	tickers := GetAaveTickers()
	//want := regexp.MustCompile([TUSD YFI BAT MANA REP UNI WBTC REN BUSD LINK sUSD DAI AAVE LEND MKR USDC SNX USDT KNC ZRX ETH ENJ])

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