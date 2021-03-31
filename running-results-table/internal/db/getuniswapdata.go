package db

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/machinebox/graphql"
)

func estimate_future_uniswap_volume_and_pool_sz(histvolume UniswapHistVolumeQuery) (float32, float32) {
	future_volume_est := 0.0
	future_sz_est := 0.0

	var count float64
	var count_sz float64
	count = 0
	count_sz = 0

	for i := 0; i < len(histvolume.DailyTimeSeries); i++ {

		//v0, _ := strconv.ParseFloat(histvolume.DailyTimeSeries[i].DailyVolumeToken0, 64)
		//v1, _ := strconv.ParseFloat(histvolume.DailyTimeSeries[i].DailyVolumeToken1, 64)
		v, _ := strconv.ParseFloat(histvolume.DailyTimeSeries[i].DailyVolumeUSD, 64)
		sz, _ := strconv.ParseFloat(histvolume.DailyTimeSeries[i].ReserveUSD, 64)

		//future_volume_est += v0 //future_volume_est += v1
		future_volume_est += v
		future_sz_est += sz

		if v != 0.0 {
			count++
		}

		if sz != 0.0 {
			count_sz++
		}

	}

	future_volume_est = future_volume_est / count
	future_sz_est = future_sz_est / count_sz

	if math.IsNaN(float64(future_volume_est)) {
		future_volume_est = -995.0
	}
	if math.IsNaN(float64(future_sz_est)) {
		future_sz_est = -996.0
	}

	return float32(future_volume_est), float32(future_sz_est) // USD
}

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

	// New

	reqUniswapAllPairs := graphql.NewRequest(`
query{				
	pairs(first: 1000, orderBy: reserveUSD, orderDirection: desc) {
	  id
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

	reqUniswapHistVolume := graphql.NewRequest(`
		query($pairid:String!){				
			pairDayDatas(first:30, orderBy: date, orderDirection: asc, 
				where: {pairAddress:$pairid}				
				) {
			id 
			date
			token0 {
				id
				symbol
			}
			token1 {
				id
				symbol
			}
			dailyVolumeToken0
			dailyVolumeToken1
			dailyVolumeUSD
			totalSupply
			reserveUSD
			}
		}
`)

	var respUniswapPoolList UniswapPoolList
	var respUniswapHist UniswapHistQuery
	var respUniswapById UniswapCurrentQuery
	var respUniswapHistVolume UniswapHistVolumeQuery
	var respUniswapPairList UniswapPairList

	reqUniswapListOfPools.Var("key", "value")
	reqUniswapListOfPools.Header.Set("Cache-Control", "no-cache")

	reqUniswapHistVolume.Var("key", "value")
	reqUniswapHistVolume.Header.Set("Cache-Control", "no-cache")

	//	respUniswapPairList.Var("key", "value")
	//	respUniswapPairList.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	// 7b - UNISWAP
	var UniswapFilteredPoolList []string      // Pairs - IDS - 0x124145
	var UniswapFilteredPoolListPairs []string // Pairs - Tokens ETH/DAI
	var UniswapFilteredTokenList []string     // Tokens - ETH, DAI

	// request all pairs
	if err := uniswapreqdata.clientUniswap.Run(ctx, reqUniswapAllPairs, &respUniswapPairList); err != nil {
		log.Fatal(err)
	}

	// arrays to hold interesting pairs
	var UniswapPairOfInterestIDList []string
	var UniswapPairOfInterestIDToken0 []string
	var UniswapPairOfInterestIDToken1 []string

	for i := 0; i < len(respUniswapPairList.Pairs); i++ {
		// check if both tokens are of interest
		if isPoolPartOfFilter(respUniswapPairList.Pairs[i].Token0.Symbol, respUniswapPairList.Pairs[i].Token1.Symbol) {
			// save pair id to database
			UniswapPairOfInterestIDList = append(UniswapPairOfInterestIDList, respUniswapPairList.Pairs[i].ID)
			UniswapPairOfInterestIDToken0 = append(UniswapPairOfInterestIDToken0, respUniswapPairList.Pairs[i].Token0.Symbol)
			UniswapPairOfInterestIDToken1 = append(UniswapPairOfInterestIDToken1, respUniswapPairList.Pairs[i].Token1.Symbol)
		}
	}

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

			//currentSize := float32(1000.000)
			// currentVolume, _ := strconv.ParseFloat(respUniswapById.Pair.VolumeUSD, 32) //
			currentInterestrate := float32(0.00)      // Zero for liquidity pool
			UniswapRewardPercentage := float32(0.003) // Placeholder

			// Do a historical volume request
			fmt.Println("REQUESTING HISTORICAL VOLUME DATA: ")
			fmt.Println(respUniswapPoolList.Pools[i].ID)
			fmt.Println(respUniswapPoolList.Pools[i].Token0.Symbol)
			fmt.Println(respUniswapPoolList.Pools[i].Token1.Symbol)

			var pairid string

			// find the pair id from 2 tokens
			fmt.Println("TRYING TO MATCH TO PAIR ID: ")
			for jjj := 0; jjj < len(UniswapPairOfInterestIDList); jjj++ {
				matches := 0
				// respUniswapPoolList.Pools[i].Token1.Symbol
				fmt.Println(jjj)
				fmt.Println("Pair of interest: ")
				fmt.Println(UniswapPairOfInterestIDToken0[jjj] + "/" + UniswapPairOfInterestIDToken1[jjj])
				fmt.Println("Vs: ")
				fmt.Println(respUniswapPoolList.Pools[i].Token0.Symbol + "/" + respUniswapPoolList.Pools[i].Token1.Symbol)

				if UniswapPairOfInterestIDToken0[jjj] == respUniswapPoolList.Pools[i].Token0.Symbol {
					matches += 1
				}
				if UniswapPairOfInterestIDToken0[jjj] == respUniswapPoolList.Pools[i].Token1.Symbol {
					matches += 1
				}
				if UniswapPairOfInterestIDToken1[jjj] == respUniswapPoolList.Pools[i].Token0.Symbol {
					matches += 1
				}
				if UniswapPairOfInterestIDToken1[jjj] == respUniswapPoolList.Pools[i].Token1.Symbol {
					matches += 1
				}
				// if match on both
				if matches == 2 {
					// then found
					fmt.Println("Matched!")
					pairid = UniswapPairOfInterestIDList[jjj]
					break
				}

			}

			fmt.Println("FOUND PAIR ID FOR: ")
			fmt.Println(respUniswapPoolList.Pools[i].Token0.Symbol)
			fmt.Println(respUniswapPoolList.Pools[i].Token1.Symbol)
			fmt.Println(pairid)

			// find the dash
			// remove it
			// func Index(s, substr string) int
			pairid = strings.TrimRight(pairid, "-")
			fmt.Println(strings.TrimRight(pairid, "-"))

			reqUniswapHistVolume.Var("pairid", pairid) // respUniswapPoolList.Pools[i].ID
			if err := uniswapreqdata.clientUniswap.Run(ctx, reqUniswapHistVolume, &respUniswapHistVolume); err != nil {
				log.Fatal(err)
			}

			fmt.Println("retrieved number of HISTORICAL VOLUME items: ")
			fmt.Println(len(respUniswapHistVolume.DailyTimeSeries))
			/*
				fmt.Println("RETRIEVED PAIR ID - XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX: ")
				fmt.Println(respUniswapHistVolume.DailyTimeSeries[0].ID)
				fmt.Println(respUniswapHistVolume.DailyTimeSeries[0].Date)
				fmt.Println(respUniswapHistVolume.DailyTimeSeries[0].DailyVolumeToken0)
				fmt.Println(respUniswapHistVolume.DailyTimeSeries[0].DailyVolumeToken1)

				fmt.Println(respUniswapHistVolume.DailyTimeSeries[1].DailyVolumeToken0)
				fmt.Println(respUniswapHistVolume.DailyTimeSeries[1].DailyVolumeToken1)

				fmt.Println(respUniswapHistVolume.DailyTimeSeries[2].DailyVolumeToken0)
				fmt.Println(respUniswapHistVolume.DailyTimeSeries[2].DailyVolumeToken1)
			*/
			future_daily_volume_est, future_pool_sz_est := estimate_future_uniswap_volume_and_pool_sz(respUniswapHistVolume)
			volatility := calculatehistoricalvolatility(retrieveDataForTokensFromDatabase(token0symbol, token1symbol, database), 30)

			ROI_raw_est := calculateROI_raw_est(currentInterestrate, UniswapRewardPercentage, float32(future_pool_sz_est), float32(future_daily_volume_est))
			ROI_vol_adj_est := calculateROI_vol_adj(ROI_raw_est, volatility)
			ROI_hist := calculateROI_hist(currentInterestrate, UniswapRewardPercentage, float32(future_pool_sz_est), float32(future_daily_volume_est))

			if math.IsNaN(float64(ROI_raw_est)) {
				fmt.Println("NAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAN1")
			}
			if math.IsNaN(float64(ROI_vol_adj_est)) {
				fmt.Println("NAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAN1")
			}
			if math.IsNaN(float64(ROI_hist)) {
				fmt.Println("NAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAN1")
			}

			var recordalreadyexists bool
			recordalreadyexists = false

			for k := 0; k < len(database.currencyinputdata); k++ {
				// Means record already exists - UPDATE IT, DO NOT APPEND
				if database.currencyinputdata[k].Pair == token0symbol+"/"+token1symbol && database.currencyinputdata[k].Pool == "Balancer" {
					recordalreadyexists = true
					database.currencyinputdata[k].PoolSize = float32(future_pool_sz_est)
					database.currencyinputdata[k].PoolVolume = float32(future_daily_volume_est)

					database.currencyinputdata[k].ROI_raw_est = ROI_raw_est
					database.currencyinputdata[k].ROI_vol_adj_est = ROI_vol_adj_est
					database.currencyinputdata[k].ROI_hist = ROI_hist

					database.currencyinputdata[k].Volatility = volatility
					database.currencyinputdata[k].Yield = currentInterestrate
				}
			}

			// APPEND IF NEW
			if !recordalreadyexists {
				database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{token0symbol + "/" + token1symbol, float32(future_pool_sz_est),
					float32(future_daily_volume_est), currentInterestrate, "Uniswap", volatility, ROI_raw_est, ROI_vol_adj_est, ROI_hist})
				//				database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{token0symbol + "/" + token1symbol, float32(currentSize), float32(currentVolume), currentInterestrate, "Uniswap", volatility, ROI})
			}
		} // if pool is within pre filtered list ends
		// } // if pool has some tokens ends
	} // Uniswap pair loop closes

}
