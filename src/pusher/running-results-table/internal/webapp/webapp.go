package webapp
import (
    "net/http"
    "pusher/running-results-table/internal/db"
    "pusher/running-results-table/internal/notifier"
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
    r.Run()
}