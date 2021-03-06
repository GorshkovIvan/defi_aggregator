package db

import (
	"testing"
)

func TestGetAllData(t *testing.T) {
	database := New()

	database.AddRecordfromAPI()

	// Test 1
	if len(database.historicalcurrencydata) == 0 {
		t.Errorf("Unsufficient data downloaded!")
	}

	// Test 2
	var teststringarray []string
	var lengtharrayDates []int
	var lengtharrayPrices []int

	for i := 0; i < len(database.historicalcurrencydata); i++ {

		teststringarray = append(teststringarray, database.historicalcurrencydata[i].Ticker)
		lengtharrayDates = append(lengtharrayDates, len(database.historicalcurrencydata[i].Date))
		lengtharrayPrices = append(lengtharrayPrices, len(database.historicalcurrencydata[i].Price))

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

	var countBalancerdatapoints int
	var countUniswapdatapoints int
	var countAavedatapoints int

	for i := 0; i < len(database.currencyinputdata); i++ {
		if database.currencyinputdata[i].Pool == "Uniswap" {
			countUniswapdatapoints++
		}

		if database.currencyinputdata[i].Pool == "Aave" {
			countAavedatapoints++
		}

		if database.currencyinputdata[i].Pool == "Balancer" {
			countBalancerdatapoints++
		}
	}

	if countUniswapdatapoints == 0 {
		t.Errorf("Error: no data appended from Uniswap")
	}

	if countAavedatapoints == 0 {
		t.Errorf("Error: no data appended from Aave")
	}

	if countBalancerdatapoints == 0 {
		t.Errorf("Error: no data appended from Balancer")
	}

}
