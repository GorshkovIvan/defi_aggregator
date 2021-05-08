package db

// ---Balancer---
type BalancerPoolID struct {
	ID         string          `json:"id"`
	TokensList []string        `json:"tokensList"`
	Tokens     []BalancerToken `json:"tokens"`
}

type BalancerPoolList struct {
	Pools []BalancerPoolID `json:"pools"`
}

type BalancerQuery struct {
	Pools []BalancerPool `json:"pools"`
}

type BalancerPool struct {
	ID              string          `json:"id"`

	SwapFee         string          `json:"swapFee"`
	Liquidity		string          `json:"liquidity"`
	TotalWeight		string          `json:"totalWeight"`
	TotalSwapVolume	string 			`json:"totalSwapVolume"`

	TokensList      []string        `json:"tokensList"`
	Tokens          []BalancerToken `json:"tokens"`
}

type BalancerById struct { // wrapper struct
	Pool 	BalancerPool `json:"pool"`
}

type BalancerToken struct {
	ID           string `json:"id"`
	Address      string `json:"address"`
	Balance      string `json:"balance"`
	Symbol       string `json:"symbol"`
}

type PoolWrapper struct {
	Swaps []BalancerHistVolumeSingleT `json:"swaps"`
}

type BalancerHistVolumeQuery struct {
	Pool 	PoolWrapper `json:"pool"`
}

type BalancerHistVolumeSingleT struct {
	Timestamp		int 		`json:"timestamp"`
	FeeValue		string 		`json:"feeValue"`
	TokenInSym		string 		`json:"tokenInSym"`
	TokenOutSym		string 		`json:"tokenOutSym"`
	TokenIn			string 		`json:"tokenIn"`
	TokenOut		string 		`json:"tokenOut"`
	TokenAmountIn	string 		`json:"tokenAmountIn"`
	TokenAmountOut 	string 		`json:"tokenAmountOut"`
	PoolLiquidity	string 		`json:"poolLiquidity"`
}