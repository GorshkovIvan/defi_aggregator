package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func append_record_to_database(pool string, tokens []string, date int64, trading_volume_usd int64, pool_sz_usd int64, fees int64, weighted_av_ir float64, util_rate float64) string {
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
	var name []string
	token0 := "N/A"
	token1 := "N/A"
	token2 := "N/A"
	token3 := "N/A"
	token4 := "N/A"
	token5 := "N/A"
	token6 := "N/A"
	token7 := "N/A"

	name = append(name, pool)
	for i := 0; i < len(tokens); i++ {
		name = append(name, tokens[i])
	}

	for i := 0; i < 8; i++ {
		switch i {
		case 0:
			if i < len(tokens) {
				token0 = tokens[i]
			} else {
				name = append(name, "token0")
			}
		case 1:
			if i < len(tokens) {
				token1 = tokens[i]
			} else {
				name = append(name, "token1")
			}
		case 2:
			if i < len(tokens) {
				token2 = tokens[i]
			} else {
				name = append(name, "token2")
			}
		case 3:
			if i < len(tokens) {
				token3 = tokens[i]
			} else {
				name = append(name, "token3")
			}
		case 4:
			if i < len(tokens) {
				token4 = tokens[i]
			} else {
				name = append(name, "token4")
			}
		case 5:
			if i < len(tokens) {
				token5 = tokens[i]
			} else {
				name = append(name, "token5")
			}
		case 6:
			if i < len(tokens) {
				token6 = tokens[i]
			} else {
				name = append(name, "token6")
			}
		case 7:
			if i < len(tokens) {
				token7 = tokens[i]
			} else {
				name = append(name, "token7")
			}
		}
	}
	v := strings.Join(name, " ")

	optimisedportfolio := Database.Collection(v)

	new_portfolio, err := optimisedportfolio.InsertOne(ctx, bson.D{
		{Key: "pool", Value: pool},
		{Key: "token0", Value: token0},
		{Key: "token1", Value: token1},
		{Key: "token2", Value: token2},
		{Key: "token3", Value: token3},
		{Key: "token4", Value: token4},
		{Key: "token5", Value: token5},
		{Key: "token6", Value: token6},
		{Key: "token7", Value: token7},
		{Key: "date", Value: date},
		{Key: "trading_Volume_USD", Value: trading_volume_usd},
		{Key: "pool_sz_USD", Value: pool_sz_usd},
		{Key: "fees", Value: fees},
		{Key: "weighted_av_IR", Value: weighted_av_ir},
	})

	if err != nil {
		log.Fatal(err)
	}

	newID := new_portfolio.InsertedID
	hexID := newID.(primitive.ObjectID).Hex()

	return hexID
}

func get_newest_timestamp_for_token_from_db(token string) int64 {
	dates := returnDatesInCollection(token)

	max := int64(0)
	for _, v := range dates {
		if v > max {
			max = v
		}
	}

	return max
}


func get_newest_timestamp_from_db(pool string, tokens []string) int64 {
	var name []string
	name = append(name, pool)
	for i := 0; i < len(tokens); i++ {
		name = append(name, tokens[i])
	}

	for i := len(tokens); i < 8; i++ {
		var tokenNum []string
		tokenNum = append(tokenNum, "token")
		tokenNum = append(tokenNum, strconv.Itoa(i))
		tokenJoined := strings.Join(tokenNum, "")
		name = append(name, tokenJoined)
	}
	v := strings.Join(name, " ")
	fmt.Print(v)
	dates := returnDatesInCollection_pools(v)

	max := int64(0)
	for _, v := range dates {
		if v > max {
			max = v
		}
	}

	return max
}

func retrieve_hist_pool_sizes_volumes_fees_ir(pool string, tokens []string) (dates []int64, tradingvolumes []int64, poolsizes []int64, fees []int64, ir []float64, util_rate []float64) {
	var name []string
	name = append(name, pool)
	for i := 0; i < len(tokens); i++ {
		name = append(name, tokens[i])
	}

	for i := len(tokens); i < 8; i++ {
		var tokenNum []string
		tokenNum = append(tokenNum, "token")
		tokenNum = append(tokenNum, strconv.Itoa(i))
		tokenJoined := strings.Join(tokenNum, "")
		name = append(name, tokenJoined)
	}
	v := strings.Join(name, " ")

	dates = returnAttributeInCollectionAsInt64(v, "date")
	tradingvolumes = returnAttributeInCollectionAsInt64(v, "trading_Volume_USD")
	poolsizes = returnAttributeInCollectionAsInt64(v, "pool_sz_USD")
	fees = returnAttributeInCollectionAsInt64(v, "fees")
	ir = returnAttributeInCollectionAsFloat64(v, "weighted_av_IR")
	util_rate = returnAttributeInCollectionAsFloat64(v, "weighted_av_IR")
	return dates, tradingvolumes, poolsizes, fees, ir, util_rate
}

