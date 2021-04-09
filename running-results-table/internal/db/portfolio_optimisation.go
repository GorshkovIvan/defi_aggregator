package db

import (
	"fmt"
	"log"
	"math"

	//"github.com/cpmech/gosl/num"

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

	var listOfAvailablePairsWithoutConversion []string // clean starting portfolio for duplicates
	// var listOfAvailablePairswithConversion []string

	// Pack risk tolerance somewhere here
	if database.Risksetting == 0 {
		fmt.Println("Risk setting set to zero!")
	}

	// Remove non unique items
	for i := 0; i < len(database.ownstartingportfolio); i++ {
		if !stringInSlice(database.ownstartingportfolio[i].Token, listOfAvailablePairsWithoutConversion) {
			listOfAvailablePairsWithoutConversion = append(listOfAvailablePairsWithoutConversion, database.ownstartingportfolio[i].Token)
		} // add
	}

	// Now create all possible PAIRS from this list of UNIQUE tokens - PERMUTE
	// Use token1 as both USD and other tokens - i.e. 2nd token in LENDING POOLS is always USD
	// Query database for best ROI on items from this list
	// Pack recommended pools into a the optimisedportfolio

	var lambda_vals [100]float32
	lambda_vals[0] = 0.01

	for i := 1; i < 100; i++ {
		lambda_vals[i] = lambda_vals[i-1] + 0.01
	}

	// why are we running 100 iterations?
	// this is as per the example; can always change later
	for i := 0; i < 100; i++ {

		fcn := func(x_weights []float64) float64 {
			print := false

			// NEW ALGO
			/*
				ret_mat := mat.NewDense(4, 4, []float64{
					11.1, 14.2, 11.3, 11.4,
					12.1, 12.2, 18.3, 13.4,
					13.1, 11.2, 17.2, 12.4,
					16.1, 11.2, 13.3, 11.4,
				})
			*/
			/*
				ret_mat := mat.NewDense(4, 4, []float64{
					11.1, 14.2, 1.3, 111.4,
					11.1, 42.2, 38.3, 113.4,
					11.1, 561.2, 37.2, 412.4,
					11.1, 21.2, 33.3, 111.4,
				})*/

			//2 assets with zero volatility
			ret_mat := mat.NewDense(4, 4, []float64{
				11.1, 14.2, 31, 111.4,
				11.2, 16, 38.3, 111.3,
				11.3, 11, 37.2, 111.2,
				11.15, 21.2, 33.3, 111.1,
			})
			/*
				ret_mat := mat.NewDense(4, 4, []float64{
					41.1, 44.2, 51.3, 41.4,
					42.1, 42.2, 58.3, 53.4,
					43.1, 41.2, 57.2, 72.4,
					46.1, 41.2, 53.3, 61.4,
				})*/

			ret_mat_pct := mat.NewDense(4, 4, nil)

			for jj := 0; jj < 4; jj++ {
				for ii := 0; ii < 4; ii++ {
					if ii > 0 {
						ret_mat_pct.Set(ii, jj, ret_mat.At(ii, jj)/ret_mat.At(ii-1, jj)-1.0)
					} else {
						ret_mat_pct.Set(ii, jj, 0.0)
					}
				} // ii
			} // jj

			v0 := ret_mat.At(3, 0)/ret_mat.At(0, 0) - 1
			v1 := ret_mat.At(3, 1)/ret_mat.At(0, 1) - 1
			v2 := ret_mat.At(3, 2)/ret_mat.At(0, 2) - 1
			v3 := ret_mat.At(3, 3)/ret_mat.At(0, 3) - 1
			/*
				fmt.Println("RETURNS: ")
				fmt.Print(v0)
				fmt.Print(" | ")
				fmt.Print(v1)
				fmt.Print(" | ")
				fmt.Print(v2)
				fmt.Print(" | ")
				fmt.Println(v3)
			*/
			ret := mat.NewVecDense(4, []float64{v0, v1, v2, v3}) // vector of returns
			if print {
				fmt.Print("ret: ")
				fmt.Println(ret)
			}
			/*
				vol := mat.NewVecDense(4, []float64{0.05, 0.15, 0.25, 03}) // vector of volatility
				if print {
					fmt.Print("vol: ")
					fmt.Println(vol)
				}
			*/
			/*
				fmt.Print("x_weights before normalisation: ")
				fmt.Println(x_weights)
			*/
			for j := 0; j < len(x_weights); j++ {
				if x_weights[j] < 0 {
					x_weights[j] = 0
				}
			}

			totl := sum(x_weights)
			for j := 0; j < len(x_weights); j++ {
				x_weights[j] = x_weights[j] / totl
			}
			/*
				fmt.Print("x_weights after normalisation: ")
				fmt.Println(x_weights)
			*/
			weights := mat.NewVecDense(4, x_weights) // vector of portfolio weights
			//weights := mat.NewVecDense(4, []float64{0.25, 0.25, 0.25, 0.25}) // vector of portfolio weights
			if print {
				fmt.Print("weights: ")
				fmt.Println(weights)
			}
			var cov *mat.SymDense = mat.NewSymDense(4, nil)
			cov.Reset()
			if print {
				fmt.Println("ISEMPTY:")
				fmt.Println(cov.IsEmpty())
			}

			if print {
				fmt.Print("ret in matrix form:")
				fmt.Println(ret_mat)
				fmt.Println(ret_mat_pct)
			}
			stat.CovarianceMatrix(cov, ret_mat_pct.T(), nil) //TRANSPOSE KEEP OR DELETE - ?
			if print {
				fmt.Print("covariance matrix return: ")
				fmt.Println(cov)
			}
			blended_return := mat.Dot(ret, weights)
			if print {
				fmt.Println("blended return: ")
				fmt.Println(blended_return)
			}
			risk_step0 := mat.NewVecDense(4, nil)
			risk_step0.MulVec(cov, weights)
			if print {
				fmt.Print("Convolution step 1: ")
				fmt.Print(risk_step0)
			}
			risk := math.Sqrt(mat.Dot(weights, risk_step0))
			if print {
				fmt.Print("risk 2: ")
				fmt.Println(risk)
			}

			return 1 * (float64(lambda_vals[i])*risk - (1-float64(lambda_vals[i]))*blended_return)
		}

		p := optimize.Problem{
			Func: fcn,
		}
		/*
			xa, xb := 0.0, 0.11
			solver := num.NewBrent(p, nil)
			xo := solver.Root(xa, xb)
			fmt.Print("xo")
			fmt.Println(xo)
		*/
		var p0 = []float64{0.25, 0.25, 0.25, 0.25} // initial value for mu

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

		xxx := []float64{result_norm.Location.X[0], result_norm.Location.X[1], result_norm.Location.X[2], result_norm.Location.X[3]}

		fmt.Println("Optimal return: ")
		fmt.Println(fcn(xxx))

		fmt.Print("p0:")
		fmt.Println(p0)

		fmt.Print("lambda:")
		fmt.Println(lambda_vals[i])
	}

	return NewOptimisedPortfolio(database)
}
