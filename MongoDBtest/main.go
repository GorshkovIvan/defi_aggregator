package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func main() {

	addOwnPortfolioRecord("ETH", 100)

}

func addOwnPortfolioRecord(token string, amount float32) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:highyield4me@cluster0.seblt.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	Database := client.Database("De-Fi_Aggregator")
	ownstartingportfolio := Database.Collection("Own Portfolio Record")

	new_portfolio, err := ownstartingportfolio.InsertOne(ctx, bson.D{
		{Key: "Token", Value: token},
		{Key: "Amount", Value: amount},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(new_portfolio)

}