func isAave1RecordsInDb() bool {
	names := getCollectionNames("De-Fi Aggregator")

	for _, name := range names {
		fmt.Println(name)
		if name == "Aave1 USDC USDC" {
			return true
		}
	}
	return false
}

func getCollectionNames(database string) []string {
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
	collection, err := Database.ListCollections(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var collectionNames []bson.M
	err = collection.All(ctx, &collectionNames)
	if err != nil {
		log.Fatal(err)
	}

	var allNames []string
	for _, name := range collectionNames {
		aName := name["name"]
		allNames = append(allNames, fmt.Sprint(aName))
	}

	return allNames
}

func returnAttributeInCollectionAsFloat64(collectionName string, attribute string) []float64 {
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

	var array []float64
	for _, record := range records {
		//fmt.Println(record)
		//fmt.Println(reflect.TypeOf(record["Date"]))
		val := record[attribute]
		//fmt.Println(date)
		//fmt.Println(reflect.TypeOf(date))
		//attributes = append(attributes, fmt.Sprint(attribute_value))
		array = append(array, val.(float64))
	}

	return array
}

func returnAttributeInCollectionAsInt64(collectionName string, attribute string) []int64 {
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

	var array []int64
	for _, record := range records {
		val := record[attribute]
		array = append(array, val.(int64))
	}

	return array
}

func addHistoricalCurrencyData(date int64, price float32, CollectionOrTicker string) string {
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
	historicaldata := Database.Collection(CollectionOrTicker)

	// check if date exists in that collection - if yes return "already exists"
	cursor, err := historicaldata.Find(ctx, bson.M{"Date": date})
	if err != nil {
		log.Fatal(err)
	}

	var collectionFiltered []bson.M
	err = cursor.All(ctx, &collectionFiltered)
	if err != nil {
		log.Fatal(err)
	}

//	fmt.Print("Collection Filtered: ")
//	fmt.Println(collectionFiltered)

	if len(collectionFiltered) > 0 {
		return "data already there"
	}

	new_data, err := historicaldata.InsertOne(ctx, bson.D{
		{Key: "Date", Value: date},
		{Key: "Price", Value: price},
	})

	if err != nil {
		log.Fatal(err)
	}

	newID := new_data.InsertedID
	hexID := newID.(primitive.ObjectID).Hex()

	return hexID
}

func addCurrencyInputData(pair string, poolsize float32, poolvolume float32, yield float32, pool string,
	volatility float32, roi_raw_est float32, roi_vol_adj_est float32, roi_hist float32) string {
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
	ownstartingportfolio := Database.Collection("Currency Input Data")

	new_data, err := ownstartingportfolio.InsertOne(ctx, bson.D{
		{Key: "Pair", Value: pair},
		{Key: "Pool Size", Value: poolsize},
		{Key: "Pool Volume", Value: poolvolume},
		{Key: "Yield", Value: yield},
		{Key: "Pool", Value: pool},
		{Key: "Volatility", Value: volatility},
		{Key: "ROI Raw Estimation", Value: roi_raw_est},
		{Key: "ROI Vol Adjusted Estimation", Value: roi_vol_adj_est},
		{Key: "ROI History", Value: roi_hist},
	})

	if err != nil {
		log.Fatal(err)
	}

	newID := new_data.InsertedID
	hexID := newID.(primitive.ObjectID).Hex()

	return hexID
}

/*
func removeRecordById(collectionName string, id string) {
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
	a_collection := Database.Collection(collectionName)

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

	Database := client.Database("test2")
	collection := Database.Collection(collectionName)

	if err = collection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
}
*/


func returnDatesInCollection_pools(collectionName string) []int64 {
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
		date := record["date"]
		dates = append(dates, date.(int64))
	}

	return dates
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

	var dates []int64
	for _, record := range records {
		date := record["Date"]
		dates = append(dates, date.(int64))
	}

	return dates
}

