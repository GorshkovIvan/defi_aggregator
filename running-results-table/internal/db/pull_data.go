package db

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

type UniswapInputStruct struct {
	clientUniswap               *graphql.Client
	reqUniswapIDFromTokenTicker *graphql.Request
	reqUniswapHist              *graphql.Request
}

func (database *Database) AddRecordfromAPI() {

	// 1 - create clients
	clientUniswap := graphql.NewClient("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2")
	//clientAave := graphql.NewClient("https://api.thegraph.com/subgraphs/name/aave/protocol")

	clientCompound := graphql.NewClient("https://api.thegraph.com/subgraphs/name/graphprotocol/compound-v2")
	clientCurve := graphql.NewClient("https://api.thegraph.com/subgraphs/name/protofire/curve")
	clientBancor := graphql.NewClient("https://api.thegraph.com/subgraphs/name/blocklytics/bancor")

	// 2 - declare queries

	reqUniswapHist := graphql.NewRequest(`
				query ($tokenid:String!){
						tokenDayDatas(first: 30 orderBy: date, orderDirection: asc,
						 where: {
						   token:$tokenid
						 }
						) {
						   date
						   priceUSD
						   token{
							   id
							   symbol
						   }
						}
				  }
			`)

	reqUniswapIDFromTokenTicker := graphql.NewRequest(`
						query ($ticker:String!){
							tokens(where:{symbol:$ticker})
							{
								id
								symbol
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

	// 3 - set query variables

	reqAave.Var("address", "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee") // Ether address
	reqCompound.Var("key", "value")
	reqCurve.Var("key", "value")
	reqBancor.Var("key", "value")

	// 4 - set query headers
	reqUniswapIDFromTokenTicker.Header.Set("Cache-Control", "no-cache")
	reqUniswapHist.Header.Set("Cache-Control", "no-cache")

	reqCompound.Header.Set("Cache-Control", "no-cache")
	reqCurve.Header.Set("Cache-Control", "no-cache")
	reqBancor.Header.Set("Cache-Control", "no-cache")
	reqAave.Header.Set("Cache-Control", "no-cache")

	// 5 - define a Context for the request
	ctx := context.Background()

	// 6 - declare query response objects
	var respCompound CompoundQuery // need to change to same format as Balancer and Uniswap
	var respCurve CurveQuery
	var respBancor BancorQuery
	// var respAave AaveQuery

	// 7 - run data queries on each pool
	U := UniswapInputStruct{clientUniswap, reqUniswapIDFromTokenTicker, reqUniswapHist}
	getBalancerData(database, U) // 1
	getUniswapData(database, U)  // 2
	/*
		3) Aave
		4) Curve
		5) Others
	*/
	// RE-RANK THEM HERE? somewhere else?

	/*
		// Checks - start
		fmt.Print("Number of historical database entries: ")
		fmt.Println(len(database.historicalcurrencydata))

		for i := 0; i < len(database.historicalcurrencydata); i++ {
			fmt.Println(database.historicalcurrencydata[i].Ticker)
		}

		// Balancer
		fmt.Print("Number of Balancer Filtered Pools: ")
		fmt.Println(len(BalancerFilteredPoolList))

		fmt.Println("Balancer Filtered Pools: ")
		for i := 0; i < len(BalancerFilteredPoolList); i++ {
			fmt.Println(BalancerFilteredPoolList[i])
		}

		fmt.Println("Balancer Filtered Pool Pairs: ")
		for i := 0; i < len(BalancerFilteredPoolListPairs); i++ {
			fmt.Println(BalancerFilteredPoolListPairs[i])
		}

		fmt.Println("Balancer Filtered Tokens: ")
		for i := 0; i < len(BalancerFilteredTokenList); i++ {
			fmt.Println(BalancerFilteredTokenList[i])
		}

		// Uniswap
		fmt.Print("Number of Uniswap Filtered Pools: ")
		fmt.Println(len(UniswapFilteredPoolList))

		fmt.Println("Uniswap Filtered Pools: ")
		for i := 0; i < len(UniswapFilteredPoolList); i++ {
			fmt.Println(UniswapFilteredPoolList[i])
		}

		fmt.Println("Uniswap Filtered Pool Pairs: ")
		for i := 0; i < len(UniswapFilteredPoolListPairs); i++ {
			fmt.Println(UniswapFilteredPoolListPairs[i])
		}

		fmt.Println("Uniswap Filtered Tokens: ")
		for i := 0; i < len(UniswapFilteredTokenList); i++ {
			fmt.Println(UniswapFilteredTokenList[i])
		}
	*/
	// Checks - end

	// OTHER POOLS - TO DO
	symbol, size, volume, interest, volatility := GetAaveData()
	ROI := calculateROI(interest, 0, volume, volatility)
	database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{symbol, size, volume, interest, "Aave", volatility, ROI})

	if err := clientCompound.Run(ctx, reqCompound, &respCompound); err != nil {
		log.Fatal(err)
	}
	if err := clientCurve.Run(ctx, reqCurve, &respCurve); err != nil {
		log.Fatal(err)
	}
	if err := clientBancor.Run(ctx, reqBancor, &respBancor); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ran all download functions and appended data")
}
