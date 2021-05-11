package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func convaavetoken(aavetoken string) string {
	assetName := " "
	if aavetoken == "Eth" {
		assetName = "ETH"
	} else if aavetoken == "Republic Token" {
		assetName = "REN"
	} else if aavetoken == "Synthetix Network Token" {
		assetName = "SNX"
	} else if aavetoken == "yearn.finance" {
		assetName = "YFI"
	} else if aavetoken == "Wrapped BTC" {
		assetName = "WBTC"
	} else if aavetoken == "Wrapped Ether" {
		assetName = "WETH"
	} else if aavetoken == "Uniswap" {
		assetName = "UNI"
	} else if aavetoken == "USD Coin" {
		assetName = "USDC"
	} else if aavetoken == "True USD" {
		assetName = "USDT"
	}
	return assetName

}

func getUsdFromVolumeAave1(aavePoolData AavePoolData) {
	fmt.Print("Asset name BEFORE: ")
	fmt.Print(aavePoolData.assetName)

	assetName := convaavetoken(aavePoolData.assetName)

	fmt.Print("Asset name: ")
	fmt.Print(assetName)

	if !isPoolPartOfFilter(assetName, assetName) {
		return
	}

	histcombo := retrieveDataForTokensFromDatabase2(assetName, "USD")
	exchangeRate := histcombo.Price
	if len(exchangeRate) == 0 {
		fmt.Print("No data!!!! 647")
	}

	exch_fixed := 0
	coinPegged, ticker := isCoinPeggedToUSD(assetName)
	if coinPegged {
		exch_fixed = 1.0
		assetName = ticker
		if exch_fixed == 0 {
		}
	}

	for i := 0; i < len(aavePoolData.volumes); i++ {
		ex := float32(0.0)
		fmt.Print("i: ")
		fmt.Print(i)
		fmt.Print("Len exch: ")
		fmt.Print(len(exchangeRate))
		fmt.Print("Len aave volumes: ")
		fmt.Print(len(aavePoolData.volumes))
		fmt.Print(" | ")
		fmt.Print("vi: ")
		fmt.Print(aavePoolData.volumes[i])

		if exch_fixed != 0 {
			ex = 1
		} else {
			if len(exchangeRate) > i {
				ex = exchangeRate[i]
			} else {
				ex = 0.0
			}

		}

		volumeUSD := float32(aavePoolData.volumes[i].Int64()) * ex // aavePoolData.volumes[i] * exchangeRate
		// aavePoolData.currentBalance to be added to AavePoolData struct by Ivan

		// check if that day is already in db

		//		data_is_old := false
		//		oldest_available_record := time.Unix(get_newest_timestamp_from_db_hist_volume_and_sz("Aave1", assetName, assetName), 0)

		//		if (time.Since(oldest_available_record).Hours()) > 24 {
		//			data_is_old = true
		//		}

		//		if data_is_old { // if no -- add
		//		addAave1PoolDataToAave1Database(assetName, float32(volumeUSD),
		//			0.0, BoD(time.Now()).Unix()-int64(86400*30+i*86400))
		//		} else {
		//fmt.Print("Data already appended")
		//		}

		// Add pool size
		// Add fees
		days_ago := 30
		pool_sz_usd := int64(0.0)
		var tokenqueue []string
		tokenqueue = append(tokenqueue, assetName)
		tokenqueue = append(tokenqueue, assetName)
		interest := float64(0.0)
		xx := append_record_to_database("Aave1", tokenqueue, BoD(time.Now()).Unix()-int64(86400*days_ago)+int64(i*86400), int64(volumeUSD), pool_sz_usd, 0, interest)
		if xx == " " {
		}
	}

}

func getUsdFromVolumeAave2(aavePoolData AavePoolData) {

	assetName := convaavetoken(aavePoolData.assetName)
	fmt.Print(assetName)

	histcombo := retrieveDataForTokensFromDatabase2(assetName,
		"USD")
	exchangeRate := histcombo.Price
	if len(exchangeRate) == 0 {
	}

	coinPegged, ticker := isCoinPeggedToUSD(assetName)
	if coinPegged {
		exchangeRate := 1.0
		assetName = ticker
		if exchangeRate == 0 {
		}

	}

	for i := 0; i < len(aavePoolData.volumes); i++ {
		//usdVolume = append(usdVolume, aavePoolData.volumes[i] * exchangeRate)
		volumeUSD := 0 // aavePoolData.volumes[i] * exchangeRate

		// aavePoolData.currentBalance to be added to AavePoolData struct by Ivan
		addAave2PoolDataToAave2Database(assetName, float32(volumeUSD),
			0.0, histcombo.Date[i]) // aavePoolData.currentBalance
	}
}

func isTokenStableCoin(coinName string) bool {
	if coinName == "USDT" {
		return true
	} else if coinName == "USDC" {
		return true
	} else if coinName == "USD" {
		return true
	} else if coinName == "TUSD" {
		return true
	} else if coinName == "DAI" {
		return true
	} else if coinName == "GUSD" {
		return true
	} else if coinName == "BUSD" {
		return true
	} else {
		return false
	}
}

func isCoinPeggedToUSD(coinName string) (bool, string) {
	if coinName == "Tether USD" {
		return true, "USDT"
	} else if coinName == "Binance USD" {
		return true, "BUSD"
	} else if coinName == "Synth sUSD" {
		return true, "SUSD"
	} else if coinName == "TrueUSD" {
		return true, "TUSD"
	} else if coinName == "USD Coin" {
		return true, "USDC"
	} else if coinName == "Dai Stablecoin" {
		return true, "DAI"
	} else if coinName == "Gemini dollar" {
		return true, "GUSD"
	} else {
		return false, "nil"
	}
}

// database add function
func addAave1PoolDataToAave1Database(assetName string, volume float32,
	currentBalance float32, timestamp int64) string {
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
	aave1data := Database.Collection("Aave1 Pool Data")

	new_aave1data, err := aave1data.InsertOne(ctx, bson.D{
		{Key: "AssetName", Value: assetName},
		{Key: "Volume", Value: volume},
		{Key: "CurrentBalance", Value: currentBalance},
		{Key: "Timestamp", Value: timestamp},
	})

	if err != nil {
		log.Fatal(err)
	}

	newID := new_aave1data.InsertedID
	hexID := newID.(primitive.ObjectID).Hex()

	return hexID

}

func addAave2PoolDataToAave2Database(assetName string, volume float32,
	currentBalance float32, timestamp int64) string {
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
	aave2data := Database.Collection("Aave2 Pool Data")

	new_aave2data, err := aave2data.InsertOne(ctx, bson.D{
		{Key: "AssetName", Value: assetName},
		{Key: "Volume", Value: volume},
		{Key: "CurrentBalance", Value: currentBalance},
		{Key: "Timestamp", Value: timestamp},
	})

	if err != nil {
		log.Fatal(err)
	}

	newID := new_aave2data.InsertedID
	hexID := newID.(primitive.ObjectID).Hex()

	return hexID

}
