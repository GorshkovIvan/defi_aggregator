package db

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// Current Exchange Rate Comes from the protocol
// Standard deviation comes from the volatility estimate, can be a 30 days estimate
// returned estimate is -x% loss in liquidity
func estimate_impermanent_loss_hist(standard_deviation float32, current_exchange_rate float32, protocol string) float32 {

	impermanent_loss := float64(0)

	if protocol == "Uniswap" {

		forecasted_exchage_rate := current_exchange_rate + standard_deviation
		price_ratio := forecasted_exchage_rate / current_exchange_rate
		impermanent_loss = 2*math.Sqrt(float64(price_ratio))/(1+float64(price_ratio)) - 1

	}

	return float32(impermanent_loss)

}

func calculate_price_return_x_days(hist_date_px_series HistoricalCurrencyData, days int) float32 {

	return 0.0
}

func calculatehistoricalvolatility(H HistoricalCurrencyData, days int) float32 {
	fmt.Print("CALCULATING HISTORICAL VOLATILITY: ")
	fmt.Println(H.Ticker)

	var vol float32
	vol = 0.05

	if len(H.Price) == 0 {
		fmt.Print("Error: no historical data found for ticker: ")
		//fmt.Print(H.Ticker)
		//fmt.Print(" ..returning -1 for volatility..")
		return -1
	}

	var vol_period int32

	// return -1 if no historical data available
	if len(H.Date) == 0 {
		vol = -1.00
	}

	vol_period = int32(math.Min(float64(len(H.Price)), float64(days))) // lower of days or available data

	// NOTE: oldest = 0
	var total float32
	total = 0.00

	fmt.Print("Vol period: ")
	fmt.Println(vol_period)

	var changes_in_price []float32
	var differencesvsmean []float32 // size = actual vol period
	var squaresofdifferencesvsmean []float32

	var actual_vol_period int32 // days with data
	actual_vol_period = 0

	if vol_period < 2 {
		return 0.0
	}

	for i := 1; i < int(vol_period); i++ {
		if !math.IsNaN(float64(H.Price[i])) && float64(H.Price[i]) > 0 && float64(H.Price[i-1]) > 0 {
			changes_in_price = append(changes_in_price, H.Price[i]/H.Price[i-1]-1)
			total = total + (H.Price[i]/H.Price[i-1] - 1) // calculate average price change
			actual_vol_period++
		}
	}

	fmt.Print("total: ")
	fmt.Println(total)

	mean := total / float32(actual_vol_period) // actual days?

	fmt.Print("mean: ")
	fmt.Println(mean)

	for i := 1; i < int(vol_period); i++ {
		if !math.IsNaN(float64(H.Price[i])) && float64(H.Price[i]) > 0 && float64(H.Price[i-1]) > 0 {
			differencesvsmean = append(differencesvsmean, H.Price[i]/H.Price[i-1]-1-mean) // calculate difference between each value and mean
			fmt.Print("Date: ")
			fmt.Print(H.Date[i])
			fmt.Print(" | ")
			fmt.Print("Price: ")
			fmt.Print(H.Price[i])
			fmt.Print(" | ")
			fmt.Print("Price - mean: ")
			fmt.Print(H.Price[i] - mean)
			squaresofdifferencesvsmean = append(squaresofdifferencesvsmean, float32(math.Pow(float64(H.Price[i]/H.Price[i-1]-1-mean), 2.0)))
			fmt.Print(" | Sqr: ")
			fmt.Println(float32(math.Pow(float64(H.Price[i]/H.Price[i-1]-1-mean), 2.0)))
		}
	}

	var avg float32
	avg = 0.0

	for i := 0; i < len(squaresofdifferencesvsmean); i++ {
		avg += squaresofdifferencesvsmean[i]
	}

	fmt.Print("Total squares: ")
	fmt.Println(avg)

	avg = avg / float32(len(squaresofdifferencesvsmean))    // average them
	vol = float32(math.Sqrt(float64(avg)) * math.Sqrt(252)) // is this the right adjustment for days?

	fmt.Print("CALCULATED VOLATILITY = ")
	fmt.Println(vol)

	if math.IsInf(float64(vol), 0) {
		return -0.99
	}
	if math.IsNaN(float64(vol)) {
		return -0.98
	}

	return float32(vol)
}

func calculateROI_hist(interestrate float32, pool_reward_pct float32, pool_sz_hist float32, daily_volume_hist float32, imp_loss_hist_est float32, px_return_hist float32) float32 {
	var ROI float32
	ROI = 0.069

	if pool_sz_hist == 0.0 {
		return -992.0
	} else {
		ROI = interestrate + (pool_reward_pct * daily_volume_hist * 365 / pool_sz_hist) + px_return_hist + imp_loss_hist_est
	}

	if math.IsInf(float64(ROI), 0) {
		return -999
	}

	if math.IsNaN(float64(ROI)) {
		return -998
	}

	return float32(ROI)
}

