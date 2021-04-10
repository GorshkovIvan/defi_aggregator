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

func get_latest_token_price(token string) float64{
	if token == "USD" {return 1.0}
	if isHistDataAlreadyDownloadedDatabase(token) {
		return returnPricesInCollection(token)[0]
	}
	return 1.00
}

func recalculate_balances_from_weights(weights_optimised []float64, total float64, pool_tkn0s []string, pool_tkn1s []string, pool_ratios []float64,own_pf []OwnPortfolioRecord) []float64 {

	var token_balances_resulting_from_entered_weights []float64 // what balances of pf translates into - len(own_pf)

	fmt.Print("IN RECALC BALANCES: ")
	fmt.Println(len(weights_optimised))
	fmt.Println(len(pool_tkn0s)) 
	fmt.Println(len(pool_tkn1s))
	fmt.Println(len(pool_ratios))
	fmt.Println(len(own_pf))

	// Resize array of token balances - data structure mirrors own_pf
	for i:=0; i < len(own_pf);i++ {
		token_balances_resulting_from_entered_weights = append(token_balances_resulting_from_entered_weights,0.0)
	}

	for j := 0; j < len(weights_optimised); j++ {
		fmt.Print("j: ")
		fmt.Print(j)
	
		idx0 := 0 // find idx of pool_tkn0s[j] in own pf
		idx1 := 0 // find idx of pool_tkn1s[j] in own pf
		
		for jj := 0; jj < len(own_pf);jj++ {
			if pool_tkn0s[j] == own_pf[jj].Token {
				idx0 = jj
				break
			}
		} // find idx0
		
		for jj := 0; jj < len(own_pf);jj++ {
			if pool_tkn1s[j] == own_pf[jj].Token {
				idx1 = jj
				break
			}
		} // find idx1		
	
		token_balances_resulting_from_entered_weights[idx0] += weights_optimised[j] * total / (1 + pool_ratios[j]) 
		token_balances_resulting_from_entered_weights[idx1] += weights_optimised[j] * pool_ratios[j] * total / (1 + pool_ratios[j])	// which one to mult by total?
	
		fmt.Print(" | token balances resulting: ")
		fmt.Println(token_balances_resulting_from_entered_weights[idx0])
		fmt.Println(token_balances_resulting_from_entered_weights[idx1])
	} // translate pool weights to total token balances

	return token_balances_resulting_from_entered_weights
}


