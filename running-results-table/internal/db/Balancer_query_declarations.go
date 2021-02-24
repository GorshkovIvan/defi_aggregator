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
	Finalized       bool            `json:"finalized"`
	PublicSwap      bool            `json:"publicSwap"`
	SwapFee         string          `json:"swapFee"`
	TotalSwapVolume string          `json:"totalSwapVolume"`
	TotalWeight     string          `json:"totalWeight"`
	TokensList      []string        `json:"tokensList"`
	Tokens          []BalancerToken `json:"tokens"`
}

type BalancerById struct {
	BalancerPool `json:"pool"`
}

type BalancerToken struct {
	ID           string `json:"id"`
	Address      string `json:"address"`
	Balance      string `json:"balance"`
	Decimals     int    `json:"decimals"`
	Symbol       string `json:"symbol"`
	DenormWeight string `json:"denormWeight"`
}
