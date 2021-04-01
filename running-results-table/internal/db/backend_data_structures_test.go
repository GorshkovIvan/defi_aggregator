package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)


func TestAddOwnPortfolioRecord(t *testing.T) {
	var id = addOwnPortfolioRecord("ETH", 100)

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
	optimisedportfolio := Database.Collection("Own Portfolio Record")

	cursor, err := optimisedportfolio.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var entries []bson.M
	if err = cursor.All(ctx, &entries); err != nil {
		log.Fatal(err)
	}

	flag := true

	for _, entry := range entries {
		if entry["_id"].(primitive.ObjectID).Hex() == id {
			flag = false
		}
	}

	if flag {
		t.Errorf("Entry was not added!")
	}

}

func TestAddOptimisedPortfolioRecord(t *testing.T) {
	var id = addOptimisedPortfolioRecord("DAI", "Uniswap", 2000, 0.42, 0.05, 0.2)

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
	optimisedportfolio := Database.Collection("Optimised Portfolio Record")

	cursor, err := optimisedportfolio.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var entries []bson.M
	if err = cursor.All(ctx, &entries); err != nil {
		log.Fatal(err)
	}

	flag := true

	for _, entry := range entries {
		if entry["_id"].(primitive.ObjectID).Hex() == id {
			flag = false
		}
	}

	if flag {
		t.Errorf("Entry was not added!")
	}

}


func TestAddHistoricalCurrencyData(t *testing.T) {
	var id = addHistoricalCurrencyData(420, 100, "DAI")

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
	optimisedportfolio := Database.Collection("Historical Currency Data")

	cursor, err := optimisedportfolio.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var entries []bson.M
	if err = cursor.All(ctx, &entries); err != nil {
		log.Fatal(err)
	}

	flag := true

	for _, entry := range entries {
		if entry["_id"].(primitive.ObjectID).Hex() == id {
			flag = false
		}
	}

	if flag {
		t.Errorf("Entry was not added!")
	}

}


func TestAddCurrencyInputData(t *testing.T) {
	var id = addCurrencyInputData("DAI/ETH", 42069, 2000, 0.42, "Uniswap", 0.3,
	0.2, 0.3, 0.2)

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
	optimisedportfolio := Database.Collection("Currency Input Data")

	cursor, err := optimisedportfolio.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var entries []bson.M
	if err = cursor.All(ctx, &entries); err != nil {
		log.Fatal(err)
	}

	flag := true

	for _, entry := range entries {
		if entry["_id"].(primitive.ObjectID).Hex() == id {
			flag = false
		}
	}

	if flag {
		t.Errorf("Entry was not added!")
	}

}



func TestRemoveRecordById(t *testing.T) {
	var id = addOwnPortfolioRecord("ETH", 100)
	removeRecordById("Own Portfolio Record", id)

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
	optimisedportfolio := Database.Collection("Own Portfolio Record")

	cursor, err := optimisedportfolio.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var entries []bson.M
	if err = cursor.All(ctx, &entries); err != nil {
		log.Fatal(err)
	}

	flag := true

	for _, entry := range entries {
		if entry["_id"].(primitive.ObjectID).Hex() == id {
			flag = false
		}
	}

	if flag == false {
		t.Errorf("Entry was not removed!")
	}
}



func TestDropEntireCollection(t *testing.T) {
	var id = addOwnPortfolioRecord("ETH", 100)
	dropEntireCollection("Own Portfolio Record")

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
	optimisedportfolio := Database.Collection("Own Portfolio Record")

	cursor, err := optimisedportfolio.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var entries []bson.M
	if err = cursor.All(ctx, &entries); err != nil {
		log.Fatal(err)
	}

	flag := true

	for _, entry := range entries {
		if entry["_id"].(primitive.ObjectID).Hex() == id {
			flag = false
		}
	}

	if flag == false {
		t.Errorf("Collection was not removed!")
	}
}

