package db

// ---Curve---
type CurveQuery struct {
	Pools []CurvePool `json:"pools"`
}

type CurvePool struct {
	Address   string      `json:"address"`
	CoinCount int         `json:"coinCount"`
	A         string      `json:"A"`
	Fee       string      `json:"fee"`
	AdminFee  string      `json:"adminFee"`
	Balances  []string    `json:"balances"`
	Coins     []CurveCoin `json:"coins"`
}

type CurveCoin struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals string `json:"decimals"`
}
