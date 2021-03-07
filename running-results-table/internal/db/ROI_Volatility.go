package db

import (
	"fmt"
	"math"
	"sort"
)

func calculatehistoricalvolatility(H HistoricalCurrencyData, days int) float32 {
	fmt.Println("Entered calculation of HISTORICAL VOLATILITY: ")

	var vol float32
	vol = 0.05

	if len(H.Price) == 0 {
		fmt.Print("Error: no historical data found for ticker: ")
		fmt.Print(H.Ticker)
		fmt.Print(" ..returning -1 for volatility..")
		return -1
	}

	var vol_period int32
	vol_period = int32(math.Min(float64(len(H.Price)), float64(days))) // lower of days or available data
	// check how many days non NaN

	// NOTE: oldest = 0
	var total float32
	total = 0.00

	//fmt.Print("Vol period: ")
	//fmt.Println(vol_period)

	var actual_vol_period int32
	actual_vol_period = 0

	for i := 0; i < int(vol_period); i++ {
		if !math.IsNaN(float64(H.Price[i])) {
			total = total + H.Price[i] // calculate average price
			actual_vol_period++
			//	fmt.Print(i)
			//	fmt.Print(" : ")
			//	fmt.Print(H.Date[i])
			//	fmt.Print(" : ")
			//	fmt.Println(H.Price[i])
		}
	}

	//fmt.Println("total: ")
	//fmt.Print(total)

	mean := total / float32(actual_vol_period) // actual days?

	fmt.Println("mean: ")
	fmt.Print(mean)

	var differencesvsmean []float32 // size = actual vol period
	for i := 0; i < int(vol_period); i++ {
		if !math.IsNaN(float64(H.Price[i])) {
			differencesvsmean = append(differencesvsmean, H.Price[i]-mean) // calculate difference between each value and mean
		}
	}

	fmt.Println("Checkpoint 2 in volatility func")

	var squaresofdifferencesvsmean []float32
	for i := 0; i < len(differencesvsmean); i++ {
		// square these values
		squaresofdifferencesvsmean = append(squaresofdifferencesvsmean, float32(math.Pow(float64(differencesvsmean[i]), 2.0)))
	}

	var avg float32
	avg = 0.0
	for i := 0; i < len(squaresofdifferencesvsmean); i++ {
		// H.Price[i] != math.NaN
		avg += squaresofdifferencesvsmean[i]
	}

	avg = avg / float32(vol_period)                         // average them
	vol = float32(math.Sqrt(float64(avg)) * math.Sqrt(252)) // is this the right adjustment for days?

	// return -1 if no historical data available
	if len(H.Date) == 0 {
		vol = -1.00
	}

	fmt.Println("RETURNING VOLATILITY: ")
	fmt.Print(vol)

	if math.IsInf(float64(vol), 0) {
		return -0.99
	}
	if math.IsNaN(float64(vol)) {
		return -0.98
	}

	return float32(vol)
}

func calculateROI(interestrate float32, shareofvolume float32, poolvolume float32, volatility float32) float32 {

	var ROI float32
	ROI = 0.069
	// TO DO: update to proper formula
	if volatility > 0 {
		ROI = (float32(interestrate) + float32(poolvolume)*float32(shareofvolume)) / float32(volatility) / 365 / 1000
	}
	if volatility == 0 {
		ROI = (float32(interestrate) + float32(poolvolume)*float32(shareofvolume)) / 365 / 1000
	}

	if math.IsInf(float64(ROI), 0) {
		return -999
	}
	if math.IsNaN(float64(ROI)) {
		return -998
	}

	return float32(ROI)
}

func isHistDataAlreadyDownloaded(token string, database *Database) bool {

	for i := 0; i < len(database.historicalcurrencydata); i++ {
		if database.historicalcurrencydata[i].Ticker == token {
			// also add date check LATER : if latest date is within 24 hours of NOW database.historicalcurrencydata[i].
			/*
				fmt.Print("Checking if data already downloaded for: ")
				fmt.Print(token)
				fmt.Print("..Data found!!")
			*/
			return true
		}
	}
	return false
}