func TestReturnAttributeInCollection(t *testing.T) {
	dropEntireCollection("Own Portfolio Record")
	addOwnPortfolioRecord("ETH", 100)
	addOwnPortfolioRecord("ETH", 100)
	addOwnPortfolioRecord("ETH", 100)

	var array = returnAttributeInCollection("Own Portfolio Record", "Token")

	for _, value := range array {
		if value != "ETH" {
			t.Errorf("Incorrect array was returned!")
		}
	}
}
/*
func TestNewOptimisedPortfolio(t *testing.T) {
	new_database := New()
	newOptimisedPortfolio := NewOptimisedPortfolio(&new_database) // returns array of portfolios

	if newOptimisedPortfolio[0].TokenOrPair != "USD" {
		t.Errorf("Token error!")
	}

	if newOptimisedPortfolio[0].Pool != "Uniswap" {
		t.Errorf("Pool error!")
	}

	if newOptimisedPortfolio[0].Amount != 100 {
		t.Errorf("Pool error!")
	}

	if newOptimisedPortfolio[0].PercentageOfPortfolio != 1 {
		t.Errorf("Percentage of Portfolio error!")
	}

	if newOptimisedPortfolio[0].ROI_raw_est != 0.0125 {
		t.Errorf("ROI error!")
	}
}

func TestNewOptimisedPortfolioWithInputLengthZero(t *testing.T) {
	new_database := New()
	newOptimisedPortfolio := NewOptimisedPortfolio(&new_database) // returns array of portfolios

	if newOptimisedPortfolio[0].TokenOrPair != "USD" {
		t.Errorf("Token error!")
	}

	if newOptimisedPortfolio[0].Pool != "Uniswap" {
		t.Errorf("Pool error!")
	}

	if newOptimisedPortfolio[0].Amount != 100 {
		t.Errorf("Pool error!")
	}

	if newOptimisedPortfolio[0].PercentageOfPortfolio != 1 {
		t.Errorf("Percentage of Portfolio error!")
	}

	if newOptimisedPortfolio[0].ROI_raw_est != 0.0125 {
		t.Errorf("ROI error!")
	}

}

func TestNewOwnPortfolioRecord(t *testing.T) {
	ownPortfolioRecord := NewOwnPortfolioRecord("ETH", 100)

	if ownPortfolioRecord.Token != "ETH" {
		t.Errorf("Failed to save token name into OwnPortfolioRecord!")
	}

	if ownPortfolioRecord.Amount != 100 {
		t.Errorf("Failed to save token amount into OwnPortfolioRecord!")
	}

}

func TestNewHistoricalCurrencyDataFromRaw(t *testing.T) {
	var rawHistoricalData []UniswapDaily
	historicalCurrencyData := NewHistoricalCurrencyDataFromRaw("ETH", rawHistoricalData)

	if historicalCurrencyData.Ticker != "ETH" {
		t.Errorf("Failed to create new historical currency data from raw data")
	}

}

func TestNewHistoricalCurrencyData(t *testing.T) {
	historicalData := NewHistoricalCurrencyData()

	if historicalData.Date[0] != 1099999999999999 {
		t.Errorf("Failed to create Historical Currency Data with the correct Date")
	}
}

func TestNewCurrencyInputData(t *testing.T) {
	new_currencyinputdata := NewCurrencyInputData()

	if new_currencyinputdata.Pair != "ETH/DAI" {
		t.Errorf("fail!")
	}

	if new_currencyinputdata.PoolSize != 420000.69 {
		t.Errorf("fail!")
	}

	if new_currencyinputdata.PoolVolume != 4200.69 {
		t.Errorf("fail!")
	}

	if new_currencyinputdata.Yield != 0.08 {
		t.Errorf("fail!")
	}

	if new_currencyinputdata.Pool != "Uniswap" {
		t.Errorf("fail!")
	}

	if new_currencyinputdata.Volatility != -0.09 {
		t.Errorf("fail!")
	}

	if new_currencyinputdata.ROI_raw_est!= 0.4269 {
		t.Errorf("fail!")
	}

}

func TestNewCurrencyInputDataAct(t *testing.T) {
	new_currencyinputdataact := NewCurrencyInputDataAct("ETH/DAI", 10.0, 5, 0.10, "Uniswap",
		0.2, 0.08)

	if new_currencyinputdataact.Pair != "ETH/DAI" {
		t.Errorf("fail!")
	}

	if new_currencyinputdataact.PoolSize != 10.0 {
		t.Errorf("fail!")
	}

	if new_currencyinputdataact.PoolVolume != 5 {
		t.Errorf("fail!")
	}

	if new_currencyinputdataact.Yield != 0.10 {
		t.Errorf("fail!")
	}

	if new_currencyinputdataact.Volatility != 0.2 {
		t.Errorf("fail!")
	}

	if new_currencyinputdataact.ROI_raw_est != 0.08 {
		t.Errorf("fail!")
	}

}

func TestNew(t *testing.T) {
	new_database := New()

	if new_database.ownstartingportfolio == nil {
		t.Errorf("fail!")
	}

	if new_database.currencyinputdata == nil {
		t.Errorf("fail!")
	}

	if new_database.optimisedportfolio == nil {
		t.Errorf("fail!")
	}

	if new_database.Risksetting != 0 {
		t.Errorf("fail!")
	}

	if new_database.historicalcurrencydata == nil {
		t.Errorf("fail!")
	}
}

func TestAddRecord(t *testing.T) {
	new_database := New()
	new_record := OwnPortfolioRecord{"DAI", 10}
	new_database.AddRecord(new_record)

	if new_database.ownstartingportfolio[0] != new_record {
		t.Errorf("fail!")
	}
}

func TestAddRiskRecord(t *testing.T) {
	new_database := New()
	new_record := RiskWrapper{0.05}
	new_database.AddRiskRecord(new_record)

	if new_database.Risksetting != new_record.Risksettinginput {
		t.Errorf("fail!")
	}
}

func TestGetOptimisedPortfolio(t *testing.T) {
	new_database := New()
	optimised_portfolio := new_database.GetOptimisedPortfolio()

	if optimised_portfolio[0].TokenOrPair != "USD" {
		t.Errorf("fail!")
	}

	if optimised_portfolio[0].Pool != "Uniswap" {
		t.Errorf("fail!")
	}
	if optimised_portfolio[0].Amount != 100 {
		t.Errorf("fail!")
	}
	if optimised_portfolio[0].PercentageOfPortfolio != 1 {
		t.Errorf("fail!")
	}
	if optimised_portfolio[0].ROI_raw_est != 0.0125{
		t.Errorf("fail!")
	}
	if optimised_portfolio[0].Risksetting != 0.00 {
		t.Errorf("fail!")
	}
}
/*
func TestGetCurrencyInputData(t *testing.T) {
	new_database := New()
	new_database.currencyinputdata[0] = NewCurrencyInputData()
	currency_input_data := new_database.GetCurrencyInputData()

	if currency_input_data[0].Pair != "ETH/DAI" {
		t.Errorf("fail!")
	}

	if currency_input_data[0].PoolSize != 420000.69 {
		t.Errorf("fail!")
	}

	if currency_input_data[0].PoolVolume != 4200.69 {
		t.Errorf("fail!")
	}
	if currency_input_data[0].Yield != 0.05 {
		t.Errorf("fail!")
	}
	if currency_input_data[0].Pool != "Uniswap" {
		t.Errorf("fail!")
	}
	if currency_input_data[0].Volatility != -9.00 {
		t.Errorf("fail!")
	}

	if currency_input_data[0].ROI_raw_est != 42.69 {
		t.Errorf("fail!")
	}

}
*/

