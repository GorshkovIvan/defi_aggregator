package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type UniswapData struct {
	Data struct {
		Pairs []struct {
			Token0 struct {
				Symbol string `json:"symbol"`
			} `json:"token0"`
			Token0Price string `json:"token0Price"`
			Token1      struct {
				Symbol string `json:"symbol"`
			} `json:"token1"`
			Token1Price string `json:"token1Price"`
		} `json:"pairs"`
	} `json:"data"`
}

type CompoundData struct {
	Data struct {
		Markets []struct {
			Symbol           string `json:"symbol"`
			BorrowRate       string `json:"borrowRate"`
			Cash             string `json:"cash"`
			CollateralFactor string `json:"collateralFactor"`
			ExchangeRate     string `json:"exchangeRate"`
		} `json:"markets"`
	} `json:"data"`
}

type BalancerData struct {
	Data struct {
		Pools []struct {
			PublicSwap bool   `json:"publicSwap"`
			SwapFee    string `json:"swapFee"`
			Tokens     []struct {
				Balance string `json:"balance"`
				Symbol  string `json:"symbol"`
			} `json:"tokens"`
			TokensList []string `json:"tokensList"`
		} `json:"pools"`
	} `json:"data"`
}

type AaveData struct {
	Data struct {
		Reserves []struct {
			ID            string `json:"id"`
			LiquidityRate string `json:"liquidityRate"`
			Name          string `json:"name"`
			Price         struct {
				ID string `json:"id"`
			} `json:"price"`
			StableBorrowRate   string `json:"stableBorrowRate"`
			VariableBorrowRate string `json:"variableBorrowRate"`
		} `json:"reserves"`
	} `json:"data"`
}

type CurveData struct {
	Data struct {
		Pool struct {
			AdminFee  string   `json:"adminFee"`
			Balances  []string `json:"balances"`
			PoolToken struct {
				ID     string `json:"id"`
				Symbol string `json:"symbol"`
			} `json:"poolToken"`
		} `json:"pool"`
	} `json:"data"`
}

