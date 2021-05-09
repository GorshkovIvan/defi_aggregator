package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
	"time"
)


func main() {
	now := time.Now()
	append_record_to_database("Uniswap", "ETH", "DAI", now.Unix(), 2000, 1000)


	//create_new_hist_volume_poolsz_entry("Balancer", "Token0", "Token1")

	fmt.Println(get_oldest_timestamp_from_db("Uniswap", "ETH", "DAI"))


	//var id = addOwnPortfolioRecord("DAI", 69)
	//returnEntryById("Own Portfolio Record", id)
	//fmt.Println(returnAttributeInCollection("Historical Currency Data", "Price"))
	//returnDatesInCollection("Uniswap ETH DAI")
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


	Database := client.Database("De-Fi_Aggregator")
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

	if len(collectionFiltered) == 0 {
		fmt.Println("already there")
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

func returnDatesInCollection(collectionName string) []int64 {
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

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var records []bson.M

	if err = cursor.All(ctx, &records); err != nil {
		log.Fatal(err)
	}

	var dates []int64
	for _, record := range records {
		//fmt.Println(record)
		//fmt.Println(reflect.TypeOf(record["Date"]))
		date := record["date"]
		//fmt.Println(date)
		//fmt.Println(reflect.TypeOf(date))
		//attributes = append(attributes, fmt.Sprint(attribute_value))
		dates = append(dates, date.(int64))
	}

	return dates
}
func append_record_to_database(pool string, token0 string, token1 string, date int64, trading_volume_usd int64, pool_sz_usd int64) string {
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
	s := []string{pool, token0, token1}
	v := strings.Join(s, " ")
	optimisedportfolio := Database.Collection(v)

	new_portfolio, err := optimisedportfolio.InsertOne(ctx, bson.D{
		{Key: "pool", Value: pool},
		{Key: "token0", Value: token0},
		{Key: "token1", Value: token1},
		{Key: "date", Value: date},
		{Key: "trading_Volume_USD", Value: trading_volume_usd},
		{Key: "pool_sz_USD", Value: pool_sz_usd},
	})

	if err != nil {
		log.Fatal(err)
	}

	newID := new_portfolio.InsertedID
	hexID := newID.(primitive.ObjectID).Hex()

	return hexID
}

func get_oldest_timestamp_from_db(pool string, token0 string, token1 string) int64 {
	s := []string{pool, token0, token1}
	v := strings.Join(s, " ")
	dates := returnDatesInCollection(v)

	min := int64(0)
	for _, v := range dates {
		if v < min {
			min = v
		}
	}

	if min == 0 {
		log.Fatal()
	}

	return min
}

func create_new_hist_volume_poolsz_entry(pool string, token0 string, token1 string) string {
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
	s := []string{pool, token0, token1, "hist volume poolsz"}
	v := strings.Join(s, " ")
	optimisedportfolio := Database.Collection(v)

	new_portfolio, err := optimisedportfolio.InsertOne(ctx, bson.D{
		{Key: "pool", Value: pool},
		{Key: "token0", Value: token0},
		{Key: "token1", Value: token1},
	})

	if err != nil {
		log.Fatal(err)
	}

	newID := new_portfolio.InsertedID
	hexID := newID.(primitive.ObjectID).Hex()

	return hexID
}