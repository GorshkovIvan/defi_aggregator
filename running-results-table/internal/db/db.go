package db

import (
	"context"
	"fmt"
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
	BorrowRate               string `json:"borrowRate"`
	Cash                     string `json:"cash"`
	CollateralFactor         string `json:"collateralFactor"`
	ExchangeRate             string `json:"exchangeRate"`
	InterestRateModelAddress string `json:"interestRateModelAddress"`
	Name                     string `json:"name"`
	Reserves                 string `json:"reserves"`
	SupplyRate               string `json:"supplyRate"`
	Symbol                   string `json:"symbol"`
	ID                       string `json:"id"`
	TotalBorrows             string `json:"totalBorrows"`
	TotalSupply              string `json:"totalSupply"`
	UnderlyingAddress        string `json:"underlyingAddress"`
	UnderlyingName           string `json:"underlyingName"`
	UnderlyingPrice          string `json:"underlyingPrice"`
	UnderlyingSymbol         string `json:"underlyingSymbol"`
	ReserveFactor            string `json:"reserveFactor"`
	UnderlyingPriceUSD       string `json:"underlyingPriceUSD"`
}

// ---Balancer---
type BalancerPoolID struct {
	ID         string          `json:"id"`
	TokensList []string        `json:"tokensList"`
	Tokens     []BalancerToken `json:"tokens"`
}

type BalancerPoolList struct {
	Pools []BalancerPoolID `json:"pools"`
}

type BalancerQuery struct {
	Pools []BalancerPool `json:"pools"`
}

type BalancerPool struct {
	ID              string          `json:"id"`
	Finalized       bool            `json:"finalized"`
	PublicSwap      bool            `json:"publicSwap"`
	SwapFee         string          `json:"swapFee"`
	TotalSwapVolume string          `json:"totalSwapVolume"`
	TotalWeight     string          `json:"totalWeight"`
	TokensList      []string        `json:"tokensList"`
	Tokens          []BalancerToken `json:"tokens"`
}

type BalancerById struct {
	BalancerPool `json:"pool"`
}

type BalancerToken struct {
	ID           string `json:"id"`
	Address      string `json:"address"`
	Balance      string `json:"balance"`
	Decimals     int    `json:"decimals"`
	Symbol       string `json:"symbol"`
	DenormWeight string `json:"denormWeight"`
}

// ---Curve---
type CurveQuery struct {
	Pools []CurvePool `json:"pools"`
}

type CurvePool struct {
	Address   string      `json:"address"`
	CoinCount int         `json:"coinCount"`
	A         string      `json:"A"`
	Fee       string      `json:"fee"`
	AdminFee  string      `json:"adminFee"`
	Balances  []string    `json:"balances"`
	Coins     []CurveCoin `json:"coins"`
}

type CurveCoin struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals string `json:"decimals"`
}

// ---Uniswap---
type UniswapQuery struct {
	UniswapFactories []UniswapFactory `json:"uniswapFactories"`
}

type UniswapFactory struct {
	ID                 string                `json:"id"`
	PairCount          int                   `json:"pairCount"`
	TotalVolumeUSD     string                `json:"totalVolumeUSD"`
	TotalVolumeETH     string                `json:"totalVolumeETH"`
	UntrackedVolumeUSD string                `json:"untrackedVolumeUSD"`
	TotalLiquidityUSD  string                `json:"totalLiquidityUSD"`
	TotalLiquidityETH  string                `json:"totalLiquidityETH"`
	TXCount            string                `json:"txCount"`
	MostLiquidTokens   []UniswapTokenDayData `json:"mostLiquidTokens"`
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
	Pool   string  `default0:"Uniswap" json:"pool_source"`
}

func NewCurrencyInputData() CurrencyInputData {
	currencyinputdata := CurrencyInputData{}
	currencyinputdata.Pair = "ETH/DAI"
	currencyinputdata.Amount = 420.69
	currencyinputdata.Yield = 0.08
	currencyinputdata.Pool = "Uniswap"
	return currencyinputdata
}