func returnPricesInCollection(collectionName string) []float64 {
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

	var prices []float64
	for _, record := range records {
		price := record["Price"]
		prices = append(prices, price.(float64))
	}

	return prices
}

func returnAttributeInCollection(collectionName string, attribute string) []string {
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

	var attributes []string
	for _, record := range records {
		attribute_value := record[attribute]
		//fmt.Println(attribute_value)
		//fmt.Println(reflect.TypeOf(attribute_value))
		attributes = append(attributes, fmt.Sprint(attribute_value))
	}

	return attributes
}

/*
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

	Database := client.Database("test2")
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
*/
type OwnPortfolioRecord struct {
	Token  string  `json:"token"`
	Amount float32 `json:"amount"`
}

// Output of OwnPortfolioRecord
type OptimisedPortfolioRecord struct {
	TokenOrPair           string  `json:"tokenorpair"`
	Pool                  string  `json:"pool"`
	Amount                float32 `json:"amount"`
	PercentageOfPortfolio float32 `json:"percentageofportfolio"`
	ROI_raw_est           float32 `json:"roi_estimate"`
	Risksetting           float32 `json:"risk_setting"`
}

func NewOptimisedPortfolio(database *Database) []OptimisedPortfolioRecord {

	// rawportfolio []OwnPortfolioRecord

	var optimisedportfolio []OptimisedPortfolioRecord
	for i := 0; i < len(database.ownstartingportfolio); i++ {
		// same token
		// placeholder values for now
		optimisedportfolio = append(optimisedportfolio, OptimisedPortfolioRecord{database.ownstartingportfolio[i].Token, "Uniswap", database.ownstartingportfolio[i].Amount, 0.420, 0.069, database.Risksetting})
	}

	// For some reason crashes with ZERO elements
	if len(database.ownstartingportfolio) == 0 {
		optimisedportfolio = append(optimisedportfolio, OptimisedPortfolioRecord{"USD", "Uniswap", 100, 1, 0.0125, database.Risksetting})
	}
	fmt.Println("Returning optimised portfolio...")
	return optimisedportfolio
}

// For adding own portfolio records
func NewOwnPortfolioRecord(token string, amount float32) OwnPortfolioRecord { // , pool_sz float32
	return OwnPortfolioRecord{token, amount} // , pool_sz
}

func NewHistoricalCurrencyDataFromRaw(token string, rawhistoricaldata []UniswapDaily) HistoricalCurrencyData {
	var historicaldata HistoricalCurrencyData

	// To DO: Date checking
	for i := 0; i < len(rawhistoricaldata); i++ {
		historicaldata.Date = append(historicaldata.Date, int64(rawhistoricaldata[i].Date))

		price, _ := strconv.ParseFloat(rawhistoricaldata[i].PriceUSD, 32)

		historicaldata.Price = append(historicaldata.Price, float32(price))
	}

	historicaldata.Ticker = token
	return historicaldata

}

func NewHistoricalCurrencyData() HistoricalCurrencyData {
	historicaldata := HistoricalCurrencyData{}
	//historicaldata.Date = append(historicaldata.Date, "Mon Jan 2 15:04:05 MST 2006")
	historicaldata.Date = append(historicaldata.Date, 1099999999999999)
	historicaldata.Price = append(historicaldata.Price, float32(420.69))
	historicaldata.Ticker = "ETH"

	return historicaldata
}

type HistoricalCurrencyData struct {
	Date   []int64   `default0:"10099999999999" json:"date"` // `default0:"Mon Jan 2 15:04:05 MST 2006" json:"date"`
	Price  []float32 `default0:"420.69" json:"price"`
	Ticker string    `default0:"ETH" json:"ticker"`
}