// Sharpe ratio
func calculateROI_vol_adj(ROI_raw_est float32, volatility_est float32) float32 {

	if math.IsInf(float64(ROI_raw_est), 0) {
		return 888.0
	}

	if volatility_est <= 0.03 { // if not volatile - do not adjust by volatility
		return ROI_raw_est
	} else {
		if !math.IsInf(float64(ROI_raw_est/volatility_est), 0) {
			return ROI_raw_est / volatility_est // sharpe ratio - risk free rate assumed to be zero
		} else {
			return 888.8
		}

	}

}

func calculateROI_raw_est(interestrate float32, pool_reward_pct float32, future_pool_sz_est float32, future_daily_volume_est float32, imp_loss_hist_est float32) float32 {

	var ROI float32
	ROI = 0.069

	if future_pool_sz_est > 0.0 {
		ROI = interestrate + (pool_reward_pct * future_daily_volume_est * 365 / future_pool_sz_est) + imp_loss_hist_est
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
	fmt.Print("RETRIEVING DATA FOR PAIR - Tokens: ")
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
			fmt.Print(" | token0: ")
			fmt.Print(token0)
			fmt.Print(" | found @ idx = ")
			fmt.Print(token0idx)
			fmt.Print(" | ")
			break
		}
	}

	for i := 0; i < len(database.historicalcurrencydata); i++ {
		if database.historicalcurrencydata[i].Ticker == token1 {
			token1dataishere = true
			token1idx = i
			fmt.Print(" | token1: ")
			fmt.Print(token1)
			fmt.Print(" | Found token data 0 at idx = ")
			fmt.Println(token1idx)
			break
		}
	}

	if !token0dataishere || !token1dataishere {
		fmt.Println("ERROR 899: ticker combo not found in database..returning blank object")
		return NewHistoricalCurrencyData()
	}

	var histcombo HistoricalCurrencyData

	histcombo.Ticker = token0 + "/" + token1

	//fmt.Println(" | ")
	//fmt.Print("Created ticker pair for hist combo")
	//fmt.Println(histcombo.Ticker)

	lengthoflookbackhist := int64(math.Min(float64(len(database.historicalcurrencydata[token0idx].Date)), float64(len(database.historicalcurrencydata[token0idx].Date))))
	lengthoflookbackhist2 := int64(math.Min(float64(len(database.historicalcurrencydata[token1idx].Date)), float64(len(database.historicalcurrencydata[token1idx].Date))))
	lengthoflookbackhist = int64(math.Min(float64(lengthoflookbackhist), float64(lengthoflookbackhist2)))

	fmt.Print("length of lookback = ")
	fmt.Println(lengthoflookbackhist)

	for i = lengthoflookbackhist - 1; i >= 0; i-- {
		// which index 	// 0 = oldest
		// add check if dates consistent across 2 datasets
		tm := time.Unix(database.historicalcurrencydata[token0idx].Date[i], 0)
		fmt.Print("i: ")
		fmt.Print(database.historicalcurrencydata[token0idx].Date[i])
		fmt.Print("parsed time: ")
		fmt.Print(tm)
		fmt.Print(" | px0: ")
		fmt.Print(database.historicalcurrencydata[token0idx].Price[i])
		fmt.Print(" | px1: ")
		fmt.Println(database.historicalcurrencydata[token1idx].Price[i])

		histcombo.Date = append(histcombo.Date, database.historicalcurrencydata[token0idx].Date[i])

		var price float32
		if database.historicalcurrencydata[token1idx].Price[i] > 0 {
			price = database.historicalcurrencydata[token0idx].Price[i] / database.historicalcurrencydata[token1idx].Price[i]
		} else {
			price = 0.0
		}
		if math.IsInf(float64(price), 0) {
			price = 0.0
			fmt.Println("WARNING 987: Inf in calculating token combo price")
		}
		if math.IsNaN(float64(price)) {
			price = 0.0
			fmt.Println("WARNING 987: Inf in calculating token combo price")
		}

		histcombo.Price = append(histcombo.Price, price)
	}

	fmt.Print("SIZE of returned combo for ticker: ")
	fmt.Print(histcombo.Ticker)
	fmt.Print(": ")
	fmt.Println(len(histcombo.Price))

	if len(histcombo.Price) >= 2 {
		fmt.Print(histcombo.Date[0])
		fmt.Print(" | ")
		fmt.Println(histcombo.Price[0])
		fmt.Print(histcombo.Date[1])
		fmt.Print(" | ")
		fmt.Print(histcombo.Price[1])
	}

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
		return database.currencyinputdata[i].ROI_raw_est > database.currencyinputdata[j].ROI_raw_est
	})
}
