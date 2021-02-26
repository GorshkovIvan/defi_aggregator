package db

import "testing"

func TestNewRecord(t *testing.T) {
	new_record := NewRecord("DAI/ETH", 12.34, 100.0)

	if new_record.Pair != "DAI/ETH" {
		t.Errorf("fail!")
	}

	if new_record.Amount != 12.34 {
		t.Errorf("fail!")
	}

	if new_record.Pool_sz != 100.0 {
		t.Errorf("fail!")
	}
}

func TestNewCurrencyInputData(t *testing.T) {
	new_currencyinputdata := NewCurrencyInputData()

	if new_currencyinputdata.Pair != "ETH/DAI" {
		t.Errorf("fail!")
	}

	if new_currencyinputdata.Amount != 420.69 {
		t.Errorf("fail!")
	}

	if new_currencyinputdata.Yield != 0.08 {
		t.Errorf("fail!")
	}
}

func TestNewCurrencyInputDataAct(t *testing.T) {
	new_currencyinputdataact := NewCurrencyInputDataAct("ETH/DAI", 10.0, 0.10)

	if new_currencyinputdataact.Pair != "ETH/DAI" {
		t.Errorf("fail!")
	}

	if new_currencyinputdataact.Amount != 10.0 {
		t.Errorf("fail!")
	}

	if new_currencyinputdataact.Yield != 0.10 {
		t.Errorf("fail!")
	}
}

func TestNew(t *testing.T) {
	new_database := New()

	if new_database.ownstartingportfolio == nil {
		t.Errorf("fail!")
	}

	if new_database.currencyinputdata == nil {
		t.Errorf("fail!")
	}
}

func TestAddRecord(t *testing.T) {
	new_database := New()
	new_record := NewRecord("DAI/ETH", 12.34, 100.0)
	new_database.AddRecord(new_record)

	if new_database.ownstartingportfolio[0] != new_record {
		t.Errorf("fail!")
	}
}

func TestGetRecords(t *testing.T) {
	new_database := New()
	new_record := NewRecord("DAI/ETH", 12.34, 100.0)
	new_database.AddRecord(new_record)
	//records := new_database.GetRecords()
	records := new_database.GetOptimisedPortfolio()

	if records[0] != new_record {
		t.Errorf("fail!")
	}
}

func TestAddRecordfromAPI(t *testing.T) {
	new_database := New()
	new_database.currencyinputdata = nil
	new_database.AddRecordfromAPI()

	if new_database.currencyinputdata == nil {
		t.Errorf("fail!")
	}
}

func TestAddRecordfromAPI2(t *testing.T) {
	new_database := New()
	new_database.currencyinputdata = nil
	new_database.AddRecordfromAPI2("DAI/ETH", 12.34, 0.11)

	if new_database.currencyinputdata[0].Pair != "DAI/ETH" {
		t.Errorf("fail!")
	}

	if new_database.currencyinputdata[0].Amount != 12.34 {
		t.Errorf("fail!")
	}

	if new_database.currencyinputdata[0].Yield != 0.11 {
		t.Errorf("fail!")
	}
}

func TestGetCurrencyInputData(t *testing.T) {
	new_database := New()
	new_database.currencyinputdata = nil
	new_database.AddRecordfromAPI2("DAI/ETH", 12.34, 0.11)
	currency_input_data := new_database.GetCurrencyInputData()

	if currency_input_data[0].Pair != "DAI/ETH" {
		t.Errorf("fail!")
	}

	if currency_input_data[0].Amount != 12.34 {
		t.Errorf("fail!")
	}

	if currency_input_data[0].Yield != 0.11 {
		t.Errorf("fail!")
	}
}

func TestRankBestCurrencies(t *testing.T) {
	new_database := New()
	new_database.AddRecordfromAPI2("DAI/ETH", 12.34, 0.05)
	new_database.AddRecordfromAPI2("DAI/USDC", 24.34, 0.13)
	new_database.AddRecordfromAPI2("DAI/BTC", 38.34, 0.01)
	new_database.RankBestCurrencies()
	currency_input_data := new_database.GetCurrencyInputData()

	if currency_input_data[0].Yield != 0.13 {
		t.Errorf("fail!")
	}

	if currency_input_data[1].Yield != 0.05 {
		t.Errorf("fail!")
	}

	if currency_input_data[2].Yield != 0.01 {
		t.Errorf("fail!")
	}
}