// Current
type CurrencyInputData struct {
	Pair            string  `default0:"ETH/DAI" json:"backend_pair"`
	PoolSize        float32 `default0:"420000.69" json:"backend_poolsize"`
	PoolVolume      float32 `default0:"4200.69" json:"backend_volume"`
	Yield           float32 `default0:"0.05" json:"backend_yield"`
	Pool            string  `default0:"Uniswap" json:"pool_source"`
	Volatility      float32 `default0:"-9.00" json:"volatility"`
	ROI_raw_est     float32 `default0:"42.69%" json:"ROIestimate"`
	ROI_vol_adj_est float32 `default0:"42.69%" json:"ROIvoladjest"`
	ROI_hist        float32 `default0:"42.69%" json:"ROIhist"`
}

func NewCurrencyInputData() CurrencyInputData {
	currencyinputdata := CurrencyInputData{}
	currencyinputdata.Pair = "ETH/DAI"
	currencyinputdata.PoolSize = 420000.69
	currencyinputdata.PoolVolume = 4200.69
	currencyinputdata.Yield = 0.08
	currencyinputdata.Pool = "Uniswap"
	currencyinputdata.Volatility = -0.09
	currencyinputdata.ROI_raw_est = 0.4269
	currencyinputdata.ROI_hist = 0.4269
	return currencyinputdata
}

func NewCurrencyInputDataAct(pair string, poolSz float32, poolVolume float32, yield float32, pool string, volatility float32, roi float32) CurrencyInputData {
	currencyinputdata := CurrencyInputData{}
	currencyinputdata.Pair = pair

	currencyinputdata.PoolSize = poolSz
	currencyinputdata.PoolVolume = poolVolume

	currencyinputdata.Yield = yield
	currencyinputdata.Pool = pool
	currencyinputdata.Volatility = volatility
	currencyinputdata.ROI_raw_est = roi
	currencyinputdata.ROI_hist = 0.0
	currencyinputdata.ROI_vol_adj_est = 0.0

	return currencyinputdata
}

type RiskWrapper struct {
	Risksettinginput float32 `json:"risk_setting"`
}

type Database struct {
	// Data structure for Optimisation
	ownstartingportfolio []OwnPortfolioRecord       // for portfolio optimisation table
	optimisedportfolio   []OptimisedPortfolioRecord // for storing output of ownstartingportfolio
	Risksetting          float32

	// Data structure for Ranking
	currencyinputdata      []CurrencyInputData      // LATEST currency pair info
	historicalcurrencydata []HistoricalCurrencyData // historical time series
}

func New() Database {
	ownstartingportfolio := make([]OwnPortfolioRecord, 0)
	optimisedportfolio := make([]OptimisedPortfolioRecord, 0)
	Risksetting := 0.00

	currencyinputdata := make([]CurrencyInputData, 0)
	historicalcurrencydata := make([]HistoricalCurrencyData, 0)

	return Database{ownstartingportfolio, optimisedportfolio, float32(Risksetting), currencyinputdata, historicalcurrencydata}
}

// Add OWN PORTFOLIO data
func (database *Database) AddRecord(r OwnPortfolioRecord) {
	fmt.Print("APPENDING OWN PORTFOLIO ITEM!!!!!!!!!!!!!!!!!!")
	database.ownstartingportfolio = append(database.ownstartingportfolio, r)
	fmt.Print(r.Token)
	fmt.Print(r.Amount)
	fmt.Print(" !!!!!!!!!! 999991 ")
	fmt.Print("..At this point, len of own starting portfolio is: ")
	fmt.Print(len(database.ownstartingportfolio))
	for i := 0; i < len(database.ownstartingportfolio);i++ {
		fmt.Print(database.ownstartingportfolio[i].Token)
		fmt.Print(" | ")
		fmt.Println(database.ownstartingportfolio[i].Amount)
	}
}

func (database *Database) AddRiskRecord(risk RiskWrapper) {
	database.Risksetting = risk.Risksettinginput
}

// Retrieve Data
func (database *Database) GetOptimisedPortfolio() []OptimisedPortfolioRecord {
//	fmt.Print("CHECKPOINT len of db being returned: ")
//	fmt.Print(len(database.ownstartingportfolio))
	return OptimisePortfolio(database)
}

// Retrieve Data
func (database *Database) GetRawPortfolio() []OwnPortfolioRecord {
//	fmt.Print("CHECKPOINT len of db being returned: ")
//	fmt.Print(len(database.ownstartingportfolio))
	return database.ownstartingportfolio
}

func (database *Database) GetCurrencyInputData() []CurrencyInputData {
	return database.currencyinputdata
}
