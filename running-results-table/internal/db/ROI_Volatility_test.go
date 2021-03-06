package db

import "testing"

/*
func TestCalculateHistoricalVolatility(t *testing.T) {
	historical_data := HistoricalCurrencyData{}
	if(historical_data.Price[0] != 420.69) {
		t.Errorf("Pool error!")
	}

	historical_volatility := calculatehistoricalvolatility(historical_data, 1)


}*/
func TestSetUniswapQueryIDForToken(t *testing.T) {

	if setUniswapQueryIDForToken("DAI", "id") != "0x6b175474e89094c44da98b954eedeac495271d0f" {

		t.Errorf("Address error!")

	}

	if setUniswapQueryIDForToken("USDC", "id") != "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48" {

		t.Errorf("Address error!")

	}

	if setUniswapQueryIDForToken("ETH", "id") != "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2" {

		t.Errorf("Address error!")

	}

	if setUniswapQueryIDForToken("WETH", "id") != "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2" {

		t.Errorf("Address error!")

	}

	if setUniswapQueryIDForToken("RANDOM", "id") != "id" {

		t.Errorf("Address error!")

	}

}

func TestConvBalancerToken(t *testing.T) {

	if convBalancerToken("ETH") != "WETH" {

		t.Errorf("Balancer Token error!")

	}

	if convBalancerToken("RANDOM") != "RANDOM" {

		t.Errorf("Balancer Token error!")

	}

}

func TestIsPoolPartOfFilter(t *testing.T) {
	result1 := isPoolPartOfFilter("WETH", "DAI")

	result2 := isPoolPartOfFilter("ETH", "DAI")

	result3 := isPoolPartOfFilter("ETH", "DAI")

	result4 := isPoolPartOfFilter("ETH", "DAI")

	if(result1 != true) {
		t.Errorf("Pool error!")
	}

	if(result2 != false) {
		t.Errorf("Pool error!")
	}

	if(result3 != false) {
		t.Errorf("Pool error!")
	}

	if(result4 != false) {
		t.Errorf("Pool error!")
	}
}

func TestStringInSlice(t *testing.T) {
	string_list := []string{"a", "b", "c", "d"}

	result := stringInSlice("a", string_list)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

