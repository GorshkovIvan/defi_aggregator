package db
import (

	"log"
	"math/big"
	"context"
	"time"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)
type AavePoolData struct {

	assetAddress string 
	assetName string
	interest_rates []*big.Int 
	volumes []*big.Int 
	rate_types []int
	flashLoanVolumes []*big.Int
	flashLoanFees []*big.Int

}


func getUsdFromVolumeAave1(aavePoolData AavePoolData) {
	assetName := aavePoolData.assetName

	if (aavePoolData.assetName == "Eth") {
		assetName = "ETH";
	} else if (aavePoolData.assetName == "Republic Token") {
		assetName = "REN";
	} else if (aavePoolData.assetName == "Synthetix Network Token") {
		assetName = "SNX"; 
	} else if (aavePoolData.assetName == "yearn.finance") {
		assetName = "YFI";
	} else if (aavePoolData.assetName == "Wrapped BTC") {
		assetName = "WBTC";
	} else if (aavePoolData.assetName == "Wrapped Ether") {
		assetName = "WETH";
	} else if (aavePoolData.assetName == "Uniswap") {
		assetName = "UNI";
	}
	
	histcombo := retrieveDataForTokensFromDatabase2(assetName,
		"USD")
	exchangeRate := histcombo.Price

	coinPegged, ticker := isCoinPeggedToUSD(assetName)
	if (coinPegged) {
		exchangeRate := 1.0
		assetName = ticker
	}
	
	for i := 0; i < len(aavePoolData.volumes); i++ {
		//usdVolume = append(usdVolume, aavePoolData.volumes[i] * exchangeRate)
		volumeUSD := aavePoolData.volumes[i] * exchangeRate

		// aavePoolData.currentBalance to be added to AavePoolData struct by Ivan
		addAave1PoolDataToAave1Database(assetName, volumeUSD, 
			aavePoolData.currentBalance, histcombo.Date)	
	}
	
}

func getUsdFromVolumeAave2(aavePoolData AavePoolData) {
	assetName := aavePoolData.assetName

	if (aavePoolData.assetName == "Eth") {
		assetName = "ETH";
	} else if (aavePoolData.assetName == "Republic Token") {
		assetName = "REN";
	} else if (aavePoolData.assetName == "Synthetix Network Token") {
		assetName = "SNX"; 
	} else if (aavePoolData.assetName == "yearn.finance") {
		assetName = "YFI";
	} else if (aavePoolData.assetName == "Wrapped BTC") {
		assetName = "WBTC";
	} else if (aavePoolData.assetName == "Wrapped Ether") {
		assetName = "WETH";
	} else if (aavePoolData.assetName == "Uniswap") {
		assetName = "UNI";
	}
	
	histcombo := retrieveDataForTokensFromDatabase2(assetName,
		"USD")
	exchangeRate := histcombo.Price

	coinPegged, ticker := isCoinPeggedToUSD(assetName)
	if (coinPegged) {
		exchangeRate := 1.0
		assetName = ticker
	}
	
	for i := 0; i < len(aavePoolData.volumes); i++ {
		//usdVolume = append(usdVolume, aavePoolData.volumes[i] * exchangeRate)
		volumeUSD := aavePoolData.volumes[i] * exchangeRate

		// aavePoolData.currentBalance to be added to AavePoolData struct by Ivan
		addAave2PoolDataToAave2Database(assetName, volumeUSD, 
			aavePoolData.currentBalance, histcombo.Date)	
	}
}

func isCoinPeggedToUSD(coinName string) (bool, string) {
	if (coinName == "Tether USD") {
		return true, "USDT"
	} else if (coinName == "Binance USD") {
		return true, "BUSD"
	} else if (coinName == "Synth sUSD") {
		return true, "SUSD"
	} else if (coinName == "TrueUSD") {
		return true, "TUSD"
	} else if (coinName == "USD Coin") {
		return true, "USDC"
	} else if (coinName == "Dai Stablecoin") {
		return true, "DAI"
	} else if (coinName == "Gemini dollar") {
		return true, "GUSD"
	} else {
		return false, nil
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
