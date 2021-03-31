package db

// To get list of pairs
type UniswapPoolList struct {
	Pools []UniswapPair `json:"pairs"`
}

type UniswapTokenDayData2 struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
}

type UniswapTickerQuery struct {
	IDsforticker []UniswapTokenDayData2 `json:"tokens"`
}

// Uniswap historical
type UniswapHistQuery struct {
	DailyTimeSeries []UniswapDaily `json:"tokenDayDatas"`
}

type UniswapDaily struct {
	Date     int                  `json:"date"`
	PriceUSD string               `json:"priceUSD"`
	Token    UniswapTokenDayData2 `json:"token"`
}

type UniswapCurrentQuery struct {
	Pair UniswapPair `json:"pair"`
}

type UniswapPair struct {
	ID                 string               `json:"id"`
	UntrackedVolumeUSD string               `json:"untrackedVolumeUSD"`
	VolumeUSD          string               `json:"volumeUSD"`
	Token0             UniswapTokenDayData2 `json:"token0"`
	Token1             UniswapTokenDayData2 `json:"token1"`
}

// Uniswap historical volume
type UniswapHistVolumeQuery struct {
	DailyTimeSeries []UniswapDailyVolume `json:"pairDayDatas"`
}

type UniswapDailyVolume struct {
	ID                	string               `json:"id"`
	Date              	int                  `json:"date"`
	Token0            	UniswapTokenDayData2 `json:"token0"`
	Token1            	UniswapTokenDayData2 `json:"token1"`
	DailyVolumeToken0 	string               `json:"dailyVolumeToken0"`
	DailyVolumeToken1 	string               `json:"dailyVolumeToken1"`
	DailyVolumeUSD 		string               `json:"dailyVolumeUSD"`
	
	TotalSupply 		string               `json:"totalSupply"`
	ReserveUSD 			string               `json:"reserveUSD"`
}

// Uniswap all pairs
type UniswapPairList struct {
	Pairs []UniswapTokenDayData3 `json:"pairs"`
}

type UniswapTokenDayData3 struct {
	ID     string               `json:"id"`
	Token0 UniswapTokenDayData2 `json:"token0"`
	Token1 UniswapTokenDayData2 `json:"token1"`
}
