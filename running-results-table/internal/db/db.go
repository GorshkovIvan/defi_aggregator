package db

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
	currencyinputdata.Pair = "ETH/BTC"
	currencyinputdata.Amount = 420.69
	currencyinputdata.Yield = 0.08
	return currencyinputdata
}

type Database struct {
	contents          []Record
	currencyinputdata []CurrencyInputData
}

func New() Database {
	contents := make([]Record, 0)
	currencyinputdata := make([]CurrencyInputData, 0)

	currencyinputdata = append(currencyinputdata, NewCurrencyInputData())

	return Database{contents, currencyinputdata}
}
func (database *Database) AddRecord(r Record) {
	database.contents = append(database.contents, r)
}
func (database *Database) GetRecords() []Record {
	return database.contents
}

func (database *Database) GetCurrencyInputData() []CurrencyInputData {
	return database.currencyinputdata
}
