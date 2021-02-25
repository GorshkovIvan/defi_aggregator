package db

import (
	"math"
	"sort"
)

func calculatehistoricalvolatility(H HistoricalCurrencyData, days int) float32 {
	var vol float32
	vol = 0.05

	var vol_period int32
	vol_period = int32(math.Min(float64(len(H.Price)), float64(days))) // lower of days or available data
	// NOTE: latest is first ? reverse indices if yet
	// TO DO: calculate deviation
	var total float32
	total = 0.00

	for i := 0; i < int(vol_period); i++ {
		total += H.Price[i] // calculate average price
	}

	mean := total / float32(vol_period) // actual days?
	var differencesvsmean []float32
	for i := 0; i < int(vol_period); i++ {
		differencesvsmean = append(differencesvsmean, H.Price[i]-mean) // calculate difference between each value and mean
	}

	var squaresofdifferencesvsmean []float32
	for i := 0; i < int(vol_period); i++ {
		// square these values
		squaresofdifferencesvsmean = append(squaresofdifferencesvsmean, float32(math.Pow(float64(differencesvsmean[i]), 2.0)))
	}

	var avg float32
	avg = 0.0
	for i := 0; i < int(vol_period); i++ {
		avg += squaresofdifferencesvsmean[i]
	}

	avg = avg / float32(vol_period)                         // average them
	vol = float32(math.Sqrt(float64(avg)) * math.Sqrt(252)) // is this the right adjustment for days?

	// return -1 if no historical data available
	if len(H.Date) == 0 {
		vol = -1.00
	}

	return float32(vol)
}

func calculateROI(interestrate float32, shareofvolume float32, poolvolume float32, volatility float32) float32 {

	var ROI float32
	ROI = 0.069
	// TO DO: update to proper formula
	ROI = (float32(interestrate) + float32(poolvolume)*float32(shareofvolume)) / float32(volatility)

	return float32(ROI)
}

func isHistDataAlreadyDownloaded(token string, database *Database) bool {

	for i := 0; i < len(database.historicalcurrencydata); i++ {
		if database.historicalcurrencydata[i].Ticker == token {
			// TO DO: also add date check: if latest date is within 24 hours of NOW database.historicalcurrencydata[i].
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
	// TO DO
	// Check if token0token1 is in database
	// check if token1token0 is in database
	// respUniswapHist
	return NewHistoricalCurrencyData()
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
