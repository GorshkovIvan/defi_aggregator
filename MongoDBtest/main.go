package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)


func main() {
	//var id = addOwnPortfolioRecord("DAI", 69)
	//returnEntryById("Own Portfolio Record", id)
	//fmt.Println(returnAttributeInCollection("Historical Currency Data", "Price"))

	//removeRecordById("Own Portfolio Record", "6065b79b738be8435f30b458")
	//dropEntireCollection("Own Portfolio Record")
	/*

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:highyield4me@cluster0.tmmmg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
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
	testDatabase := Database.Collection("Test")

	test, err := testDatabase.Indexes().CreateOne(ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "createdAt", Value: 1}},
			Options: options.Index().SetExpireAfterSeconds(10),
		})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(test)
	 */
	/*

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:highyield4me@cluster0.tmmmg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	Database := client.Database("test2")

	fmt.Println(Database.ListCollectionNames(ctx, bson.M{}))
	*/

	//fmt.Println(isHistDataAlreadyDownloadedDatabase("token1"))



	//fmt.Println(returnAttributeInCollection("token1", "date"))
	//addOwnPortfolioRecord("DAI", 288)


}
/*
		cursor, err := ownstartingportfolio.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}

		var portfolioRecords []bson.M
		if err = cursor.All(ctx, &portfolioRecords); err != nil {
			log.Fatal(err)
		}

		//var last_id stri
		for _, portfolioRecord := range portfolioRecords {
			last_id := portfolioRecord["Token"]
			fmt.Println(last_id)
			fmt.Println(reflect.TypeOf(last_id))
		}


 */


func addOwnPortfolioRecord(token string, amount float32) string {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:highyield4me@cluster0.tmmmg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
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
	ownstartingportfolio := Database.Collection("Test")

	new_portfolio, err := ownstartingportfolio.InsertOne(ctx, bson.D{
		{Key: "createdAt", Value: time.Now()},
		{Key: "Token", Value: token},
		{Key: "Amount", Value: amount},
	})

	if err != nil {
		log.Fatal(err)
	}

	newID := new_portfolio.InsertedID
	hexID := newID.(primitive.ObjectID).Hex()

	return hexID

}

func removeRecordById(collection string, id string) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:highyield4me@cluster0.tmmmg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
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
	a_collection := Database.Collection(collection)

	objID, _ := primitive.ObjectIDFromHex(id)

	result, err := a_collection.DeleteOne(ctx, bson.M{"_id": objID})

	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

	if err != nil {
		log.Fatal(err)
	}
}

func dropEntireCollection(collectionName string) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:highyield4me@cluster0.tmmmg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
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
	collection := Database.Collection(collectionName)


	if err = collection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
}

func returnDateInCollection(collectionName string, attribute string) [] primitive.DateTime {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:highyield4me@cluster0.tmmmg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)


	Database := client.Database("test2")
	collection := Database.Collection(collectionName)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}


	var records []bson.M

	if err = cursor.All(ctx, &records); err != nil {
		log.Fatal(err)
	}

	var dates [] primitive.DateTime
	for _, record := range records {
		//fmt.Println(record)
		//fmt.Println(reflect.TypeOf(record[attribute]))
		date := record[attribute]
		//fmt.Println(attribute_value)
		//fmt.Println(reflect.TypeOf(attribute_value))
		//attributes = append(attributes, fmt.Sprint(attribute_value))
		dates = append(dates, date.(primitive.DateTime))
	}

	return dates
}


func returnEntryById(collectionName string, id string) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:highyield4me@cluster0.tmmmg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
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
	collection := Database.Collection(collectionName)

	objID, _ := primitive.ObjectIDFromHex(id)


	filterCursor, err := collection.Find(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Fatal(err)
	}
	var collectionFiltered []bson.M
	if err = filterCursor.All(ctx, &collectionFiltered); err != nil {
		log.Fatal(err)
	}

	fmt.Println(collectionFiltered)
}

func isHistDataAlreadyDownloadedDatabase(token string) bool {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:highyield4me@cluster0.tmmmg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	Database := client.Database("test2")

	array, err:= Database.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(array); i++ {
		if array[i] == token {
			// also add date check LATER : if latest date is within 24 hours of NOW database.historicalcurrencydata[i].
			/*
				fmt.Print("Checking if data already downloaded for: ")
				fmt.Print(token)
				fmt.Print("..Data found!!")
			*/
			return true
		}
	}
	return false
}