func NewCurrencyInputDataAct(pair string, amount float32, yield float32, pool string) CurrencyInputData {
	currencyinputdata := CurrencyInputData{}
	currencyinputdata.Pair = pair
	currencyinputdata.Amount = amount
	currencyinputdata.Yield = yield
	currencyinputdata.Pool = pool
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
	clientAave := graphql.NewClient("https://api.thegraph.com/subgraphs/name/aave/protocol")

	reqBalancerListOfPools := graphql.NewRequest(`
	query {      
        pools{
        id
        tokens {
          id
          symbol
        }
      }
			}
`)

	// make requests
	reqBalancer := graphql.NewRequest(`
		query {
			pools(first: 10, where: {publicSwap: true}) {
				id
				finalized
				publicSwap
				swapFee
				totalSwapVolume
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

	reqBalancerbyIDvar := graphql.NewRequest(`
	query ($poolid:String!){      
		pool(id:$poolid) {
			id
			finalized
			publicSwap
			swapFee
			totalSwapVolume
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

	reqAave := graphql.NewRequest(`
		query($address: ID!)
		{
				reserve(id: $address){
				id
				symbol
				liquidityRate
				stableBorrowRate
				variableBorrowRate
				totalBorrows
			}
			}
				`)

	// set any variables
	reqBalancer.Var("key", "value")
	reqCompound.Var("key", "value")
	reqUniswap.Var("key", "value")
	reqCurve.Var("key", "value")
	reqBancor.Var("key", "value")
	reqBalancerListOfPools.Var("key", "value")
	reqAave.Var("address", "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee") // Ether address

	// set header fields
	reqBalancer.Header.Set("Cache-Control", "no-cache")
	reqBalancerbyIDvar.Header.Set("Cache-Control", "no-cache")
	reqCompound.Header.Set("Cache-Control", "no-cache")
	reqUniswap.Header.Set("Cache-Control", "no-cache")
	reqCurve.Header.Set("Cache-Control", "no-cache")
	reqBancor.Header.Set("Cache-Control", "no-cache")
	reqBalancerListOfPools.Header.Set("Cache-Control", "no-cache")
	reqAave.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	//	var respBalancer BalancerQuery
	var respCompound CompoundQuery
	var respUniswap UniswapQuery
	var respCurve CurveQuery
	var respBancor BancorQuery
	var respBalancerById BalancerById
	var respBalancerPoolList BalancerPoolList
	var respAave AaveQuery

	if err := clientAave.Run(ctx, reqAave, &respAave); err != nil {
		log.Fatal(err)
	}

	value, _ := strconv.ParseFloat(respAave.Reserve.TotalBorrows, 32)
	float := float32(value)
	float2 := float32(111)
	database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{respAave.Reserve.Symbol, float, float2, "Aave"})

	if err := clientBalancer.Run(ctx, reqBalancerListOfPools, &respBalancerPoolList); err != nil {
		log.Fatal(err)
	}

	var BalancerWETHPoolList []string
	fmt.Println(len(BalancerWETHPoolList))
	fmt.Println("how many pools in total:")
	fmt.Println(len(respBalancerPoolList.Pools))

	for i := 0; i < len(respBalancerPoolList.Pools); i++ {
		if len(respBalancerPoolList.Pools[i].Tokens) > 1 {
			if respBalancerPoolList.Pools[i].Tokens[0].Symbol == "WETH" || respBalancerPoolList.Pools[i].Tokens[1].Symbol == "WETH" {
				//fmt.Print(respBalancerPoolList.Pools[i].Tokens[0].Symbol)
				//fmt.Println(respBalancerPoolList.Pools[i].Tokens[1].Symbol)
				BalancerWETHPoolList = append(BalancerWETHPoolList, respBalancerPoolList.Pools[i].ID)
			}
		}

	}

	//fmt.Println(len(BalancerWETHPoolList))
	var BalancerETHPools BalancerQuery

	x := len(BalancerWETHPoolList) - 0

	for i := 0; i < x; i++ {
		reqBalancerbyIDvar.Var("poolid", BalancerWETHPoolList[i])
		fmt.Print(i)
		fmt.Print(": ")
		fmt.Print(BalancerWETHPoolList[i])
		fmt.Print(": ")
		if err := clientBalancer.Run(ctx, reqBalancerbyIDvar, &respBalancerById); err != nil {
			log.Fatal(err)
		}
		// var xx BalancerPool

		BalancerETHPools.Pools = append(BalancerETHPools.Pools, respBalancerById.BalancerPool) //
		fmt.Print(BalancerETHPools.Pools[i].Tokens[0].Symbol)
		fmt.Println(BalancerETHPools.Pools[i].Tokens[1].Symbol)

		value, _ := strconv.ParseFloat(BalancerETHPools.Pools[i].TotalSwapVolume, 32)
		value2, _ := strconv.ParseFloat(BalancerETHPools.Pools[i].TotalSwapVolume, 32) // float32(), 0 // strconv.ParseFloat(float32(len(respBalancerPoolList.Pools)), 32)
		float := float32(value)                                                        // float32(respBalancerById.Data.Pool.TotalSwapVolume)
		float2 := float32(value2)                                                      // float32(respBalancerById.Data.Pool.TotalSwapVolume)

		//		fmt.Print("Appending to database: ")
		//		fmt.Print(i)
		//		fmt.Print(": ")
		fmt.Println(BalancerETHPools.Pools[i].Tokens[0].Symbol + "/" + BalancerETHPools.Pools[i].Tokens[1].Symbol)
		database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{BalancerETHPools.Pools[i].Tokens[0].Symbol + "/" + BalancerETHPools.Pools[i].Tokens[1].Symbol, float, float2, "Balancer"})

	}

	//	fmt.Println("Downloaded data for pools with ETH: ")
	//	fmt.Println(len(BalancerETHPools.Pools))

	//if err := clientBalancer.Run(ctx, reqBalancer, &respBalancer); err != nil {
	//	log.Fatal(err)
	//}

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
	/*
		for i := 0; i < len(respBalancer.Pools); i++ {
			value, _ := strconv.ParseFloat(respBalancer.Pools[i].Tokens[0].Balance, 32)
			value2, _ := strconv.ParseFloat(respBalancer.Pools[i].Tokens[1].Balance, 32)
			float := float32(value)
			float2 := float32(value2)
			database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{respBalancer.Pools[i].Tokens[0].Symbol + "/" + respBalancer.Pools[i].Tokens[1].Symbol, float, float2})
		}
	*/

	//c CurrencyInputData = NewCurrencyInputData()
	//database.currencyinputdata = append(database.currencyinputdata, NewCurrencyInputData())
}

func (database *Database) AddRecordfromAPI2(pair string, amount float32, yield float32, pool string) {
	//c CurrencyInputData = NewCurrencyInputData()
	database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{pair, amount, yield, pool})
}

func (database *Database) GetCurrencyInputData() []CurrencyInputData {
	return database.currencyinputdata
}

func (database *Database) RankBestCurrencies() {

	sort.Slice(database.currencyinputdata, func(i, j int) bool {
		return database.currencyinputdata[i].Yield > database.currencyinputdata[j].Yield
	})
}