func retrieveDataForTokensFromDatabase(token0 string, token1 string, database *Database) HistoricalCurrencyData {
	// respUniswapHist
	fmt.Print("Entered retrievedatafor Tokens for PAIR: ")
	fmt.Println(token0 + "/" + token1 + " : ")

	var i int64

	token0dataishere := false // bool
	token1dataishere := false

	token0idx := 0
	token1idx := 1

	fmt.Println("CURRENT SHAPE OF THE DATABASE: ")
	for i := 0; i < len(database.historicalcurrencydata); i++ {
		fmt.Print(database.historicalcurrencydata[i].Ticker)
		fmt.Print(" : ")
		fmt.Print(len(database.historicalcurrencydata[i].Date))
		fmt.Print(" : ")
		fmt.Println(len(database.historicalcurrencydata[i].Price))
	}

	fmt.Println(" ")
	fmt.Println("_____________________________________")

	for i := 0; i < len(database.historicalcurrencydata); i++ {
		if database.historicalcurrencydata[i].Ticker == token0 {
			token0dataishere = true
			token0idx = i
			fmt.Println(" | token: ")
			fmt.Print(token0)
			fmt.Print(" | found @ idx = ")
			fmt.Println(token0idx)
			break
		}
	}

	for i := 0; i < len(database.historicalcurrencydata); i++ {
		if database.historicalcurrencydata[i].Ticker == token0 {
			token1dataishere = true
			token1idx = i
			fmt.Print(" | token: ")
			fmt.Print(token1)
			fmt.Print(" | Found token data 0 at idx = ")
			fmt.Println(token1idx)
			break
		}
	}

	if !token0dataishere || !token1dataishere {
		fmt.Println("Error: ticker combo not found in database..returning blank object")
		return NewHistoricalCurrencyData()
	}

	var histcombo HistoricalCurrencyData

	histcombo.Ticker = token0 + "/" + token1

	//fmt.Println(" | ")
	//fmt.Print("Created ticker pair for hist combo")
	//fmt.Println(histcombo.Ticker)

	lengthoflookbackhist := int64(math.Min(float64(len(database.historicalcurrencydata[token0idx].Date)), float64(len(database.historicalcurrencydata[token0idx].Date))))

	fmt.Print("length of lookback = ")
	fmt.Print(lengthoflookbackhist)

	for i = lengthoflookbackhist - 1; i >= 0; i-- {
		// which index 	// 0 = oldest
		// add check if dates consistent across 2 datasets
		histcombo.Date = append(histcombo.Date, database.historicalcurrencydata[token0idx].Date[i]/(60*60*24))
		price := database.historicalcurrencydata[token0idx].Price[i] / database.historicalcurrencydata[token1idx].Price[i]
		histcombo.Price = append(histcombo.Price, price)
	}

	fmt.Print("In histcombo: size of returned combo for ticker: ")
	fmt.Print(histcombo.Ticker)
	fmt.Print(": ")
	fmt.Println(len(histcombo.Price))

	return histcombo
}

// For looking up Uniswap IDs of tokens
func setUniswapQueryIDForToken(token string, ID string) string {
	if token == "DAI" {
		return "0x6b175474e89094c44da98b954eedeac495271d0f"
	}
	if token == "USDC" {
		// TO CHECK
		return "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48" // "check the uniswap id for dai"
	}
	if token == "ETH" || token == "WETH" {
		// TO CHECK
		return "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
	}
	return ID
}

// Convert Balancer token to a Uniswap token format
func convBalancerToken(t string) string {
	if t == "ETH" {
		return "WETH"
	}
	return t
}

// Checks if string is already in a vector
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Checks if a pair is within a pre-set list
func isPoolPartOfFilter(token0 string, token1 string) bool {
	t0 := "DAI"
	t1 := "USDC"
	t2 := "USDT"
	t3 := "WETH"
	t4 := "WBTC"

	var t0ok bool
	var t1ok bool

	if token0 == t0 || token0 == t1 || token0 == t2 || token0 == t3 || token0 == t4 {
		t0ok = true
	}

	if token1 == t0 || token1 == t1 || token1 == t2 || token1 == t3 || token1 == t4 {
		t1ok = true
	}

	if t0ok && t1ok {
		return true
	}
	return false

}

// ROI Ranking Function
func (database *Database) RankBestCurrencies() {

	sort.Slice(database.currencyinputdata, func(i, j int) bool {
		return database.currencyinputdata[i].ROIestimate > database.currencyinputdata[j].ROIestimate
	})
}
