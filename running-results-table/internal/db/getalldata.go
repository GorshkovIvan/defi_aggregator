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
						tokenDayDatas(first: 30 orderBy: date, orderDirection: desc,
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

	// 3 - set query headers
	reqUniswapIDFromTokenTicker.Header.Set("Cache-Control", "no-cache")
	reqUniswapHist.Header.Set("Cache-Control", "no-cache")

	// 4 - run data queries on each pool
	U := UniswapInputStruct{clientUniswap, reqUniswapIDFromTokenTicker, reqUniswapHist}

	getUniswapData(database, U)
	getBalancerData(database, U)
	//getCurveData()
	//getAave1Data()
	//getAave2Data()
	// getAaveData(database, U)     //

	fmt.Println("Ran all download functions and appended data")
}
