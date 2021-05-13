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

func get_latest_token_price(token string) float64 {
	if token == "USD" {
		return 1.0
	}
	if isHistDataAlreadyDownloadedDatabase(token) {
		return returnPricesInCollection(token)[0]
	}
	return 1.00
}

func recalculate_balances_from_weights(weights_optimised []float64, total float64, pool_tkn0s []string, pool_tkn1s []string, pool_ratios []float64, own_pf []OwnPortfolioRecord) []float64 {

	var token_balances_resulting_from_entered_weights []float64 // what balances of pf translates into - len(own_pf)
	/*
		fmt.Print("IN RECALC BALANCES: ")
		fmt.Println(len(weights_optimised))
		fmt.Println(len(pool_tkn0s))
		fmt.Println(len(pool_tkn1s))
		fmt.Println(len(pool_ratios))
		fmt.Println(len(own_pf))
	*/
	// Resize array of token balances - data structure mirrors own_pf
	for i := 0; i < len(own_pf); i++ {
		token_balances_resulting_from_entered_weights = append(token_balances_resulting_from_entered_weights, 0.0)
	}

	for j := 0; j < len(weights_optimised); j++ {
		fmt.Print("j: ")
		fmt.Print(j)

		idx0 := 0 // find idx of pool_tkn0s[j] in own pf
		idx1 := 0 // find idx of pool_tkn1s[j] in own pf

		for jj := 0; jj < len(own_pf); jj++ {
			if pool_tkn0s[j] == own_pf[jj].Token {
				idx0 = jj
				break
			}
		} // find idx0

		for jj := 0; jj < len(own_pf); jj++ {
			if pool_tkn1s[j] == own_pf[jj].Token {
				idx1 = jj
				break
			}
		} // find idx1

		token_balances_resulting_from_entered_weights[idx0] += weights_optimised[j] * total / (1 + pool_ratios[j])
		token_balances_resulting_from_entered_weights[idx1] += weights_optimised[j] * pool_ratios[j] * total / (1 + pool_ratios[j]) // which one to mult by total?

		//fmt.Print(" | token balances resulting: ")
		//fmt.Println(token_balances_resulting_from_entered_weights[idx0])
		//fmt.Println(token_balances_resulting_from_entered_weights[idx1])
	} // translate pool weights to total token balances

	return token_balances_resulting_from_entered_weights
}

