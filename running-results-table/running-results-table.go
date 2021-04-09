package main

import (
	"fmt"
	"pusher/defi_aggregator/running-results-table/internal/db"
	"pusher/defi_aggregator/running-results-table/internal/notifier"
	"pusher/defi_aggregator/running-results-table/internal/webapp"
)

func main() {
	database := db.New()

	// for loading the optimised portfolio without having to load all the data
	results := database.GetOptimisedPortfolio()
	if false {
		fmt.Print(results)
	}
	//
	notifierClient := notifier.New(&database)
	webapp.StartServer(&database, &notifierClient)
}