func nrm_pool_wgts(pool_weights_raw []float64, pool_tkn0s []string, pool_tkn1s []string, pool_ratios []float64, own_pf []OwnPortfolioRecord) []float64 {
	var weights_optimised []float64
	fmt.Print("..IN nrm pool wgts..checkpoint 991..")
	if len(pool_tkn0s) == 0 {return weights_optimised}
	
	// eliminate negatives
	for j := 0; j < len(pool_weights_raw); j++ {
		if pool_weights_raw[j] < 0 {
				weights_optimised = append(weights_optimised,0.0)
		} else if pool_weights_raw[j] > 1 {
				weights_optimised = append(weights_optimised,1.00)
		} else {
				weights_optimised = append(weights_optimised,pool_weights_raw[j])
		}
	}
		fmt.Print("..IN nrm pool wgts..checkpoint 992..")
	total := 0.0 // sum current prices x balances of deployable tokens
	
	for i:= 0; i < len(own_pf);i++ {
		total += get_latest_token_price(own_pf[i].Token) * float64(own_pf[i].Amount)
	}
	fmt.Print("..IN nrm pool wgts..checkpoint 993..")
	// now translate pool weights to total token balances
	token_balances_resulting_from_entered_weights := recalculate_balances_from_weights(weights_optimised,total,pool_tkn0s,pool_tkn1s,pool_ratios,own_pf)

	violation_count := 0	
	for i:=0; i < len(own_pf); i++ {
		if token_balances_resulting_from_entered_weights[i] > float64(own_pf[i].Amount) * get_latest_token_price(own_pf[i].Token) {
			violation_count++
		}
	}
		fmt.Print("..IN nrm pool wgts..checkpoint 994..")
	for violation_count > 0 { // make sure they sum to individual token balances
				for j := 0; j < len(weights_optimised); j++ {
						
				idx0 := 0 // find idx of pool_tkn0s[j] in own pf
				idx1 := 0 // find idx of pool_tkn1s[j] in own pf
				
				for jj := 0; jj < len(own_pf);jj++ {
					if pool_tkn0s[j] == own_pf[jj].Token {
						idx0 = jj
						break
					}
				} // find idx0
				
				for jj := 0; jj < len(own_pf);jj++ {
					if pool_tkn1s[j] == own_pf[jj].Token {
						idx1 = jj
						break
					}
				} // find idx1
				
				amt0 := float64(own_pf[idx0].Amount) * get_latest_token_price(own_pf[idx0].Token) // USD terms
				amt1 := float64(own_pf[idx1].Amount) * get_latest_token_price(own_pf[idx1].Token) // USD terms
				
				rat0 := amt0 / token_balances_resulting_from_entered_weights[idx0]
				rat1 := amt1 / token_balances_resulting_from_entered_weights[idx1]
				rat := math.Min(rat0,rat1)

				if rat < 1 { 				// scale pool % by rat
					weights_optimised[j] = weights_optimised[j] * rat	
				} // if rat
			} // for len pool weights raw
	fmt.Print("..IN nrm pool wgts..checkpoint 995..")
		token_balances_resulting_from_entered_weights = recalculate_balances_from_weights(weights_optimised,total,pool_tkn0s,pool_tkn1s,pool_ratios,own_pf)		
		violation_count = 0
		
		for i:=0; i < len(own_pf); i++ {
			if token_balances_resulting_from_entered_weights[i] > float64(own_pf[i].Amount) * get_latest_token_price(own_pf[i].Token) {
				violation_count++
			}
		}
		
	} // violation count loop ends
		fmt.Print("..IN nrm pool wgts..checkpoint 996..")
	return weights_optimised
}

