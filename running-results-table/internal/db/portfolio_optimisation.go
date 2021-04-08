package db

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/optimize"
	"math"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
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

	// NEW ALGO
	ret := mat.NewVecDense(4, []float64{0.1, 0.2, 0.3, 0.4}) // vector of returns
	fmt.Print("ret: ")
	fmt.Println(ret)

	vol := mat.NewVecDense(4, []float64{0.05, 0.15, 0.25, 03}) // vector of volatility
	fmt.Print("vol: ")
	fmt.Println(vol)

	//weights_array := []float64{0.2, 0.2, 0.2, 0.2, 0.2, 0.2}
	weights := mat.NewVecDense(4, []float64{0.25, 0.25, 0.25, 0.25}) // vector of portfolio weights
	fmt.Print("weights: ")
	fmt.Println(weights)

	//	cov := mat.NewDense(4, 4, nil)
	//	fmt.Print("cov: ")
	//	fmt.Println(cov)

	var cov *mat.SymDense = mat.NewSymDense(4, nil)
	cov.Reset()
	fmt.Println("ISEMPTY:")
	fmt.Println(cov.IsEmpty())

	//ret2 = matrix from ret
	//weights2 = array of float64 (weights)
	//	ret_mat := ret.T()
	ret_mat := mat.NewDense(4, 4, []float64{
		41.1, 44.2, 51.3, 41.4,
		42.1, 42.2, 58.3, 53.4,
		43.1, 41.2, 57.2, 72.4,
		46.1, 41.2, 53.3, 61.4,
	})
	fmt.Print("ret in matrix form:")
	fmt.Println(ret_mat)
	//fmt.Println(weights_array)

	stat.CovarianceMatrix(cov, ret_mat.T(), nil) //covariance matrix of returns, transposed
	fmt.Print("covariance matrix return: ")
	fmt.Println(cov)

	blended_return := mat.Dot(ret, weights)
	fmt.Println("blended return: ")
	fmt.Println(blended_return)

	x := mat.NewVecDense(4, nil)
	x.MulVec(cov, weights)
	fmt.Print("Convolution step 1: ")
	fmt.Print(x)

	portfolio_volatility := 0.0 
	fmt.Print("portfolio_volatility 1: ")
	fmt.Println(portfolio_volatility)

	portfolio_volatility = math.Sqrt(mat.Dot(weights, x))

	fmt.Print("portfolio_volatility 2: ")
	fmt.Println(portfolio_volatility)
	
	var lambda_vals [99]float32
		lambda_vals[0] = 0.01
	
		for i := 1; i < 100;i++ {
			lambda_vals[i] = lambda_vals[i-1] + 0.01 
		}

		MeanVarB := mat.NewDense(101, 2, nil)
	
		w_lower := 0
		w_upper := 1
		
		// why are we running 100 iterations?
		for i := 1; i < 101; i++ { 
			//lambda := lambda_vals[i]
			fcn := func(x []float64) float64 {
				return float64(lambda_vals[i])*portfolio_volatility + (1-float64(lambda_vals[i]))*blended_return
			}

			p := optimize.Problem{
				Func: fcn,
			}

			//var c mat.Dense
			//c.Mul(weights.T(), sigma)
			//c.Mul(sigma, weights)

			//var meth = &optimize.Newton{} // meth - does this do anything?
			var p0 = []float64{1} // initial value for mu
			//	weights := mat.NewVecDense(4, []float64{0.25, 0.25, 0.25, 0.25}) // vector of portfolio weights

			result, err := optimize.Minimize(p, p0, nil, nil)
			if err != nil {
				log.Fatal(err)
			}
			if err = result.Status.Err(); err != nil {
				log.Fatal(err)
			}

			//MeanVarB[i,:]= [evaluate(ret),evaluate(risk)] - need to change from python?
		}


	return NewOptimisedPortfolio(database)
}