package db

import "testing"

func TestNewOptimisedPortfolio(t *testing.T) {
	ownPortfolioRecord := OwnPortfolioRecord{"ETH", 100}
	var ownPortfolioRecordArray [] OwnPortfolioRecord
	ownPortfolioRecordArray = append(ownPortfolioRecordArray, ownPortfolioRecord)
	newOptimisedPortfolio := NewOptimisedPortfolio(ownPortfolioRecordArray) // returns array of portfolios

	if newOptimisedPortfolio[0].TokenOrPair != "ETH" {
		t.Errorf("Token error!")
	}

	if newOptimisedPortfolio[0].Pool != "Uniswap" {
		t.Errorf("Pool error!")
	}

	if newOptimisedPortfolio[0].Amount != 100 {
		t.Errorf("Pool error!")
	}

	if newOptimisedPortfolio[0].PercentageOfPortfolio != 0.420 {
		t.Errorf("Percentage of Portfolio error!")
	}

	if newOptimisedPortfolio[0].ROIestimate != 0.069 {
		t.Errorf("ROI error!")
	}

}

func TestNewOptimisedPortfolioWithInputLengthZero(t *testing.T) {
	var ownPortfolioRecordArray [] OwnPortfolioRecord
	newOptimisedPortfolio := NewOptimisedPortfolio(ownPortfolioRecordArray) // returns array of portfolios

	if newOptimisedPortfolio[0].TokenOrPair != "USD" {
		t.Errorf("Token error!")
	}

	if newOptimisedPortfolio[0].Pool != "Uniswap" {
		t.Errorf("Pool error!")
	}

	if newOptimisedPortfolio[0].Amount != 100 {
		t.Errorf("Pool error!")
	}

	if newOptimisedPortfolio[0].PercentageOfPortfolio != 1 {
		t.Errorf("Percentage of Portfolio error!")
	}

	if newOptimisedPortfolio[0].ROIestimate != 0.0125 {
		t.Errorf("ROI error!")
	}

}

func TestNewOwnPortfolioRecord(t *testing.T) {
	ownPortfolioRecord := NewOwnPortfolioRecord("ETH", 100)

	if ownPortfolioRecord.Token != "ETH" {
		t.Errorf("Failed to save token name into OwnPortfolioRecord!")
	}

	if ownPortfolioRecord.Amount != 100 {
		t.Errorf("Failed to save token amount into OwnPortfolioRecord!")
	}

}

func TestNewHistoricalCurrencyDataFromRaw(t *testing.T) {
	var rawHistoricalData []UniswapDaily
	historicalCurrencyData := NewHistoricalCurrencyDataFromRaw("ETH", rawHistoricalData)

	if historicalCurrencyData.Ticker != "ETH" {
		t.Errorf("Failed to create new historical currency data from raw data")
	}

}

func TestNewHistoricalCurrencyData(t *testing.T) {
	historicalData := NewHistoricalCurrencyData()

	if historicalData.Date[0] != 1099999999999999 {
		t.Errorf("Failed to create Historical Currency Data with the correct Date")
	}
}

func TestNewCurrencyInputData(t *testing.T) {
	new_currencyinputdata := NewCurrencyInputData()

	if new_currencyinputdata.Pair != "ETH/DAI" {
		t.Errorf("fail!")
	}

	if new_currencyinputdata.PoolSize != 420000.69 {
		t.Errorf("fail!")
	}

	if new_currencyinputdata.PoolVolume != 4200.69 {
		t.Errorf("fail!")
	}

	if new_currencyinputdata.Yield != 0.08 {
		t.Errorf("fail!")
	}

	/*if new_currencyinputdata.Pool != "Uniswap" {
		t.Errorf("fail!")
	}*/

}