func OptimisePortfolio(database *Database) []OptimisedPortfolioRecord {
	fmt.Println("----ENTERING PORTFOLIO OPTIMISATION----------")

	var startingTokenTickers []string 

	if database.Risksetting == 0 {
		fmt.Println("WARNING: Risk setting set to zero!")
	}
	
	// For testing
	if len(database.ownstartingportfolio) == 0 {
		database.ownstartingportfolio = append(database.ownstartingportfolio,OwnPortfolioRecord{"WETH",42.0})
		database.ownstartingportfolio = append(database.ownstartingportfolio,OwnPortfolioRecord{"DAI",1269.0})
		database.ownstartingportfolio = append(database.ownstartingportfolio,OwnPortfolioRecord{"USD",420.69})
	}
	
	/// DUPLICATE - CREATE A MORE EFFICIENT IMPLEMENTATION
	// THIS CALCULATES NUMBER OF ELEMENTS IN WEIGHTS - AND FEEDS THIS NUMBER TO P0
		var h_array []HistoricalCurrencyData // are these the same size? 

		// 0 - Clean starting portfolio for duplicates
		for i := 0; i < len(database.ownstartingportfolio); i++ {
			if !stringInSlice(database.ownstartingportfolio[i].Token, startingTokenTickers) {
				startingTokenTickers = append(startingTokenTickers, database.ownstartingportfolio[i].Token)
			}
		}

		// 1 - Calculate available pools for deployment + PULL THEIR DATA
		for i := 0; i < len(database.currencyinputdata); i++ {					
			s := strings.Split(database.currencyinputdata[i].Pair, "/")
			var Available bool			
			if len(s) == 2 {
				Available = stringInSlice(s[0],startingTokenTickers) && stringInSlice(s[1],startingTokenTickers)
			} else {
				Available = stringInSlice(s[0],startingTokenTickers)
			}
			
			if Available {
				var h HistoricalCurrencyData
				if len(s) == 2 {
					h = retrieveDataForTokensFromDatabase2(s[0], s[1])
					h_array = append(h_array,h)
				} else {
					h = retrieveDataForTokensFromDatabase2(s[0], "USD")
					h_array = append(h_array,h)
				}
			} // if Available
		} // for pools
	
		
	// Define optimization function
	fcn := func(x_weights []float64) float64 {
	
		var h_array []HistoricalCurrencyData
		var deployable_portfolio_array []OwnPortfolioRecord
		var pool_tkn0s []string
		var pool_tkn1s []string 
		var pool_ratios []float64 
		
		fmt.Print("..Checkpoint 000..")

		// 0 - Clean starting portfolio for duplicates
		for i := 0; i < len(database.ownstartingportfolio); i++ {
			if !stringInSlice(database.ownstartingportfolio[i].Token, startingTokenTickers) {
				startingTokenTickers = append(startingTokenTickers, database.ownstartingportfolio[i].Token)
			}
		}

		fmt.Print("..Checkpoint 001..")
		for i := 0; i < len(startingTokenTickers); i++ {
			fmt.Println(startingTokenTickers[i])
		}
		
		// 1 - Calculate available pools for deployment + PULL THEIR DATA
		for i := 0; i < len(database.currencyinputdata); i++ {					
			// for each pair - if pair is made up of portfolio tokens - add to available pool list
			s := strings.Split(database.currencyinputdata[i].Pair, "/")
			// Filter out pools for which available true						
			var Available bool			
			if len(s) == 2 {
				Available = stringInSlice(s[0],startingTokenTickers) && stringInSlice(s[1],startingTokenTickers)
			} else {
				Available = stringInSlice(s[0],startingTokenTickers)
			}
			
			if Available {
				var h HistoricalCurrencyData
				if len(s) == 2 {
					h = retrieveDataForTokensFromDatabase2(s[0], s[1])
					h_array = append(h_array,h)
				} else {
					h = retrieveDataForTokensFromDatabase2(s[0], "USD")
					h_array = append(h_array,h)
				}
			} // if Available
			
		} // for pools

		fmt.Print("..Checkpoint 002..")
		
		// 2 - Convert data to return matrix		
		ret_mat_xxx := mat.NewDense(1, 1, nil)
		if len(h_array) > 0 {
			ret_mat_xxx = mat.NewDense(len(h_array[0].Price), len(h_array), nil)
			for ii:= 0; ii < len(h_array[0].Price); ii++ { // row?
				for jj:= 0; jj < len(h_array); jj++ {		
					ret_mat_xxx.Set(ii,jj,float64(h_array[jj].Price[ii]))
				}
			}			
		} else {
			ret_mat_xxx.Zero()
		} // if len(h_array) > 0
				
		number_of_tokens := int(math.Max(float64(len(h_array)),1))
		number_of_days := 2
		if len(h_array) > 0 {
			number_of_days = int(math.Max(float64(len(h_array[0].Price)),2))
		}
		
		fmt.Print("len len(h_array): ")
		fmt.Println(len(h_array))
		fmt.Println(number_of_days)
		fmt.Println(number_of_tokens)

		ret_mat := ret_mat_xxx
		fmt.Print("RET MAT: ")
		fmt.Println(ret_mat)
		ret_mat_pct := mat.NewDense(number_of_days - 1, number_of_tokens, nil)
		fmt.Print(" xx ")		
		
		for ii := 0; ii < number_of_days - 1; ii++ { // row
					for jj := 0; jj < number_of_tokens; jj++ { // col
						if ret_mat.At(ii, jj) != 0.0 {
							ret_mat_pct.Set(ii, jj, ret_mat.At(ii+1, jj)/ret_mat.At(ii, jj)-1.0)						
						} else {
							ret_mat_pct.Set(ii, jj, 0.0)
						}
				} // jj
		} // ii

		fmt.Print("..Checkpoint 002.5..")

		var avg_returns []float64 
		for jj := 0; jj < number_of_tokens; jj++ {
				total := 0.0
				for ii := 0; ii < number_of_days - 1; ii++ { 
					total += ret_mat_pct.At(ii,jj)
				}
				avg_returns = append(avg_returns,252*total/float64((number_of_days - 1)))
			}

		ret := mat.NewVecDense(number_of_tokens, avg_returns) // vector of returns			

		fmt.Print("..Checkpoint 003..")

		// 3 & 4 - PREPARE VARIABLES - TO FEED INTO FUNC TO NORMALIZE WEIGHTS
		// 3 - Populate list of pool token0 and pool token1, pool ratios
		for j:=0; j < len(h_array); j++ {
			s := strings.Split(h_array[j].Ticker, "/")
			pool_tkn0s = append(pool_tkn1s,s[0])
				if len(s) == 1 {
					pool_tkn1s = append(pool_tkn1s,"USD")
				} else {
					pool_tkn1s = append(pool_tkn1s,s[1])
				}	
			pool_ratios = append(pool_ratios,1.25) // NEED TO CHANGE THIS
		}		

		fmt.Print("..Checkpoint 004..")
		
		// 4 - Filter out the portfolio items which have pools to be deployed into
		for i:=0; i < len(database.ownstartingportfolio);i++ {
			// if tokens are in our pool_tkn0s or pool_tkn1s - then filter it out
			if stringInSlice(database.ownstartingportfolio[i].Token,pool_tkn0s) || stringInSlice(database.ownstartingportfolio[i].Token,pool_tkn1s) {
				deployable_portfolio_array = append(deployable_portfolio_array,database.ownstartingportfolio[i])
			}
		}

		fmt.Print("..Checkpoint 005..")

		// 5 - Normalise weights
		x_weights = nrm_pool_wgts(x_weights, pool_tkn0s, pool_tkn1s, pool_ratios, deployable_portfolio_array)		
		weights := mat.NewVecDense(number_of_tokens, x_weights)
		
		fmt.Print("..Checkpoint 006..")
		
		// 6 - Calculate covariance matrix
		var cov *mat.SymDense = mat.NewSymDense(number_of_tokens, nil)
		cov.Reset()
		stat.CovarianceMatrix(cov, ret_mat_pct, nil) 
		var cov2 *mat.SymDense = mat.NewSymDense(number_of_tokens, nil)				
		for ii := 0; ii < number_of_tokens; ii++ { // row
					for jj := 0; jj < number_of_tokens; jj++ { // col
						cov2.SetSym(ii, jj, cov.At(ii, jj)*252) // annualise them
				} // ii
		} // jj

		fmt.Print("..Checkpoint 007..")
		
		// 7 - Calculate blended return and risk
		blended_return := mat.Dot(ret, weights)			
		risk_step0 := mat.NewVecDense(number_of_tokens, nil)
		risk_step0.MulVec(cov2, weights)
		risk := math.Sqrt(mat.Dot(weights, risk_step0))

		fmt.Print("..Checkpoint 008..")

		// 8 - Return sharpe ratio
		sharpe := -blended_return / risk
		
		if math.IsNaN(sharpe) {
				return 0.0
		}
		
		if math.IsInf(sharpe,0) {
				return 0.0
		}
		
			return sharpe
		} // fcn definition complete
		
		
		// var p0 = []float64{0.25, 0.25, 0.25, 0.25} // should be sized same as number of assets
		
		var p0 []float64
		
		if len(h_array) > 0 {
		for i := 0; i < len(h_array);i++{
			p0 = append(p0,1/float64(len(h_array)))
			}
		} else {
			p0 = append(p0,0.0)
		}	// 1/number_of_tokens

		p := optimize.Problem{
			Func: fcn,
		}

		fmt.Print("ABOUT TO CALL MINIMIZE")
		result, err := optimize.Minimize(p, p0, nil, nil)
		if err != nil {
			log.Fatal(err)
		}
		if err = result.Status.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Print("RAW WEIGHTS OPTIMIZED:")
		fmt.Println(result)

		// fmt.Print("FINAL WEIGHTS OPTIMIZED:")
		// result_norm := nrm_pool_wgts(result) - ADD ALL THE OTHER INPUTS
		// fmt.Print(result_norm)
				
		fmt.Println("OPTIMIZATION COMPLETE..RETURNING!!!..")

	return NewOptimisedPortfolio(database)
}