package webapp

import (
	"fmt"
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

	//Returns currency for table2
	r.GET("/currencyoutputtable", func(c *gin.Context) {
		currencyoutputtable := database.GetCurrencyInputData()
		c.JSON(http.StatusOK, gin.H{
			"currencyoutputtable": currencyoutputtable,
		})
	})

	//Returns currency for table2
	r.GET("/results_original", func(c *gin.Context) {
		data2 := database.GetRawPortfolio()
		c.JSON(http.StatusOK, gin.H{
			"results_original": data2,
		})
	})

	// Ranked by ROI table - download data and rank currencies
	r.POST("/results", func(c *gin.Context) {
		var json db.OwnPortfolioRecord //		var json db.Record
		// fmt.Print("RUNNING ")
		if err := c.BindJSON(&json); err == nil {
			database.AddRecord(json)
			c.JSON(http.StatusCreated, json)
			notifierClient.Notify()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{})
		}
	})

	// Post data from slider into db
	r.POST("/results2", func(c *gin.Context) {
		var json db.RiskWrapper
		if err := c.BindJSON(&json); err == nil {
			fmt.Println("ADDING RISK RECORD FROM BUTTON!!")
			database.AddRiskRecord(json) // json
			c.JSON(http.StatusCreated, json)
			notifierClient.Notify()

			fmt.Println(database.Risksetting)

		} else {
			fmt.Println("ERROR IN PARSING JSON RISK SETTING!!")
			c.JSON(http.StatusBadRequest, gin.H{})
		}
	})

	/*
		// Optimised portfolio table
		r.GET("/results2", func(c *gin.Context) {
			Risksetting := database.AddRiskRecord(json) // json
			c.JSON(http.StatusOK, gin.H{
				"results2": Risksetting,
			})
		})
	*/

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
