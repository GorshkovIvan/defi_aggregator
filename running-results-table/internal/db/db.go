package db

import (
	"context"
	"log"
	"sort"
	"strconv"

	"github.com/machinebox/graphql"
)

// ---Compound---
type CompoundQuery struct {
	Markets []CompoundMarket `json:"markets"`
}

type CompoundMarket struct {
	BorrowRate string `json:"borrowRate"`
	Cash string `json:"cash"`
	CollateralFactor string `json:"collateralFactor"`
	ExchangeRate string `json:"exchangeRate"`
	InterestRateModelAddress string `json:"interestRateModelAddress"`
	Name string `json:"name"`
	Reserves string `json:"reserves"`
	SupplyRate string `json:"supplyRate"`
	Symbol string `json:"symbol"`
	ID string `json:"id"`
	TotalBorrows string `json:"totalBorrows"`
	TotalSupply string `json:"totalSupply"`
	UnderlyingAddress string `json:"underlyingAddress"`
	UnderlyingName string `json:"underlyingName"`
	UnderlyingPrice string `json:"underlyingPrice"`
	UnderlyingSymbol string `json:"underlyingSymbol"`
	ReserveFactor string `json:"reserveFactor"`
	UnderlyingPriceUSD string `json:"underlyingPriceUSD"`
}


// ---Balancer---
type BalancerQuery struct {
	Pools []BalancerPool `json:"pools"`
}

type BalancerPool struct {
	ID string `json:"id"`
	Finalized bool `json:"finalized"`
	PublicSwap bool `json:"publicSwap"`
	SwapFee string `json:"swapFee"`
	TotalWeight string `json:"totalWeight"`
	TokensList []string `json:"tokensList"`
	Tokens []BalancerToken `json:"tokens"`
}

type BalancerToken struct {
	ID string `json:"id"`
	Address string `json:"address"`
	Balance string `json:"balance"`
	Decimals int `json:"decimals"`
	Symbol string `json:"symbol"`
	DenormWeight string `json:"denormWeight"`
}


// ---Curve---
type CurveQuery struct {
	Pools []CurvePool `json:"pools"`
}

type CurvePool struct {
	Address string `json:"address"`
	CoinCount int `json:"coinCount"`
	A string `json:"A"`
	Fee string `json:"fee"`
	AdminFee string `json:"adminFee"`
	Balances []string `json:"balances"`
	Coins []CurveCoin `json:"coins"`
}

