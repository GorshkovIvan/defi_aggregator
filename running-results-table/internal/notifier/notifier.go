package notifier

import (
	"pusher/defi_aggregator/running-results-table/internal/db"

	"github.com/pusher/pusher-http-go"
)

type Notifier struct {
	notifyChannel chan<- bool
}

func notifier(database *db.Database, notifyChannel <-chan bool) {
	client := pusher.Client{
		AppID:   "1139323",
		Key:     "7885860875bb513c3e34",
		Secret:  "3633fcf50bba02630b0c",
		Cluster: "eu",
		Secure:  true,
	}
	// infinite loop for both results and currencyoutputtable
	// we have to put everything in this one for loop
	for {
		<-notifyChannel
		data := map[string][]db.Record{"results": database.GetRecords()}
		client.Trigger("results", "results", data)
		currencyoutputtable := map[string][]db.CurrencyInputData{"currencyoutputtable": database.GetCurrencyInputData()}
		client.Trigger("currencyoutputtable", "currencyoutputtable", currencyoutputtable)
	}
}
func New(database *db.Database) Notifier {
	notifyChannel := make(chan bool)
	go notifier(database, notifyChannel)
	return Notifier{
		notifyChannel,
	}
}
func (notifier *Notifier) Notify() {
	notifier.notifyChannel <- true
}
