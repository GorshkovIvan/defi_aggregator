package db

import (
	"fmt"
	"log"
	"math"
	"strings"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize"
	"gonum.org/v1/gonum/stat"
)

func sum(array []float64) float64 {
	result := 0.0
	for _, v := range array {
		result += v
	}
	return result
}


func OptimisePortfolio(database *Database) []OptimisedPortfolioRecord {
	fmt.Println("----ENTERING PORTFOLIO OPTIMISATION----------")

	var startingTokenTickers []string 
	var startingTokenAmounts []float32
	// var listOfAvailablePairswithConversion []string

	if database.Risksetting == 0 {
		fmt.Println("WARNING: Risk setting set to zero!")
	}
	
	// Get some data for testing locally
	/*
	database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{"WETH" + "/" + "DAI", float32(99),
					float32(98), 0.05, "Uniswap", 0.15, 0.16, 0.17, 0.18})
	database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{"WETH", float32(199),
					float32(198), 0.15, "Aave", 0.05, 0.06, 0.07, 0.08})
	database.currencyinputdata = append(database.currencyinputdata, CurrencyInputData{"USDC", float32(199),
					float32(198), 0.15, "Balancer", 0.05, 0.06, 0.07, 0.08})
	*/

	database.ownstartingportfolio = append(database.ownstartingportfolio,OwnPortfolioRecord{"WETH",420.0})
	database.ownstartingportfolio = append(database.ownstartingportfolio,OwnPortfolioRecord{"DAI",69.0})
		database.ownstartingportfolio = append(database.ownstartingportfolio,OwnPortfolioRecord{"USD",69.0})

	// 0 - Clean starting portfolio for duplicates
	for i := 0; i < len(database.ownstartingportfolio); i++ {
		if !stringInSlice(database.ownstartingportfolio[i].Token, startingTokenTickers) {
			startingTokenTickers = append(startingTokenTickers, database.ownstartingportfolio[i].Token)
		}
	}
	
	for i := 0; i < len(startingTokenTickers); i++ {
		for j:=0; j < len(database.ownstartingportfolio); j++ {
				if database.ownstartingportfolio[j].Token == startingTokenTickers[i]{
					startingTokenAmounts = append(startingTokenAmounts, database.ownstartingportfolio[j].Amount)			
				}
			}
	}

		fmt.Println("TRYING TO OPTIMISE PORTFOLIO: ")
		fmt.Print("risk: ")
		fmt.Println(database.Risksetting)
		
		for i := 0; i < len(startingTokenTickers); i++ {
			fmt.Print(startingTokenTickers[i])
			fmt.Print(" | ")
			fmt.Println(startingTokenAmounts[i])
		}
				
		fmt.Println("AVAILABLE POOLS TO DEPLOY INTO: ")
		for i := 0; i < len(database.currencyinputdata); i++ {		
			fmt.Print(database.currencyinputdata[i].Pool)
			fmt.Print(" | ")
			fmt.Print(database.currencyinputdata[i].Pair)
			fmt.Print(" | ")
			fmt.Print(database.currencyinputdata[i].ROI_raw_est)
			
			// for each pair - if pair is made up of portfolio tokens - add to available pool list
			s := strings.Split(database.currencyinputdata[i].Pair, "/")
			fmt.Print(" | Extracted tokens: ")
			fmt.Print(s[0])
			fmt.Print("   |   ")
			if len(s) > 1 {
				fmt.Print(s[1])
			}
			
			fmt.Print(" Available given portfolio?: ")
			if len(s) == 2 {
				fmt.Println(stringInSlice(s[0],startingTokenTickers) && stringInSlice(s[1],startingTokenTickers))
			} else {
				fmt.Println(stringInSlice(s[0],startingTokenTickers))
			}
		}
		
		// Filter out pools for which available true
		// ret raw = matrix(get prices of constituent tokens of filtered pools)
		// +hist return + other return - swap costs
		// figure out how to sum pools to 1	- 2 assets
		// hist return is not just prices here - it is also other stuff

		// Define optimization function
		fcn := func(x_weights []float64) float64 {
			print := false
			
			number_of_tokens := 4 // cols (=intersection of pools + portfolio)
			number_of_days := 4 // rows (=lowest number of days in historical data)

			ret_mat := mat.NewDense(number_of_days, number_of_tokens, []float64{
				11.1, 14.2, 31, 111.4,
				11.2, 16, 38.3, 111.3,
				11.3, 11, 37.2, 111.2,
				11.15, 21.2, 33.3, 111.1,
			})

			//fmt.Println("Checkpoint 1")

			ret_mat_pct := mat.NewDense(number_of_days - 1, number_of_tokens, nil)

				for ii := 0; ii < number_of_days - 1; ii++ { // row
					for jj := 0; jj < number_of_tokens; jj++ { // col
						ret_mat_pct.Set(ii, jj, ret_mat.At(ii+1, jj)/ret_mat.At(ii, jj)-1.0)
				} // ii
			} // jj

			//fmt.Println("Checkpoint 2")

			var avg_returns []float64 

			for jj := 0; jj < number_of_tokens; jj++ {
				total := 0.0
				for ii := 0; ii < number_of_days - 1; ii++ { 
					total += ret_mat_pct.At(ii,jj)
				}
			avg_returns = append(avg_returns,252*total/float64((number_of_days - 1)))
		}

			//fmt.Println("Checkpoint 3")
		
		if print {
				fmt.Print("RET %:")

				fmt.Println(ret_mat_pct)
			
				fmt.Println("RETURNS: ")
				fmt.Print(avg_returns[0])
				fmt.Print(" | ")
				fmt.Print(avg_returns[1])
				fmt.Print(" | ")
				fmt.Print(avg_returns[2])
				fmt.Print(" | ")
				fmt.Println(avg_returns[3])
		}
			
			
			ret := mat.NewVecDense(4, avg_returns) // vector of returns
			if print {
				fmt.Print("RET: ")
				fmt.Println(ret)
			}
			

			if print {
				fmt.Print("x_weights before normalisation: ")
				fmt.Println(x_weights)
			}
			for j := 0; j < len(x_weights); j++ {
				if x_weights[j] < 0 {
					x_weights[j] = 0
				}
			}

			totl := sum(x_weights)
			for j := 0; j < len(x_weights); j++ {
				x_weights[j] = x_weights[j] / totl
			}
			if print {
				fmt.Print("x_weights after normalisation: ")
				fmt.Println(x_weights)
			}
			weights := mat.NewVecDense(number_of_tokens, x_weights) // vector of portfolio weights
			//weights := mat.NewVecDense(number_of_tokens, []float64{0.25, 0.25, 0.25, 0.25}) // vector of portfolio weights

			var cov *mat.SymDense = mat.NewSymDense(number_of_tokens, nil)
			cov.Reset()

			stat.CovarianceMatrix(cov, ret_mat_pct, nil) 
			if print {
				fmt.Print("COV WITH T: ")
				fmt.Println(cov)
			}

			var cov2 *mat.SymDense = mat.NewSymDense(number_of_tokens, nil)
			//cov2.Reset()
						
			for ii := 0; ii < number_of_tokens; ii++ { // row
					for jj := 0; jj < number_of_tokens; jj++ { // col
						cov2.SetSym(ii, jj, cov.At(ii, jj)*252) // annualise them
				} // ii
			} // jj
			
			// fmt.Println("Checkpoint 3.5")
			
			blended_return := mat.Dot(ret, weights)
			if print {
				fmt.Println("BLENDED RETURN: ")
				fmt.Println(blended_return)
			}
			
			//fmt.Println("Checkpoint 4")
			
			risk_step0 := mat.NewVecDense(4, nil)
			risk_step0.MulVec(cov2, weights)
			if print {
				fmt.Print("Portfolio var step 0: ")
				fmt.Print(risk_step0)
				fmt.Print(" | ")
			}
			risk := math.Sqrt(mat.Dot(weights, risk_step0))
			if print {
				fmt.Print("PORTVOLIO VOLATILITY: ")
				fmt.Println(risk)
			}

			return -blended_return / risk
		}

		var p0 = []float64{0.25, 0.25, 0.25, 0.25} 
		fmt.Println("TESTING WITH INITIAL 25% WEIGHTS: ")
		fmt.Print(fcn(p0))

		p := optimize.Problem{
			Func: fcn,
		}

		result, err := optimize.Minimize(p, p0, nil, nil)
		if err != nil {
			log.Fatal(err)
		}
		if err = result.Status.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Print("RAW WEIGHTS OPTIMIZED:")
		fmt.Println(result)

		fmt.Print("FINAL WEIGHTS OPTIMIZED:")
		result_norm := result
		for j := 0; j < len(result_norm.Location.X); j++ {
			if result_norm.Location.X[j] < 0 {
				result_norm.Location.X[j] = 0
			}
		}

		totl := sum(result_norm.Location.X)
		for j := 0; j < len(result_norm.Location.X); j++ {
			result_norm.Location.X[j] = result.Location.X[j] / totl
		}
		fmt.Print(result_norm.Location.X[0])
		fmt.Print(" | ")
		fmt.Print(result_norm.Location.X[1])
		fmt.Print(" | ")
		fmt.Print(result_norm.Location.X[2])
		fmt.Print(" | ")
		fmt.Println(result_norm.Location.X[3])
		
		fmt.Println("OPTIMIZATION COMPLETE")

	return NewOptimisedPortfolio(database)
}