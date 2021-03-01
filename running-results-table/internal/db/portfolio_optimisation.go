package db

import "fmt"

func OptimisePortfolio(database *Database) []OptimisedPortfolioRecord {
	// rawportfolio []OwnPortfolioRecord, risktolerance float32,
	var listOfAvailablePairsWithoutConversion []string // clean starting portfolio for duplicates
	// var listOfAvailablePairswithConversion []string

	// Pack risk tolerance somewhere here
	if database.risksetting == 0 {
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

	// For now just a placeholder for result
	return NewOptimisedPortfolio(database.ownstartingportfolio)
}
