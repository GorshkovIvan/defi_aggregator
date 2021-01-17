package webapp

import (
	"net/http"
	"pusher/defi_aggregator/running-results-table/internal/db"
	"pusher/defi_aggregator/running-results-table/internal/notifier"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer(database *db.Database, notifierClient *notifier.Notifier) {

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/results", func(c *gin.Context) {
		results := database.GetRecords()
		c.JSON(http.StatusOK, gin.H{
			"results": results,
		})
	})

	r.GET("/currencyoutputtable", func(c *gin.Context) {
		currencyoutputtable := database.GetCurrencyInputData()
		c.JSON(http.StatusOK, gin.H{
			"currencyoutputtable": currencyoutputtable,
		})
	})

	r.POST("/results", func(c *gin.Context) {
		var json db.Record
		if err := c.BindJSON(&json); err == nil {
			database.AddRecord(json)
			c.JSON(http.StatusCreated, json)
			notifierClient.Notify()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{})
		}
	})

	// post it to DB - table 2
	database.AddRecordfromAPI()
	database.AddRecordfromAPI2("USDT/USDC", 420420, 6969)
	database.AddRecordfromAPI2("HighYield4meToken", 1337, 420.69)

	// This is the backend algo
	database.RankBestCurrencies()

	// Next steps:
	// Get API pull to work with Notifier update frequency
	// Connect a real database
	// Create an ugly front end like Curve
	// Profit

	/*
		r.POST("/currencyoutputtable", func(c *gin.Context) {
			var json db.CurrencyInputData
			// {"ETH/DAI",420,0.069} // download data from API - here?

			if err := c.BindJSON(&json); err == nil {
				database.AddRecordfromAPI(json)
				c.JSON(http.StatusCreated, json)
				notifierClient.Notify()
			} else {
				c.JSON(http.StatusBadRequest, gin.H{})
			}
		})
	*/

	r.Run()
}
