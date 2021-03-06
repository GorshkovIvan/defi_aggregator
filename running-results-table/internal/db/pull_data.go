package db

import (
	"fmt"

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
	//	clientCompound := graphql.NewClient("https://api.thegraph.com/subgraphs/name/graphprotocol/compound-v2")
	//	clientCurve := graphql.NewClient("https://api.thegraph.com/subgraphs/name/protofire/curve")
	//	clientBancor := graphql.NewClient("https://api.thegraph.com/subgraphs/name/blocklytics/bancor")

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

	// 4 - set query headers
	reqUniswapIDFromTokenTicker.Header.Set("Cache-Control", "no-cache")
	reqUniswapHist.Header.Set("Cache-Control", "no-cache")

	// 5 - define a Context for the request
	// ctx := context.Background()

	// 7 - run data queries on each pool
	U := UniswapInputStruct{clientUniswap, reqUniswapIDFromTokenTicker, reqUniswapHist}
	getBalancerData(database, U) // 1
	getUniswapData(database, U)  // 2
	getAaveData(database, U)     //	3
	/*
		4) Curve
		5) Others
	*/

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

	fmt.Println("Ran all download functions and appended data")
}
