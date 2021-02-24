package db

import (
	"sort"
)

// ---Bancor---
type BancorQuery struct {
	Swaps []BancorSwap `json:"swaps"`
}

type BancorSwap struct {
	ID string `json:"id"`
	// Need separate structs for these datatypes
	//FromToken BancorToken `json:"fromToken"`
	//ToToken BancorToken `json:"toToken"`
	AmountPurchased                     string `json:"amountPurchased"`
	AmountReturned                      string `json:"amountReturned"`
	Price                               string `json:"price"`
	InversePrice                        string `json:"inversePrice"`
	ConverterWeight                     string `json:"converterWeight"`
	ConverterFromTokenBalanceBeforeSwap string `json:"converterFromTokenBalanceBeforeSwap"`
	ConverterFromTokenBalanceAfterSwap  string `json:"converterFromTokenBalanceAfterSwap"`
	ConverterToTokenBalanceBeforeSwap   string `json:"converterToTokenBalanceBeforeSwap"`
	ConverterToTokenBalanceAfterSwap    string `json:"converterToTokenBalanceAfterSwap"`
	Slippage                            string `json:"slippage"`
	ConversionFee                       string `json:"conversionFee"`
	// Need separate structs for these datatypes
	//ConverterUsed Converter `json:"converterUsed"`
	//Transaction Transaction `json:"transaction"`
	//Trader User `json:"trader"`
	Timestamp string `json:"timestamp"`
	LogIndex  int    `json:"logIndex"`
}

// --AAVE--
/*
type AaveQuery struct {
	Reserve AaveData `json:"reserve"`
}

type AaveData struct {
	ID                 string `json:"id"`
	Symbol             string `json:"symbol"`
	LiquidityRate      string `json:"liquidityRate"`
	StableBorrowRate   string `json:"stableBorrowRate"`
	VariableBorrowRate string `json:"variableBorrowRate"`
	TotalBorrows       string `json:"totalBorrows"`
}
*/

type Record struct {
	Pair    string  `json:"pair"`
	Amount  float32 `json:"amount"`
	Pool_sz float32 `json:"pool_sz"`
}

func NewRecord(pair string, amount float32, pool_sz float32) Record {
	return Record{pair, amount, pool_sz}
}

type CurrencyInputData struct {
	Pair        string  `default0:"ETH/DAI" json:"backend_pair"`
	Amount      float32 `default0:"123" json:"backend_amount"`
	Yield       float32 `default0:"0.05" json:"backend_yield"`
	Pool        string  `default0:"Uniswap" json:"pool_source"`
	Volatility  float32 `default0:"-9.00" json:"volatility"`
	ROIestimate float32 `default0:"42.69%" json:"ROIestimate"`
}

func NewHistoricalCurrencyDataFromRaw(token string, rawhistoricaldata []UniswapDaily) HistoricalCurrencyData {
	// TO DO: parsing
	historicaldata := HistoricalCurrencyData{}
	historicaldata.Ticker = token
	return historicaldata
}

func NewHistoricalCurrencyData() HistoricalCurrencyData {
	// TO DO: fill defaults
	historicaldata := HistoricalCurrencyData{}
	return historicaldata
}

type HistoricalCurrencyData struct {
	Date   []float32 `default0:"ETH/DAI" json:"date"`
	Price  []float32 `default0:"123" json:"price"`
	Ticker string    `default0:"0.05" json:"ticker"`
}

func NewCurrencyInputData() CurrencyInputData {
	currencyinputdata := CurrencyInputData{}
	currencyinputdata.Pair = "ETH/DAI"
	currencyinputdata.Amount = 420.69
	currencyinputdata.Yield = 0.08
	currencyinputdata.Pool = "Uniswap"
	currencyinputdata.Volatility = -0.09
	currencyinputdata.ROIestimate = 0.4269
	return currencyinputdata
}

func NewCurrencyInputDataAct(pair string, amount float32, yield float32, pool string, volatility float32, roi float32) CurrencyInputData {
	currencyinputdata := CurrencyInputData{}
	currencyinputdata.Pair = pair
	currencyinputdata.Amount = amount
	currencyinputdata.Yield = yield
	currencyinputdata.Pool = pool
	currencyinputdata.Volatility = volatility
	currencyinputdata.ROIestimate = roi
	return currencyinputdata
}

type Database struct {
	contents               []Record
	currencyinputdata      []CurrencyInputData // this will store the LATEST currency pair info
	historicalcurrencydata []HistoricalCurrencyData
}

func New() Database {
	contents := make([]Record, 0)
	currencyinputdata := make([]CurrencyInputData, 0)

	historicalcurrencydata := make([]HistoricalCurrencyData, 0)

	return Database{contents, currencyinputdata, historicalcurrencydata}
}
func (database *Database) AddRecord(r Record) {
	database.contents = append(database.contents, r)
}
func (database *Database) GetRecords() []Record {
	return database.contents
}

func (database *Database) AddRecordfromAPI2(pair string, amount float32, yield float32, pool string, volatility float32, ROIestimate float32) {
	//c CurrencyInputData = NewCurrencyInputData()
	database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{pair, amount, yield, pool, volatility, ROIestimate})
}

func (database *Database) GetCurrencyInputData() []CurrencyInputData {
	return database.currencyinputdata
}

func (database *Database) RankBestCurrencies() {

	sort.Slice(database.currencyinputdata, func(i, j int) bool {
		return database.currencyinputdata[i].Yield > database.currencyinputdata[j].Yield
	})
}
