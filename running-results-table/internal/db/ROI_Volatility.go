package db

func calculatehistoricalvolatility(H HistoricalCurrencyData, days int) float32 {
	vol := 0.05

	// math.Min(int(len(H.DailyTimeSeries)),
	for i := 0; i < days; i++ {
		// TO DO: calculate deviation
	}

	// placeholder for case where no historical data is not available
	if len(H.Date) == 0 {
		return -1
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

func isHistDataAlreadyDownloaded(token string) bool {
	// TO DO: Check pair for database
	return false
}

func retrieveDataForTokensFromDatabase(token0 string, token1 string, database *Database) HistoricalCurrencyData {
	// Check if token0token1 is in database
	// check if token1token0 is in database
	// respUniswapHist
	return NewHistoricalCurrencyData()
}

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

func convBalancerToken(t string) string {
	if t == "ETH" {
		return "WETH"
	}
	return t
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isPoolPartOfFilter(token0 string, token1 string) bool {
	// count occurences
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