func main() {
	// Uniswap
	jsonData1 := map[string]string{
		"query": `
		{
			pairs(first: 2) 
			{
			 token0{
			  symbol
				}
			 token1{
			  symbol
				}
			  token0Price
			  token1Price
			}
		}
        `,
	}

	//2-Compound
	jsonData2 := map[string]string{
		"query": `
			{
				markets(first: 2 orderBy:borrowRate orderDirection: desc) 
				{
					symbol
					borrowRate
					cash
				  	collateralFactor
				  	exchangeRate
				}
			}
			`,
	}

	//3-Balancer
	jsonData3 := map[string]string{
		"query": `
		{
			pools(first: 1, where: {publicSwap: true}) {
			  id
			  publicSwap
			  swapFee
			  tokensList
			  tokens {
				balance
				symbol
			  }
			}
		}
			`,
	}

	//4-Aave
	jsonData4 := map[string]string{
		"query": `
		{
			reserves (first: 3 orderBy:variableBorrowRate orderDirection:desc where: {
			  usageAsCollateralEnabled: true
			}) {
			  id
			  name
			  price {
				id
			  }
			  liquidityRate
			  variableBorrowRate
			  stableBorrowRate
			}
		  }
				`,
	}

	//5-Curve
	jsonData5 := map[string]string{
		"query": `
		{
			pool(id:"0x06364f10b501e868329afbc005b3492902d6c763"){
			  balances
			  adminFee
			  coins
			  poolToken {
				id
				symbol
			  }
			}
		  }
					`,
	}

	// Address
	uniswapaddress := "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"          // Liquidity
	compound_address := "https://api.thegraph.com/subgraphs/name/graphprotocol/compound-v2" // Interest
	balancer_address := "https://api.thegraph.com/subgraphs/name/balancer-labs/balancer"    // liquidity
	aave_address := "https://api.thegraph.com/subgraphs/name/aave/protocol"                 // Interest
	curve_address := "https://api.thegraph.com/subgraphs/name/protofire/curve"
	/*
		bancor_address := "https://api.thegraph.com/subgraphs/name/blocklytics/bancor"
	*/

	// 1-Uniswap
	jsonValue, _ := json.Marshal(jsonData1)
	request, err := http.NewRequest("POST", uniswapaddress, bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	var datastr string
	datastr = string(data)
	var dx UniswapData // Data
	err = json.Unmarshal([]byte(datastr), &dx)

	fmt.Println("Uniswap data:")
	fmt.Printf("%+v", dx.Data.Pairs[0].Token0.Symbol)
	fmt.Printf("/%+v: ", dx.Data.Pairs[0].Token1.Symbol)
	fmt.Printf("%+v\n", dx.Data.Pairs[0].Token0Price)

	fmt.Printf("%+v", dx.Data.Pairs[1].Token0.Symbol)
	fmt.Printf("/%+v: ", dx.Data.Pairs[1].Token1.Symbol)
	fmt.Printf("%+v\n", dx.Data.Pairs[1].Token0Price)

	// 2 - Compound
	jsonValue2, _ := json.Marshal(jsonData2)
	request2, err := http.NewRequest("POST", compound_address, bytes.NewBuffer(jsonValue2))
	client2 := &http.Client{Timeout: time.Second * 10}
	response2, err := client2.Do(request2)
	if err != nil {
		fmt.Printf("Compound request failed with error %s\n", err)
	}

	defer response2.Body.Close()
	data2, _ := ioutil.ReadAll(response2.Body)
	fmt.Println("Compound:")

	var dCompound CompoundData
	err = json.Unmarshal([]byte(data2), &dCompound)
	if err != nil {
		fmt.Printf("Compound unmarshalling failed with error: %s\n", err)
	}

	// 3 - Balancer
	jsonValue3, _ := json.Marshal(jsonData3)
	request3, err := http.NewRequest("POST", balancer_address, bytes.NewBuffer(jsonValue3))
	client3 := &http.Client{Timeout: time.Second * 10}
	response3, err := client3.Do(request3)
	if err != nil {
		fmt.Printf("Balancer request failed with error %s\n", err)
	}

	defer response3.Body.Close()
	data3, _ := ioutil.ReadAll(response3.Body)
	fmt.Println("Balancer:")

	var dBalancer BalancerData
	err = json.Unmarshal([]byte(data3), &dBalancer)
	if err != nil {
		fmt.Printf("Balancer unmarshalling failed with error: %s\n", err)
	}

	fmt.Printf("%+v: ", dBalancer.Data.Pools[0].Tokens[0].Symbol)
	fmt.Printf("%+v", dBalancer.Data.Pools[0].Tokens[0].Balance)

	// 4 - Aave - Lending
	jsonValue4, _ := json.Marshal(jsonData4)
	request4, err := http.NewRequest("POST", aave_address, bytes.NewBuffer(jsonValue4))
	client4 := &http.Client{Timeout: time.Second * 10}
	response4, err := client4.Do(request4)
	if err != nil {
		fmt.Printf("Aave request failed with error %s\n", err)
	}

	defer response4.Body.Close()
	data4, _ := ioutil.ReadAll(response4.Body)
	fmt.Println("\nAave:")

	var dAave AaveData
	err = json.Unmarshal([]byte(data4), &dAave)
	if err != nil {
		fmt.Printf("Aave unmarshalling failed with error: %s\n", err)
	}

	fmt.Printf("%+v: ", len(dAave.Data.Reserves))
	fmt.Printf("%+v: ", dAave.Data.Reserves[0].Name)
	fmt.Printf("%+v: \n", dAave.Data.Reserves[0].VariableBorrowRate)
	fmt.Printf("%+v: ", dAave.Data.Reserves[1].Name)
	fmt.Printf("%+v: ", dAave.Data.Reserves[1].VariableBorrowRate)
	// 5 - Curve
	jsonValue5, _ := json.Marshal(jsonData5)
	request5, err := http.NewRequest("POST", curve_address, bytes.NewBuffer(jsonValue5))
	client5 := &http.Client{Timeout: time.Second * 10}
	response5, err := client5.Do(request5)
	if err != nil {
		fmt.Printf("Curve request failed with error %s\n", err)
	}

	defer response5.Body.Close()
	data5, _ := ioutil.ReadAll(response5.Body)
	fmt.Println("\nCurve:")

	var dCurve CurveData
	err = json.Unmarshal([]byte(data5), &dCurve)
	if err != nil {
		fmt.Printf("Curve unmarshalling failed with error: %s\n", err)
	}

	fmt.Printf("%+v\n", dCurve.Data.Pool.PoolToken.Symbol)
	fmt.Printf("%+v", dCurve.Data.Pool.Balances[0])
}

// dec := json.NewDecoder(bytes.NewReader(data))
// var dd map[string]string
// dec.Decode(&dd)
