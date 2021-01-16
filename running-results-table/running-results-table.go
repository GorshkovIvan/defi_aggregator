package main
import (
		"pusher/running-results-table/internal/db"
		"pusher/running-results-table/internal/notifier"
		"pusher/running-results-table/internal/webapp"
)
func main() {
		database := db.New()
		notifierClient := notifier.New(&database)
		webapp.StartServer(&database, &notifierClient)
}