type CurveCoin struct {
	Address string `json:"address"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Decimals string `json:"decimals"`
}


// ---Uniswap---
type UniswapQuery struct {
	UniswapFactories []UniswapFactory `json:"uniswapFactories"`
}

type UniswapFactory struct {
	ID string `json:"id"`
	PairCount int `json:"pairCount"`
	TotalVolumeUSD string `json:"totalVolumeUSD"`
	TotalVolumeETH string `json:"totalVolumeETH"`
	UntrackedVolumeUSD string `json:"untrackedVolumeUSD"`
	TotalLiquidityUSD string `json:"totalLiquidityUSD"`
	TotalLiquidityETH string `json:"totalLiquidityETH"`
	TXCount string `json:"txCount"`
	MostLiquidTokens []UniswapTokenDayData `json:"mostLiquidTokens"`
}

type UniswapTokenDayData struct {
	ID string `json:"id"`
	// Other fields can be added as necessary
}


// ---Bancor---
type BancorQuery struct {
	Swaps []BancorSwap `json:"swaps"`
}

type BancorSwap struct {
	ID string `json:"id"`
	// Need separate structs for these datatypes
	//FromToken BancorToken `json:"fromToken"`
	//ToToken BancorToken `json:"toToken"`
	AmountPurchased string `json:"amountPurchased"`
	AmountReturned string `json:"amountReturned"`
	Price string `json:"price"`
	InversePrice string `json:"inversePrice"`
	ConverterWeight string `json:"converterWeight"`
	ConverterFromTokenBalanceBeforeSwap string `json:"converterFromTokenBalanceBeforeSwap"`
	ConverterFromTokenBalanceAfterSwap string `json:"converterFromTokenBalanceAfterSwap"`
	ConverterToTokenBalanceBeforeSwap string `json:"converterToTokenBalanceBeforeSwap"`
	ConverterToTokenBalanceAfterSwap string `json:"converterToTokenBalanceAfterSwap"`
	Slippage string `json:"slippage"`
	ConversionFee string `json:"conversionFee"`
	// Need separate structs for these datatypes
	//ConverterUsed Converter `json:"converterUsed"`
	//Transaction Transaction `json:"transaction"`
	//Trader User `json:"trader"`
	Timestamp string `json:"timestamp"`
	LogIndex int `json:"logIndex"`
}

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
	
	// create clients
	clientBalancer := graphql.NewClient("https://api.thegraph.com/subgraphs/name/balancer-labs/balancer")
	clientCompound := graphql.NewClient("https://api.thegraph.com/subgraphs/name/graphprotocol/compound-v2")
	clientUniswap := graphql.NewClient("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2")
	clientCurve := graphql.NewClient("https://api.thegraph.com/subgraphs/name/protofire/curve")
	clientBancor := graphql.NewClient("https://api.thegraph.com/subgraphs/name/blocklytics/bancor")

	// make requests
	reqBalancer := graphql.NewRequest(`
		query {
			pools(first: 10, where: {publicSwap: true}) {
				id
				finalized
				publicSwap
				swapFee
				totalWeight
				tokensList
				tokens {
					id
					address
					balance
					decimals
					symbol
					denormWeight
				}
			}
		}
	`)

	reqCompound := graphql.NewRequest(`
		query {
			markets(first: 10) {
				borrowRate
				cash
				collateralFactor
				exchangeRate
				interestRateModelAddress
				name
				reserves
				supplyRate
				symbol
				id
				totalBorrows
				totalSupply
				underlyingAddress
				underlyingName
				underlyingPrice
				underlyingSymbol
				reserveFactor
				underlyingPriceUSD
			}
		}  
	`)

	// More parameters to be added as necessary
	reqUniswap := graphql.NewRequest(`
		query {
			uniswapFactories(first: 10) {
				id
				pairCount
				totalVolumeUSD
				totalVolumeETH
				untrackedVolumeUSD
				totalLiquidityUSD
				totalLiquidityETH
				txCount
				mostLiquidTokens {
					id
				}
			}
		}
	`)

	reqCurve := graphql.NewRequest(`
		query {
			pools(orderBy: addedAt, first: 10) {
				address
				coinCount
				A
				fee
				adminFee
				balances
				coins {
					address
					name
					symbol
					decimals
			  	}
			}
		}
	`)

	reqBancor := graphql.NewRequest(`
		query {
			swaps(first: 10, skip: 0, orderBy: timestamp, orderDirection: desc) {
				id
				amountPurchased
				amountReturned
				price
				inversePrice
				converterWeight
				converterFromTokenBalanceBeforeSwap
				converterFromTokenBalanceAfterSwap
				converterToTokenBalanceBeforeSwap
				converterToTokenBalanceAfterSwap
				slippage
				conversionFee
				timestamp
				logIndex
			}
		}
	`)

	// set any variables
	reqBalancer.Var("key", "value")
	reqCompound.Var("key", "value")
	reqUniswap.Var("key", "value")
	reqCurve.Var("key", "value")
	reqBancor.Var("key", "value")

	// set header fields
	reqBalancer.Header.Set("Cache-Control", "no-cache")
	reqCompound.Header.Set("Cache-Control", "no-cache")
	reqUniswap.Header.Set("Cache-Control", "no-cache")
	reqCurve.Header.Set("Cache-Control", "no-cache")
	reqBancor.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respBalancer BalancerQuery
	var respCompound CompoundQuery
	var respUniswap UniswapQuery
	var respCurve CurveQuery
	var respBancor BancorQuery

	if err := clientBalancer.Run(ctx, reqBalancer, &respBalancer); err != nil {
		log.Fatal(err)
	}
	if err := clientCompound.Run(ctx, reqCompound, &respCompound); err != nil {
		log.Fatal(err)
	}
	if err := clientUniswap.Run(ctx, reqUniswap, &respUniswap); err != nil {
		log.Fatal(err)
	}
	if err := clientCurve.Run(ctx, reqCurve, &respCurve); err != nil {
		log.Fatal(err)
	}
	if err := clientBancor.Run(ctx, reqBancor, &respBancor); err != nil {
		log.Fatal(err)
	}

	// Example data append function
	// To be completed when we know what inputs we want for the ROI calculation

	for i := 0; i < len(respBalancer.Pools); i++ {
		value, _ := strconv.ParseFloat(respBalancer.Pools[i].Tokens[0].Balance, 32)
		value2, _ := strconv.ParseFloat(respBalancer.Pools[i].Tokens[1].Balance, 32)
		float := float32(value)
		float2 := float32(value2)
		database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{respBalancer.Pools[i].Tokens[0].Symbol + "/" + respBalancer.Pools[i].Tokens[1].Symbol, float, float2})
	}
	//c CurrencyInputData = NewCurrencyInputData()
	//database.currencyinputdata = append(database.currencyinputdata, NewCurrencyInputData())
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
