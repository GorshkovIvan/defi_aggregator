package webapp

import (
	"net/http"
	"pusher/defi_aggregator/running-results-table/internal/db"
	"pusher/defi_aggregator/running-results-table/internal/notifier"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// for adding data after a certain number of seconds
	"time"
)

func StartServer(database *db.Database, notifierClient *notifier.Notifier) {

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/currencyoutputtable", func(c *gin.Context) {
		currencyoutputtable := database.GetCurrencyInputData()
		c.JSON(http.StatusOK, gin.H{
			"currencyoutputtable": currencyoutputtable,
		})
	})

	// Optimised portfolio table
	r.GET("/results", func(c *gin.Context) {
		results := database.GetOptimisedPortfolio() //results := database.GetRecords()
		c.JSON(http.StatusOK, gin.H{
			"results": results,
		})
	})

	// Ranked by ROI table - download data and rank currencies
	r.POST("/results", func(c *gin.Context) {
		var json db.OwnPortfolioRecord //		var json db.Record
		if err := c.BindJSON(&json); err == nil {
			database.AddRecord(json)
			c.JSON(http.StatusCreated, json)
			notifierClient.Notify()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{})
		}
	})

	// Ranked by ROI table - download data and rank currencies
	r.POST("/results", func(c *gin.Context) {
		var json db.OwnPortfolioRecord //		var json db.Record
		if err := c.BindJSON(&json); err == nil {
			database.AddRecord(json)
			c.JSON(http.StatusCreated, json)
			notifierClient.Notify()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{})
		}
	})

	// Run just once
	time.AfterFunc(10*time.Second, func() {
		database.AddRecordfromAPI()
		database.RankBestCurrencies()
		notifierClient.Notify()
	})

	//	database.AddRecordfromAPI() 	// post it to DB - table 2
	//	database.RankBestCurrencies() 	// This is the backend algo
	/*
		time.AfterFunc(5*time.Second, func() {
			database.AddRecordfromAPI()
			database.RankBestCurrencies() // backend algo
			notifierClient.Notify()
		})
	*/

	r.Run()
}
