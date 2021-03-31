package db

import (
	"fmt"
	"strconv"
)

type OwnPortfolioRecord struct {
	Token  string  `json:"token"`
	Amount float32 `json:"amount"`
}

// Output of OwnPortfolioRecord
type OptimisedPortfolioRecord struct {
	TokenOrPair           string  `json:"tokenorpair"`
	Pool                  string  `json:"pool"`
	Amount                float32 `json:"amount"`
	PercentageOfPortfolio float32 `json:"percentageofportfolio"`
	ROI_raw_est           float32 `json:"roi_estimate"`
	Risksetting           float32 `json:"risk_setting"`
}

func NewOptimisedPortfolio(database *Database) []OptimisedPortfolioRecord {

	// rawportfolio []OwnPortfolioRecord

	var optimisedportfolio []OptimisedPortfolioRecord
	for i := 0; i < len(database.ownstartingportfolio); i++ {
		// same token
		// placeholder values for now
		optimisedportfolio = append(optimisedportfolio, OptimisedPortfolioRecord{database.ownstartingportfolio[i].Token, "Uniswap", database.ownstartingportfolio[i].Amount, 0.420, 0.069, database.Risksetting})
	}

	// For some reason crashes with ZERO elements
	if len(database.ownstartingportfolio) == 0 {
		optimisedportfolio = append(optimisedportfolio, OptimisedPortfolioRecord{"USD", "Uniswap", 100, 1, 0.0125, database.Risksetting})
	}
	fmt.Println("Returning optimised portfolio...")
	return optimisedportfolio
}

// For adding own portfolio records
func NewOwnPortfolioRecord(token string, amount float32) OwnPortfolioRecord { // , pool_sz float32
	return OwnPortfolioRecord{token, amount} // , pool_sz
}

func NewHistoricalCurrencyDataFromRaw(token string, rawhistoricaldata []UniswapDaily) HistoricalCurrencyData {
	var historicaldata HistoricalCurrencyData

	// To DO: Date checking
	for i := 0; i < len(rawhistoricaldata); i++ {
		historicaldata.Date = append(historicaldata.Date, int64(rawhistoricaldata[i].Date))

		price, _ := strconv.ParseFloat(rawhistoricaldata[i].PriceUSD, 32)

		historicaldata.Price = append(historicaldata.Price, float32(price))
	}

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
	ROI_raw_est float32 `default0:"42.69%" json:"ROIestimate"`
	ROI_vol_adj_est float32 `default0:"42.69%" json:"ROIvoladjest"`
	ROI_hist float32 `default0:"42.69%" json:"ROIhist"`
}

func NewCurrencyInputData() CurrencyInputData {
	currencyinputdata := CurrencyInputData{}
	currencyinputdata.Pair = "ETH/DAI"
	currencyinputdata.PoolSize = 420000.69
	currencyinputdata.PoolVolume = 4200.69
	currencyinputdata.Yield = 0.08
	currencyinputdata.Pool = "Uniswap"
	currencyinputdata.Volatility = -0.09
	currencyinputdata.ROI_raw_est = 0.4269
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
	currencyinputdata.ROI_raw_est = roi



	return currencyinputdata
}

type RiskWrapper struct {
	Risksettinginput float32 `json:"risk_setting"`
}

type Database struct {
	// Data structure for Optimisation
	ownstartingportfolio []OwnPortfolioRecord       // for portfolio optimisation table
	optimisedportfolio   []OptimisedPortfolioRecord // for storing output of ownstartingportfolio
	Risksetting          float32

	// Data structure for Ranking
	currencyinputdata      []CurrencyInputData      // LATEST currency pair info
	historicalcurrencydata []HistoricalCurrencyData // historical time series
}

func New() Database {
	ownstartingportfolio := make([]OwnPortfolioRecord, 0)
	optimisedportfolio := make([]OptimisedPortfolioRecord, 0)
	Risksetting := 0.00

	currencyinputdata := make([]CurrencyInputData, 0)
	historicalcurrencydata := make([]HistoricalCurrencyData, 0)

	return Database{ownstartingportfolio, optimisedportfolio, float32(Risksetting), currencyinputdata, historicalcurrencydata}
}

// Add OWN PORTFOLIO data
func (database *Database) AddRecord(r OwnPortfolioRecord) {
	database.ownstartingportfolio = append(database.ownstartingportfolio, r)
}

func (database *Database) AddRiskRecord(risk RiskWrapper) {
	database.Risksetting = risk.Risksettinginput
}

// Retrieve Data
func (database *Database) GetOptimisedPortfolio() []OptimisedPortfolioRecord {
	return OptimisePortfolio(database)
}

func (database *Database) GetCurrencyInputData() []CurrencyInputData {
	return database.currencyinputdata
}
