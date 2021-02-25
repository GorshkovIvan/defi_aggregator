package db

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/machinebox/graphql"
)

func (database *Database) AddRecordfromAPI() {

	// 1 - create clients
	clientBalancer := graphql.NewClient("https://api.thegraph.com/subgraphs/name/balancer-labs/balancer")
	clientUniswap := graphql.NewClient("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2")
	//clientAave := graphql.NewClient("https://api.thegraph.com/subgraphs/name/aave/protocol")

	clientCompound := graphql.NewClient("https://api.thegraph.com/subgraphs/name/graphprotocol/compound-v2")
	clientCurve := graphql.NewClient("https://api.thegraph.com/subgraphs/name/protofire/curve")
	clientBancor := graphql.NewClient("https://api.thegraph.com/subgraphs/name/blocklytics/bancor")

	// 2 - declare queries
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

	reqBalancerByPoolID := graphql.NewRequest(`
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
	reqBalancerListOfPools.Var("key", "value")
	reqUniswapListOfPools.Var("key", "value")

	reqAave.Var("address", "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee") // Ether address
	reqCompound.Var("key", "value")
	reqCurve.Var("key", "value")
	reqBancor.Var("key", "value")

	// 4 - set query headers
	reqBalancerListOfPools.Header.Set("Cache-Control", "no-cache")
	reqBalancerByPoolID.Header.Set("Cache-Control", "no-cache")

	reqUniswapListOfPools.Header.Set("Cache-Control", "no-cache")
	reqUniswapIDFromTokenTicker.Header.Set("Cache-Control", "no-cache")
	reqUniswapHist.Header.Set("Cache-Control", "no-cache")

	reqCompound.Header.Set("Cache-Control", "no-cache")
	reqCurve.Header.Set("Cache-Control", "no-cache")
	reqBancor.Header.Set("Cache-Control", "no-cache")
	reqAave.Header.Set("Cache-Control", "no-cache")

	// 5 - define a Context for the request
	ctx := context.Background()

	// 6 - declare query response objects
	var respBalancerPoolList BalancerPoolList
	var respBalancerById BalancerById

	var respUniswapPoolList UniswapPoolList
	var respUniswapTicker UniswapTickerQuery // Used in Balancer to look up Uniswap IDs of 'ETH' etc
	var respUniswapHist UniswapHistQuery
	var respUniswapById UniswapCurrentQuery

	var respCompound CompoundQuery // need to change to same format as Balancer and Uniswap
	var respCurve CurveQuery
	var respBancor BancorQuery
	// var respAave AaveQuery

	// 7 - run data queries on each pool
	/*
		Order:
		1) Balancer
		2) Uniswap
		3) Aave
		4) Curve
		5) Others
	*/
	// 7a - BALANCER
	var BalancerFilteredPoolList []string      // Pairs - IDS - 0x124145
	var BalancerFilteredPoolListPairs []string // Pairs - Tokens ETH/DAI
	var BalancerFilteredTokenList []string     // Tokens - ETH, DAI

	if err := clientBalancer.Run(ctx, reqBalancerListOfPools, &respBalancerPoolList); err != nil {
		log.Fatal(err)
	}

	// Process received list of pools (PAIRS)
	for i := 0; i < len(respBalancerPoolList.Pools); i++ {
		if len(respBalancerPoolList.Pools[i].Tokens) > 1 {
			token0symbol := respBalancerPoolList.Pools[i].Tokens[0].Symbol
			token1symbol := respBalancerPoolList.Pools[i].Tokens[1].Symbol

			if isPoolPartOfFilter(token0symbol, token1symbol) {
				// Filter pools to allowed components (WETH, DAI, USDC, USDT)
				BalancerFilteredPoolList = append(BalancerFilteredPoolList, respBalancerPoolList.Pools[i].ID)
				BalancerFilteredPoolListPairs = append(BalancerFilteredPoolListPairs, token0symbol+"/"+token1symbol)

				var tokenqueue []string

				// Split list of pairs into single tokens
				if !stringInSlice(token0symbol, BalancerFilteredTokenList) {
					BalancerFilteredTokenList = append(BalancerFilteredTokenList, token0symbol)
					tokenqueue = append(tokenqueue, token0symbol)
				}
				if !stringInSlice(token1symbol, BalancerFilteredTokenList) {
					BalancerFilteredTokenList = append(BalancerFilteredTokenList, token1symbol)
					tokenqueue = append(tokenqueue, token1symbol)
				}

				/*
					fmt.Print("tokenqueue len: ")
					fmt.Println(len(tokenqueue))
				*/

				for j := 0; j < len(tokenqueue); j++ {
					// Check if database already has historical data
					if !isHistDataAlreadyDownloaded(convBalancerToken(tokenqueue[j]), database) {
						// Get Uniswap Ids of these tokens
						reqUniswapIDFromTokenTicker.Var("ticker", convBalancerToken(tokenqueue[j]))
						if err := clientUniswap.Run(ctx, reqUniswapIDFromTokenTicker, &respUniswapTicker); err != nil {
							log.Fatal(err)
						}
						// Download historical data for each token for which data is missing
						if len(respUniswapTicker.IDsforticker) >= 1 {
							// request data from uniswap using this queried ticker
							reqUniswapHist.Var("tokenid", setUniswapQueryIDForToken(tokenqueue[j], respUniswapTicker.IDsforticker[0].ID))

							fmt.Print("Querying historical data for: ")
							fmt.Print(tokenqueue[j])
							if err := clientUniswap.Run(ctx, reqUniswapHist, &respUniswapHist); err != nil {
								log.Fatal(err)
							}

							fmt.Print("| returned days: ")
							fmt.Println(len(respUniswapHist.DailyTimeSeries))

							// if returned data - append it to database
							if len(respUniswapHist.DailyTimeSeries) > 0 {
								// Append to database
								database.historicalcurrencydata = append(database.historicalcurrencydata, NewHistoricalCurrencyDataFromRaw(tokenqueue[j], respUniswapHist.DailyTimeSeries))
							}
						} // if managed to find some IDs for this TOKEN
					} // if historical data needs updating
				} // tokenqueue loop ends

				// if historical data is in order - get current data
				reqBalancerByPoolID.Var("poolid", respBalancerPoolList.Pools[i].ID)

				if err := clientBalancer.Run(ctx, reqBalancerByPoolID, &respBalancerById); err != nil {
					log.Fatal(err)
				}

				currentSize := float32(10000.0)                                              // TO DO: Size
				currentVolume, _ := strconv.ParseFloat(respBalancerById.TotalSwapVolume, 32) // No historical for now
				currentInterestrate := float32(0.00)                                         // Zero for liquidity pool
				BalancerRewardPercentage := float32(0.003)                                   // TO DO

				volatility := calculatehistoricalvolatility(retrieveDataForTokensFromDatabase(token0symbol, token1symbol, database), 30)
				ROI := calculateROI(currentInterestrate, BalancerRewardPercentage, float32(currentVolume), volatility)

				database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{token0symbol + "/" + token1symbol, float32(currentSize), float32(currentVolume), currentInterestrate, "Balancer", volatility, ROI})

			} // if pool is within pre filtered list ends
		} // if pool has some tokens ends
	} // balancer pair loop closes

	// 7b - UNISWAP
	var UniswapFilteredPoolList []string      // Pairs - IDS - 0x124145
	var UniswapFilteredPoolListPairs []string // Pairs - Tokens ETH/DAI
	var UniswapFilteredTokenList []string     // Tokens - ETH, DAI

	if err := clientUniswap.Run(ctx, reqUniswapListOfPools, &respUniswapPoolList); err != nil {
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
					reqUniswapHist.Var("tokenid", tokenqueueIDs[j])
					fmt.Print("Querying historical data for: ")
					fmt.Print(tokenqueueIDs[j])
					fmt.Print(" : ")
					fmt.Print(tokenqueue[j])
					if err := clientUniswap.Run(ctx, reqUniswapHist, &respUniswapHist); err != nil {
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

			if err := clientUniswap.Run(ctx, reqUniswapByPoolID, &respUniswapById); err != nil {
				log.Fatal(err)
			}

			currentSize := float32(1000.000)
			currentVolume, _ := strconv.ParseFloat(respUniswapById.Pair.VolumeUSD, 32) // No historical for now
			currentInterestrate := float32(0.00)                                       // Zero for liquidity pool
			UniswapRewardPercentage := float32(0.003)                                  // Placeholder

			volatility := calculatehistoricalvolatility(retrieveDataForTokensFromDatabase(token0symbol, token1symbol, database), 30)
			ROI := calculateROI(currentInterestrate, UniswapRewardPercentage, float32(currentVolume), volatility)

			database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{token0symbol + "/" + token1symbol, float32(currentSize), float32(currentVolume), currentInterestrate, "Uniswap", volatility, ROI})
		} // if pool is within pre filtered list ends
		// } // if pool has some tokens ends
	} // Uniswap pair loop closes

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

	// AAVE
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