func nrm_pool_wgts(pool_weights_raw []float64, pool_tkn0s []string, pool_tkn1s []string, pool_ratios []float64, own_pf []OwnPortfolioRecord, own_pf_px []float64) ([]float64,[]string,[]float64) {
	var weights_optimised []float64

	var leftovertokens []string
	var leftoveramounts []float64

//	fmt.Print("..Normalizing optimized raw weights: ")
//	fmt.Print("starting point: ")
//	fmt.Println(len(pool_weights_raw))
//	fmt.Println(len(pool_tkn0s))
//	fmt.Println(len(pool_tkn1s))
//	fmt.Println(len(pool_ratios))
//	fmt.Println(len(own_pf_px))
	if len(pool_weights_raw) == len(pool_tkn0s) && len(pool_weights_raw) == len(own_pf_px) {
	for i:=0; i < len(pool_weights_raw); i++ {
	//	fmt.Print("i: ")
	//	fmt.Print(i)
	//	fmt.Print(" | ")
	//	fmt.Print(pool_weights_raw[i])
	//	fmt.Print(" | ")
	//	fmt.Print(pool_tkn0s[i])
	//	fmt.Print(" | ")
	//	fmt.Println(pool_tkn1s[i])
	}
}

//fmt.Print("Checkpoint 992")

if len(pool_tkn0s) == 0 {

		for jjj := 0; jjj < len(own_pf);jjj++ {
			weights_optimised = append(weights_optimised,0.0)		
		}

		total := 0.0
		for i := 0; i < len(own_pf); i++ {
			total += own_pf_px[i] * float64(own_pf[i].Amount)
		}
	
		token_balances_resulting_from_entered_weights := recalculate_balances_from_weights(weights_optimised, total, pool_tkn0s, pool_tkn1s, pool_ratios, own_pf)

		//fmt.Print("Checkpoint 993")
	//	fmt.Print("len: ")
	//	fmt.Print(len(own_pf))
		for ii := 0; ii < len(own_pf); ii++ {
		/*
			fmt.Print("Checkpoint 994")
			fmt.Print("ii: ")
			fmt.Print(ii)
			fmt.Print(" | tkn: ")
			fmt.Print(own_pf[ii].Token)
			fmt.Print(" | ")
			fmt.Print(own_pf[ii].Amount)
			fmt.Print(" | tkn bal res: ")
			fmt.Print(token_balances_resulting_from_entered_weights[ii])
		*/
			diff := float64(own_pf[ii].Amount)*own_pf_px[ii] - token_balances_resulting_from_entered_weights[ii]
			fmt.Print(" | diff: ")
			fmt.Println(diff)

			//if diff > 0.0000000000001 {
				leftovertokens = append(leftovertokens,own_pf[ii].Token)
				leftoveramounts = append(leftoveramounts,diff)
			//}
		}

	//	fmt.Print("Number of leftovers: ")
	//	fmt.Print(len(leftovertokens))
		return weights_optimised, leftovertokens,leftoveramounts
	}

	// eliminate negatives
	for j := 0; j < len(pool_weights_raw); j++ {
		if pool_weights_raw[j] < 0 {
			weights_optimised = append(weights_optimised, 0.0)
		} else if pool_weights_raw[j] > 1 {
			weights_optimised = append(weights_optimised, 1.00)
		} else {
			weights_optimised = append(weights_optimised, pool_weights_raw[j])
		}
	}
	//fmt.Print("..IN nrm pool wgts..checkpoint 992..")
	total := 0.0 // sum current prices x balances of deployable tokens

	for i := 0; i < len(own_pf); i++ {
		total += own_pf_px[i] * float64(own_pf[i].Amount)
	}

	tot_wgt := 0.0

	for j := 0; j < len(weights_optimised); j++ {
		tot_wgt += weights_optimised[j] 
	}
	// re-normalize so weights sum up to 1
	for j := 0; j < len(weights_optimised); j++ {
		weights_optimised[j] = weights_optimised[j] / tot_wgt
		fmt.Print(j)
		fmt.Print(" | ")
		fmt.Println(weights_optimised[j])
	}
	
	tot_test := 0.0
	//t_test := 0.0

	for j := 0; j < len(weights_optimised); j++ {
		tot_test += weights_optimised[j] 
	}
	fmt.Print("new TOTAL: ")
	fmt.Print(tot_test)


	//fmt.Print("..IN nrm pool wgts..checkpoint 993..")
	// now translate pool weights to total token balances
	token_balances_resulting_from_entered_weights := recalculate_balances_from_weights(weights_optimised, total, pool_tkn0s, pool_tkn1s, pool_ratios, own_pf)

	violation_count := 0
	for i := 0; i < len(own_pf); i++ { // get_latest_token_price(own_pf[i].Token)
		if token_balances_resulting_from_entered_weights[i] > float64(own_pf[i].Amount)*own_pf_px[i] {
			violation_count++
		}
	}
	// fmt.Print("..994..")
	for violation_count > 0 { // make sure they sum to individual token balances
		//fmt.Print("..violations BEFORE change: ")
		//fmt.Print(violation_count)

		for j := 0; j < len(weights_optimised); j++ {

			idx0 := 0 // find idx of pool_tkn0s[j] in own pf
			idx1 := 0 // find idx of pool_tkn1s[j] in own pf

			for jj := 0; jj < len(own_pf); jj++ {
				if pool_tkn0s[j] == own_pf[jj].Token {
					idx0 = jj
					break
				}
			} // find idx0

			for jj := 0; jj < len(own_pf); jj++ {
				if pool_tkn1s[j] == own_pf[jj].Token {
					idx1 = jj
					break
				}
			} // find idx1

			amt0 := float64(own_pf[idx0].Amount) * own_pf_px[idx0] // get_latest_token_price(own_pf[idx0].Token) // USD terms
			amt1 := float64(own_pf[idx1].Amount) * own_pf_px[idx1] // get_latest_token_price(own_pf[idx1].Token) // USD terms

			rat0 := 0.0
			rat1 := 0.0

			if token_balances_resulting_from_entered_weights[idx0] > 0.01 {
				rat0 = amt0 / token_balances_resulting_from_entered_weights[idx0]				
			}
			if token_balances_resulting_from_entered_weights[idx1] > 0.01 {
				rat1 = amt1 / token_balances_resulting_from_entered_weights[idx1]
			}

			rat := math.Min(rat0, rat1)

		//	fmt.Print("ratio in optimiser: ")	
		//	fmt.Print(rat)
			if rat < 1 && rat > 0 { // scale pool % by rat
				weights_optimised[j] = weights_optimised[j] * rat
			} // if rat
		} // for len pool weights raw
		// fmt.Print("..995..")
		token_balances_resulting_from_entered_weights = recalculate_balances_from_weights(weights_optimised, total, pool_tkn0s, pool_tkn1s, pool_ratios, own_pf)
		violation_count = 0

		for i := 0; i < len(own_pf); i++ { // get_latest_token_price(own_pf[i].Token)
			if token_balances_resulting_from_entered_weights[i] > float64(own_pf[i].Amount)*own_pf_px[i] {
				violation_count++
			}
		}
		

		// get leftovers
		// Update return matrix with actual token returns
		for ii := 0; ii < len(own_pf); ii++ {
			fmt.Print("ii: ")
			fmt.Print(ii)
			fmt.Print(" | tkn: ")
			fmt.Print(own_pf[ii].Token)
			fmt.Print(" | ")
			fmt.Print(own_pf[ii].Amount)
			fmt.Print(" | tkn bal res: ")
			fmt.Print(token_balances_resulting_from_entered_weights[ii])

			diff := float64(own_pf[ii].Amount)*own_pf_px[ii] - token_balances_resulting_from_entered_weights[ii]
			fmt.Print(" | diff: ")
			fmt.Println(diff)

			//if diff > 0.0000000000001 {
				leftovertokens = append(leftovertokens,own_pf[ii].Token)
				leftoveramounts = append(leftoveramounts,diff)
			//}
		}
		//fmt.Print("..violations AFTER change: ")
		//fmt.Print(violation_count)

	} // violation count loop ends
	fmt.Print("..returning optimised, normalised weights..")
	return weights_optimised, leftovertokens,leftoveramounts
}

