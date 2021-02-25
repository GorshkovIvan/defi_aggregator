package db

// ---Bancor---
type BancorQuery struct {
	Swaps []BancorSwap `json:"swaps"`
}

type BancorSwap struct {
	ID string `json:"id"`
	// Need separate structs for these datatypes
	//FromToken BancorToken `json:"fromToken"`
	//ToToken BancorToken `json:"toToken"`
	AmountPurchased                     string `json:"amountPurchased"`
	AmountReturned                      string `json:"amountReturned"`
	Price                               string `json:"price"`
	InversePrice                        string `json:"inversePrice"`
	ConverterWeight                     string `json:"converterWeight"`
	ConverterFromTokenBalanceBeforeSwap string `json:"converterFromTokenBalanceBeforeSwap"`
	ConverterFromTokenBalanceAfterSwap  string `json:"converterFromTokenBalanceAfterSwap"`
	ConverterToTokenBalanceBeforeSwap   string `json:"converterToTokenBalanceBeforeSwap"`
	ConverterToTokenBalanceAfterSwap    string `json:"converterToTokenBalanceAfterSwap"`
	Slippage                            string `json:"slippage"`
	ConversionFee                       string `json:"conversionFee"`
	// Need separate structs for these datatypes
	//ConverterUsed Converter `json:"converterUsed"`
	//Transaction Transaction `json:"transaction"`
	//Trader User `json:"trader"`
	Timestamp string `json:"timestamp"`
	LogIndex  int    `json:"logIndex"`
}

// --AAVE--
/*
type AaveQuery struct {
	Reserve AaveData `json:"reserve"`
}

type AaveData struct {
	ID                 string `json:"id"`
	Symbol             string `json:"symbol"`
	LiquidityRate      string `json:"liquidityRate"`
	StableBorrowRate   string `json:"stableBorrowRate"`
	VariableBorrowRate string `json:"variableBorrowRate"`
	TotalBorrows       string `json:"totalBorrows"`
}
*/
