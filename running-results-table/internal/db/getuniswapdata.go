package db

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/machinebox/graphql"
)

func getUniswapData(database *Database, uniswapreqdata UniswapInputStruct) {

	reqUniswapListOfPools := graphql.NewRequest(`
	query{
		pairs(first: 50, orderBy: volumeUSD, orderDirection: desc) {
			id
			untrackedVolumeUSD
			volumeUSD
			token0 {
				id
				symbol
			}
			token1 {
				id
				symbol
			}
		}
	}
`)

	reqUniswapByPoolID := graphql.NewRequest(`
query ($poolid:String!){
	pair(id:$poolid) {
		id
		untrackedVolumeUSD
		volumeUSD
		token0 {
		id
		symbol
		}
		token1 {
		id
		symbol
		}
	}
}
`)

	var respUniswapPoolList UniswapPoolList
	var respUniswapHist UniswapHistQuery
	var respUniswapById UniswapCurrentQuery

	reqUniswapListOfPools.Var("key", "value")
	reqUniswapListOfPools.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	// 7b - UNISWAP
	var UniswapFilteredPoolList []string      // Pairs - IDS - 0x124145
	var UniswapFilteredPoolListPairs []string // Pairs - Tokens ETH/DAI
	var UniswapFilteredTokenList []string     // Tokens - ETH, DAI

	if err := uniswapreqdata.clientUniswap.Run(ctx, reqUniswapListOfPools, &respUniswapPoolList); err != nil {
		log.Fatal(err)
	}

	// Process received list of pools (PAIRS)
	for i := 0; i < len(respUniswapPoolList.Pools); i++ {
		// if len(respUniswapPoolList.Pools[i].Token0) > 1 {
		token0symbol := respUniswapPoolList.Pools[i].Token0.Symbol
		token1symbol := respUniswapPoolList.Pools[i].Token1.Symbol

		if isPoolPartOfFilter(token0symbol, token1symbol) {
			// Filter pools to allowed components (WETH, DAI, USDC, USDT)
			UniswapFilteredPoolList = append(UniswapFilteredPoolList, respUniswapPoolList.Pools[i].ID)
			UniswapFilteredPoolListPairs = append(UniswapFilteredPoolListPairs, token0symbol+"/"+token1symbol)

			var tokenqueue []string
			var tokenqueueIDs []string

			// Split list of pairs into single tokens
			if !stringInSlice(token0symbol, UniswapFilteredTokenList) {
				UniswapFilteredTokenList = append(UniswapFilteredTokenList, token0symbol)
				tokenqueue = append(tokenqueue, token0symbol)
				tokenqueueIDs = append(tokenqueueIDs, respUniswapPoolList.Pools[i].Token0.ID)
			}
			if !stringInSlice(token1symbol, UniswapFilteredTokenList) {
				UniswapFilteredTokenList = append(UniswapFilteredTokenList, token1symbol)
				tokenqueue = append(tokenqueue, token1symbol)
				tokenqueueIDs = append(tokenqueueIDs, respUniswapPoolList.Pools[i].Token1.ID)
			}

			for j := 0; j < len(tokenqueueIDs); j++ {
				// Check if database already has historical data
				if !isHistDataAlreadyDownloaded(tokenqueue[j], database) {
					// No need to get uniswap ids of these tokens
					// Download historical data for each token for which data is missing
					// request data from uniswap using this queried ticker
					uniswapreqdata.reqUniswapHist.Var("tokenid", tokenqueueIDs[j])
					fmt.Print("Querying historical data for: ")
					fmt.Print(tokenqueueIDs[j])
					fmt.Print(" : ")
					fmt.Print(tokenqueue[j])
					if err := uniswapreqdata.clientUniswap.Run(ctx, uniswapreqdata.reqUniswapHist, &respUniswapHist); err != nil {
						log.Fatal(err)
					}
					fmt.Print("| returned days: ")
					fmt.Println(len(respUniswapHist.DailyTimeSeries))
					// if returned data - append it to database
					if len(respUniswapHist.DailyTimeSeries) > 0 {
						database.historicalcurrencydata = append(database.historicalcurrencydata, NewHistoricalCurrencyDataFromRaw(tokenqueue[j], respUniswapHist.DailyTimeSeries))
					}
				} // if historical data needs updating
			} // tokenqueue loop ends

			// if historical data is in order - get current data
			reqUniswapByPoolID.Var("poolid", respUniswapPoolList.Pools[i].ID)

			if err := uniswapreqdata.clientUniswap.Run(ctx, reqUniswapByPoolID, &respUniswapById); err != nil {
				log.Fatal(err)
			}

			currentSize := float32(1000.000)
			currentVolume, _ := strconv.ParseFloat(respUniswapById.Pair.VolumeUSD, 32) // No historical for now
			currentInterestrate := float32(0.00)                                       // Zero for liquidity pool
			UniswapRewardPercentage := float32(0.003)                                  // Placeholder

			volatility := calculatehistoricalvolatility(retrieveDataForTokensFromDatabase(token0symbol, token1symbol, database), 30)
			ROI := calculateROI(currentInterestrate, UniswapRewardPercentage, float32(currentVolume), volatility)

			var recordalreadyexists bool
			recordalreadyexists = false

			for k := 0; k < len(database.currencyinputdata); k++ {
				// Means record already exists - UPDATE IT, DO NOT APPEND
				if database.currencyinputdata[k].Pair == token0symbol+"/"+token1symbol && database.currencyinputdata[k].Pool == "Balancer" {
					recordalreadyexists = true
					database.currencyinputdata[k].PoolSize = currentSize
					database.currencyinputdata[k].PoolVolume = float32(currentVolume)
					database.currencyinputdata[k].ROIestimate = ROI
					database.currencyinputdata[k].Volatility = volatility
					database.currencyinputdata[k].Yield = currentInterestrate
				}
			}

			// APPEND IF NEW
			if !recordalreadyexists {
				database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{token0symbol + "/" + token1symbol, float32(currentSize), float32(currentVolume), currentInterestrate, "Uniswap", volatility, ROI})
			}
		} // if pool is within pre filtered list ends
		// } // if pool has some tokens ends
	} // Uniswap pair loop closes

}
