package main

import (
	"pusher/defi_aggregator/running-results-table/internal/db"
	"pusher/defi_aggregator/running-results-table/internal/notifier"
	"pusher/defi_aggregator/running-results-table/internal/webapp"
)

func main() {
	database := db.New()
	notifierClient := notifier.New(&database)
	webapp.StartServer(&database, &notifierClient)
}