func OptimisePortfolio(database *Database) []OptimisedPortfolioRecord {
	fmt.Println("----ENTERING PORTFOLIO OPTIMISATION----------")

	if database.Risksetting == 0 {
		fmt.Println("WARNING: Risk setting set to zero!")
	}

	var startingTokenTickers []string    // Unique own portfolio token list
	var h_array []HistoricalCurrencyData // For storing historical px data of pools to deploy into
	var deployable_portfolio_array []OwnPortfolioRecord
	var pool_name_array []string
	var pool_tkn0s []string   // Pool list token0
	var pool_tkn1s []string   // Pool list token1
	var pool_ratios []float64 // Need to pull this - XXX
	var own_pf_px []float64
	var avg_returns []float64
	var Available bool
	var optimised_pf []OptimisedPortfolioRecord // return value

	// Add some tokens to own portfolio - for testing - remove in final version
	if len(database.ownstartingportfolio) == 0 {
		database.ownstartingportfolio = append(database.ownstartingportfolio, OwnPortfolioRecord{"WETH", 0.01})
		database.ownstartingportfolio = append(database.ownstartingportfolio, OwnPortfolioRecord{"DAI", 200})
		database.ownstartingportfolio = append(database.ownstartingportfolio, OwnPortfolioRecord{"USDC", 200})
	}

	// 0 - Clean starting portfolio for duplicates
	for i := 0; i < len(database.ownstartingportfolio); i++ {
		if !stringInSlice(database.ownstartingportfolio[i].Token, startingTokenTickers) {
			startingTokenTickers = append(startingTokenTickers, database.ownstartingportfolio[i].Token)
		}
	}

	// 1 - Calculate available pools for deployment + PULL THEIR DATA
	for i := 0; i < len(database.currencyinputdata); i++ {
		s := strings.Split(database.currencyinputdata[i].Pair, "/")

		// Loop through pools - if pairs in pool made up of portfolio tokens - add to AVAILABLE pools
		if len(s) == 2 {
			Available = stringInSlice(s[0], startingTokenTickers) && stringInSlice(s[1], startingTokenTickers)
		} else {
			Available = stringInSlice(s[0], startingTokenTickers)
		}

		if Available {
			pool_name_array = append(pool_name_array, database.currencyinputdata[i].Pool)
			var h HistoricalCurrencyData
			if len(s) == 2 {
				h = retrieveDataForTokensFromDatabase2(s[0], s[1])
				h_array = append(h_array, h)
			} else {
				h = retrieveDataForTokensFromDatabase2(s[0], "USD")
				h_array = append(h_array, h)
			}
		} // if Available

	} // Calculate available pools for deployment + PULL THEIR DATA - end

	// 2 -  pulled pool price data to return matrix (prices first)
	ret_mat_xxx := mat.NewDense(1, 1, nil)
	if len(h_array) > 0 {
		ret_mat_xxx = mat.NewDense(len(h_array[0].Price), len(h_array), nil)
		for ii := 0; ii < len(h_array[0].Price); ii++ { // row?
			for jj := 0; jj < len(h_array); jj++ {
				ret_mat_xxx.Set(ii, jj, float64(h_array[jj].Price[ii]))
			}
		}
	} else {
		ret_mat_xxx.Zero()
	} // if len(h_array) > 0

	// 3 & 4 - PREPARE VARIABLES - TO FEED INTO FUNC TO NORMALIZE WEIGHTS
	// 3 - Populate list of pool token0 and pool token1, pool ratios
	for j := 0; j < len(h_array); j++ {
		s := strings.Split(h_array[j].Ticker, "/")
		pool_tkn0s = append(pool_tkn1s, s[0])
		if len(s) == 1 {
			pool_tkn1s = append(pool_tkn1s, "USD")
		} else {
			pool_tkn1s = append(pool_tkn1s, s[1])
		}
		pool_ratios = append(pool_ratios, 1.25) // NEED TO CHANGE THIS
	}

fmt.Print("Before filter tokn0s: ")
for i:=0; i < len(pool_tkn0s); i++ {
	fmt.Print(i)
	fmt.Print(": ")
	fmt.Println(pool_tkn0s)
}

fmt.Print("Before filter pool_tkn1s: ")
for i:=0; i < len(pool_tkn1s); i++ {
	fmt.Print(i)
	fmt.Print(": ")
	fmt.Println(pool_tkn1s)
}

	// 4 - Filter out OWN portfolio tokens which have pools to be deployed into
	for i := 0; i < len(database.ownstartingportfolio); i++ {
		fmt.Print(i)
		fmt.Print(": ")
		fmt.Print(database.ownstartingportfolio[i].Token)
		// if tokens are in our pool_tkn0s or pool_tkn1s - then filter it out
		if stringInSlice(database.ownstartingportfolio[i].Token, pool_tkn0s) || stringInSlice(database.ownstartingportfolio[i].Token, pool_tkn1s) {
			deployable_portfolio_array = append(deployable_portfolio_array, database.ownstartingportfolio[i])
			own_pf_px = append(own_pf_px, get_latest_token_price(database.ownstartingportfolio[i].Token))
			fmt.Print(" | deployable portfolio i : ")
			fmt.Print(i)
			fmt.Print(" | ")
			fmt.Print(deployable_portfolio_array[int(len(deployable_portfolio_array)-1)])
			fmt.Print(" | ")
			fmt.Print(" | px: ")
			fmt.Print(own_pf_px[int(len(deployable_portfolio_array)-1)])
		}
	}

	// 5 - Define dimensions
	number_of_tokens := int(math.Max(float64(len(h_array)), 1)) // Number of pools to deploy into
	number_of_days := 2                                         // starting value to prevent errors in sizing
	if len(h_array) > 0 {
		number_of_days = int(math.Max(float64(len(h_array[0].Price)), 2))
	}

	// 6 - Declare matrix for returns data in %
	ret_mat_pct := mat.NewDense(number_of_days-1, number_of_tokens, nil)

	// 7 - Populate this matrix with returns in %
	for ii := 0; ii < number_of_days-1; ii++ { // row
		for jj := 0; jj < number_of_tokens; jj++ { // col
			if ret_mat_xxx.At(ii, jj) != 0.0 {
				ret_mat_pct.Set(ii, jj, ret_mat_xxx.At(ii+1, jj)/ret_mat_xxx.At(ii, jj)-1.0)
			} else {
				ret_mat_pct.Set(ii, jj, 0.0)
			}
		} // jj
	} // ii

	// 8 - Calculate average returns by token to be deployed into
	for jj := 0; jj < number_of_tokens; jj++ {
		total := 0.0
		for ii := 0; ii < number_of_days-1; ii++ {
			total += ret_mat_pct.At(ii, jj)
		}
		avg_returns = append(avg_returns, 252*total/float64((number_of_days-1)))
	}
	
	
	
		
	ret := mat.NewVecDense(number_of_tokens, avg_returns) // vector of returns

	// 9 - Calculate covariance matrix
	var cov *mat.SymDense = mat.NewSymDense(number_of_tokens, nil)
	cov.Reset()
	stat.CovarianceMatrix(cov, ret_mat_pct, nil)
	var cov2 *mat.SymDense = mat.NewSymDense(number_of_tokens, nil)


	for ii := 0; ii < number_of_tokens; ii++ { // row
		for jj := 0; jj < number_of_tokens; jj++ { // col
			cov2.SetSym(ii, jj, cov.At(ii, jj)*252) // annualise them
		} // ii
	} // jj

	// 10 - Define optimization function
	fcn := func(x_weights []float64) float64 {
		//var leftovertokens []string
		//var leftoveramounts []float64
		// Normalise weights
	//	fmt.Print("calculating x weights...")
	//	fmt.Print("pool_tkn0s len: ")
	//	fmt.Print(len(pool_tkn0s))

		x_weights, _, _ = nrm_pool_wgts(x_weights, pool_tkn0s, pool_tkn1s, pool_ratios, deployable_portfolio_array, own_pf_px)
	//	fmt.Print("calculated x weights...")
		weights := mat.NewVecDense(number_of_tokens, x_weights)
		// Calculate covariance matrix is outside
		// Calculate blended return and risk
		blended_return := mat.Dot(ret, weights)
		risk_step0 := mat.NewVecDense(number_of_tokens, nil)
		risk_step0.MulVec(cov2, weights)
		risk := math.Sqrt(mat.Dot(weights, risk_step0))

		// Return sharpe ratio
		sharpe := -blended_return / risk

		if math.IsNaN(sharpe) {
			return 0.0
		}

		if math.IsInf(sharpe, 0) {
			return 0.0
		}

		return sharpe
	} // fcn definition complete

	// 11 - Call the optimizer
	var p0 []float64

	if len(h_array) > 0 { // sized same as number of tokens
		for i := 0; i < len(h_array); i++ {
			p0 = append(p0, 1/float64(len(h_array)))
		}
	} else {
		p0 = append(p0, 0.0)
	} // 1/number_of_tokens

	// 12 - Feed fcn into optimizer
	p := optimize.Problem{
		Func: fcn,
	}

	fmt.Print("13 - ABOUT TO CALL MINIMIZE")
	result, err := optimize.Minimize(p, p0, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = result.Status.Err(); err != nil {
		log.Fatal(err)
	}
	
	//var leftovertokens []string
	//var leftoveramounts []float64
	//var result_norm
	// Print results out
	fmt.Print("RAW WEIGHTS OPTIMIZED: ")
	fmt.Println(result)
	fmt.Print("------------------------------------")
	result_norm, leftovertokens, leftoveramounts := nrm_pool_wgts(result.X, pool_tkn0s, pool_tkn1s, pool_ratios, deployable_portfolio_array, own_pf_px)
	fmt.Print("FINAL WEIGHTS OPTIMIZED: ")
	fmt.Println(result_norm)
	fmt.Print("..sz: ")
	fmt.Println(len(result_norm))
	fmt.Println("..OPTIMIZATION COMPLETE..")

	if len(leftovertokens) == 0 && len(result_norm) == 0 && len(database.ownstartingportfolio) > 0 {
		// populate leftovers here
		for i:= 0; i < len(database.ownstartingportfolio); i++ {
			leftovertokens = append(leftovertokens,database.ownstartingportfolio[i].Token)
			leftoveramounts = append(leftoveramounts,float64(database.ownstartingportfolio[i].Amount))
		}
	}

	total_pf_val := 0.0
	for i:=0; i < len(deployable_portfolio_array); i++ {
		total_pf_val += float64(deployable_portfolio_array[i].Amount)*own_pf_px[i]
	}

	// Pack results into output struct array
	for i := 0; i < len(result_norm); i++ {
		amount := float32(total_pf_val * result_norm[i])
		yield := float32(0.06723)
		// search database
		if amount > 0 {
		for ii := 0; ii < len(database.currencyinputdata); ii++ {
			/*
			fmt.Print("SEARCHING FOR YIELD IN DB!!!: ")
			fmt.Print(ii)
			fmt.Print(" pair: ")
			fmt.Print(database.currencyinputdata[ii].Pair)
			fmt.Print(" | ")
			fmt.Print(database.currencyinputdata[ii].ROI_raw_est)
			fmt.Print(" | name: ")
			fmt.Print(database.currencyinputdata[ii].Pool)
			fmt.Print(" | tryna match: ")
			fmt.Print(pool_tkn0s[i] + "/" + pool_tkn1s[i])
			fmt.Print(" | our pool name: ")
			fmt.Print(pool_name_array[i])
			fmt.Println(" | ")
			*/
			if database.currencyinputdata[ii].Pair == (pool_tkn0s[i] + "/" + pool_tkn1s[i]) && database.currencyinputdata[ii].Pool == pool_name_array[i] {
	//			fmt.Print("MATCHED!!!!! ")
				yield = database.currencyinputdata[ii].ROI_raw_est
			}
		}

		optimised_pf = append(optimised_pf, OptimisedPortfolioRecord{pool_tkn0s[i] + "/" + pool_tkn1s[i], pool_name_array[i], amount, float32(result_norm[i]), yield, database.Risksetting})
	}
}

//	if totl > 0 && totl < 1 {
	//fmt.Print("NUMBER OF LEFTOVER TOKENS: ")
	//fmt.Print(len(leftovertokens))
	for i:= 0; i < len(leftovertokens); i++ {
		fmt.Print("adding leftover..")
		fmt.Print(leftovertokens[i])
		//amount := float32(420) // XX - get total pool amt
		yield := float32(0.0) // XX - get pool yield for that pool
		//leftover := 1 - totl
		pool := "NA"
		//	fmt.Print("Remainder to Uniswap DAI!!")
		// same loop searching yieold in db - slightly different
		pct := float32(0.0)
		if total_pf_val > 0 {
			pct = float32(leftoveramounts[i]/total_pf_val)			
		}
		//pct = float32(leftoveramounts[i]/total_pf_val)
		optimised_pf = append(optimised_pf, OptimisedPortfolioRecord{leftovertokens[i], pool, float32(leftoveramounts[i]), pct, yield, database.Risksetting})
	}

	for i := 0; i < len(optimised_pf); i++ {
		fmt.Print(optimised_pf[i].TokenOrPair)
		fmt.Print(optimised_pf[i].PercentageOfPortfolio)
		fmt.Print(optimised_pf[i].Pool)
		fmt.Print(optimised_pf[i].Amount)
	}


	return optimised_pf
}

/*
	fmt.Print("len len(h_array): ")
	fmt.Println(len(h_array))
	fmt.Println(number_of_days)
	fmt.Println(number_of_tokens)
*/

/*
	for i := 0; i < len(startingTokenTickers); i++ {
		fmt.Println(startingTokenTickers[i])
	}
*/
