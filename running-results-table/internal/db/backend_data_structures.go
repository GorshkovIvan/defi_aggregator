package db

type OwnPortfolioRecord struct {
	Token  string  `json:"token"`
	Amount float32 `json:"amount"`
	//	Pool_sz float32 `json:"pool_sz"`
}

// For adding own portfolio records
func NewOwnPortfolioRecord(token string, amount float32) OwnPortfolioRecord { // , pool_sz float32
	return OwnPortfolioRecord{token, amount} // , pool_sz
}

func NewHistoricalCurrencyDataFromRaw(token string, rawhistoricaldata []UniswapDaily) HistoricalCurrencyData {
	// TO DO: parsing
	historicaldata := HistoricalCurrencyData{}
	historicaldata.Ticker = token
	return historicaldata
}

func NewHistoricalCurrencyData() HistoricalCurrencyData {
	historicaldata := HistoricalCurrencyData{}
	//historicaldata.Date = append(historicaldata.Date, "Mon Jan 2 15:04:05 MST 2006")
	historicaldata.Date = append(historicaldata.Date, 1099999999999999)
	historicaldata.Price = append(historicaldata.Price, float32(420.69))
	historicaldata.Ticker = "ETH"

	return historicaldata
}

type HistoricalCurrencyData struct {
	Date   []int64   `default0:"10099999999999" json:"date"` // `default0:"Mon Jan 2 15:04:05 MST 2006" json:"date"`
	Price  []float32 `default0:"420.69" json:"price"`
	Ticker string    `default0:"ETH" json:"ticker"`
}

// Current
type CurrencyInputData struct {
	Pair        string  `default0:"ETH/DAI" json:"backend_pair"`
	PoolSize    float32 `default0:"420000.69" json:"backend_poolsize"`
	PoolVolume  float32 `default0:"4200.69" json:"backend_volume"`
	Yield       float32 `default0:"0.05" json:"backend_yield"`
	Pool        string  `default0:"Uniswap" json:"pool_source"`
	Volatility  float32 `default0:"-9.00" json:"volatility"`
	ROIestimate float32 `default0:"42.69%" json:"ROIestimate"`
}

func NewCurrencyInputData() CurrencyInputData {
	currencyinputdata := CurrencyInputData{}
	currencyinputdata.Pair = "ETH/DAI"
	currencyinputdata.PoolSize = 420000.69
	currencyinputdata.PoolVolume = 4200.69
	currencyinputdata.Yield = 0.08
	currencyinputdata.Pool = "Uniswap"
	currencyinputdata.Volatility = -0.09
	currencyinputdata.ROIestimate = 0.4269
	return currencyinputdata
}

func NewCurrencyInputDataAct(pair string, poolSz float32, poolVolume float32, yield float32, pool string, volatility float32, roi float32) CurrencyInputData {
	currencyinputdata := CurrencyInputData{}
	currencyinputdata.Pair = pair

	currencyinputdata.PoolSize = poolSz
	currencyinputdata.PoolVolume = poolVolume

	currencyinputdata.Yield = yield
	currencyinputdata.Pool = pool
	currencyinputdata.Volatility = volatility
	currencyinputdata.ROIestimate = roi
	return currencyinputdata
}

type Database struct {
	ownstartingportfolio   []OwnPortfolioRecord     // for portfolio optimisation table
	currencyinputdata      []CurrencyInputData      // LATEST currency pair info
	historicalcurrencydata []HistoricalCurrencyData // historical time series
}

func New() Database {
	ownstartingportfolio := make([]OwnPortfolioRecord, 0)
	currencyinputdata := make([]CurrencyInputData, 0)
	historicalcurrencydata := make([]HistoricalCurrencyData, 0)
	return Database{ownstartingportfolio, currencyinputdata, historicalcurrencydata}
}

// Add OWN PORTFOLIO data
func (database *Database) AddRecord(r OwnPortfolioRecord) {
	database.ownstartingportfolio = append(database.ownstartingportfolio, r)
}

func (database *Database) GetRecords() []OwnPortfolioRecord {
	return database.ownstartingportfolio
}

// Add CURRENT pair data
func (database *Database) AddRecordfromAPI2(pair string, poolSz float32, poolVolume float32, yield float32, pool string, volatility float32, ROIestimate float32) {
	database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{pair, poolSz, poolVolume, yield, pool, volatility, ROIestimate})
}

func (database *Database) GetCurrencyInputData() []CurrencyInputData {
	return database.currencyinputdata
}
