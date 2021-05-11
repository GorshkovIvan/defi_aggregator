package db

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type CurvePoolData struct {
	poolAddress common.Address
	//poolCurrentBalances [8]*big.Int
	assetAddresses     [8]common.Address
	assetDecimals      [8]*big.Int
	volumes            []*[8]*big.Int
	fees               []*[8]*big.Float
	balances           []*[8]*big.Int
	normalsiedBalances []*big.Float
}

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
