package db

import "sort"

type Record struct {
	Pair    string  `json:"pair"`
	Amount  float32 `json:"amount"`
	Pool_sz float32 `json:"pool_sz"`
}

func NewRecord(pair string, amount float32, pool_sz float32) Record {
	return Record{pair, amount, pool_sz}
}

type CurrencyInputData struct {
	Pair   string  `default0:"ETH/DAI" json:"backend_pair"`
	Amount float32 `default0:"123" json:"backend_amount"`
	Yield  float32 `default0:"0.05" json:"backend_yield"`
}

func NewCurrencyInputData() CurrencyInputData {
	currencyinputdata := CurrencyInputData{}
	currencyinputdata.Pair = "ETH/DAI"
	currencyinputdata.Amount = 420.69
	currencyinputdata.Yield = 0.08
	return currencyinputdata
}

func NewCurrencyInputDataAct(pair string, amount float32, yield float32) CurrencyInputData {
	currencyinputdata := CurrencyInputData{}
	currencyinputdata.Pair = pair
	currencyinputdata.Amount = amount
	currencyinputdata.Yield = yield
	return currencyinputdata
}

type Database struct {
	contents []Record
	// currencyinputdata[0] = ETH/DAI [1] = DAI USDC
	currencyinputdata []CurrencyInputData // this will store the LATEST currency pair info
}

func New() Database {
	contents := make([]Record, 0)
	currencyinputdata := make([]CurrencyInputData, 0)

	// append being moved from HERE
	return Database{contents, currencyinputdata}
}
func (database *Database) AddRecord(r Record) {
	database.contents = append(database.contents, r)
}
func (database *Database) GetRecords() []Record {
	return database.contents
}

func (database *Database) AddRecordfromAPI() {
	//c CurrencyInputData = NewCurrencyInputData()
	database.currencyinputdata = append(database.currencyinputdata, NewCurrencyInputData())
}

func (database *Database) AddRecordfromAPI2(pair string, amount float32, yield float32) {
	//c CurrencyInputData = NewCurrencyInputData()
	database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{pair, amount, yield})
}

func (database *Database) GetCurrencyInputData() []CurrencyInputData {
	return database.currencyinputdata
}

func (database *Database) RankBestCurrencies() {

	sort.Slice(database.currencyinputdata, func(i, j int) bool {
		return database.currencyinputdata[i].Yield > database.currencyinputdata[j].Yield
	})
}
