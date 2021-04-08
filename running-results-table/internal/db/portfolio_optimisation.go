package db

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	//	"math"
)

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

	// NEW ALGO
	ret := mat.NewVecDense(4, []float64{0.1, 0.2, 0.3, 0.4}) // vector of returns
	fmt.Print("ret: ")
	fmt.Println(ret)

	vol := mat.NewVecDense(4, []float64{0.05, 0.15, 0.25, 03}) // vector of volatility
	fmt.Print("vol: ")
	fmt.Println(vol)

	weights := mat.NewVecDense(4, []float64{0.25, 0.25, 0.2}) // vector of portfolio weights
	fmt.Print("weights: ")
	fmt.Println(weights)

	cov := mat.NewDense(4, 4, nil)
	fmt.Print("cov: ")
	fmt.Println(cov)

	// cov = CovarianceMatrix(dst *mat.SymDns mat.Matrix,weights) // (vol) // covariance matrix of returns?
	var blended_return mat.Dense
	blended_return.Mul(ret.T(), weights)
	fmt.Print("blended return: ")
	fmt.Println(blended_return)
	/*
		portfolio_volatility := mat.Dot(weights.T(), mat.Dot(cov, weights))
		fmt.Print("portfolio_volatility: ")
		fmt.Println(portfolio_volatility)

		portfolio_volatility_r := new(mat.Dense)
		portfolio_volatility_r.Apply(func(i, j int, v float64) float64 { return math.Sqrt(v) }, portfolio_volatility)

		fmt.Print("portfolio_volatility_r: ")
		fmt.Println(portfolio_volatility_r)
	*/
	lambda := 0.0 // weighting parameter - https://jump.dev/Convex.jl/stable/examples/portfolio_optimization/portfolio_optimization2/
	fmt.Print("lambda: ")
	fmt.Println(lambda)

	return NewOptimisedPortfolio(database)
}
