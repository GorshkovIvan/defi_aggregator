package db

/*
func convCurveToken(token string) string {
	assetName := " "

	if token == "Eth" {
		assetName = "ETH"
	} else if token == "Republic Token" {
		assetName = "REN"
	} else if token == "Synthetix Network Token" {
		assetName = "SNX"
	} else if token == "yearn.finance" {
		assetName = "YFI"
	} else if token == "Wrapped BTC" {
		assetName = "WBTC"
	} else if token == "Wrapped Ether" {
		assetName = "WETH"
	} else if token == "Uniswap" {
		assetName = "UNI"
	}
	return assetName
}
*/

/*
func getUsdFromVolumeCurve(curvePoolData CurvePoolData) {
	// more tokens

	// 0) Connect to client
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	instance0, err := token.NewToken(curvePoolData.assetAddresses[0], client)
	if err != nil {
		log.Fatal(err)
	}

	symbol0, err := instance0.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	instance1, err := token.NewToken(curvePoolData.assetAddresses[0], client)
	if err != nil {
		log.Fatal(err)
	}

	symbol1, err := instance1.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	assetName0 := convCurveToken(symbol0)
	assetName1 := convCurveToken(symbol1)

	if !isPoolPartOfFilter(assetName0, assetName1) {
		return
	}

	histcombo := retrieveDataForTokensFromDatabase2(assetName0, assetName1)
	//exchangeRate := histcombo.Price
	// ex := 0
	/*
		coinPegged0, ticker0 := isCoinPeggedToUSD(assetName0)
		if coinPegged {
			ex = 1.0
			assetName0 = ticker0
		}

		coinPegged, ticker1 = isCoinPeggedToUSD(assetName1)
		if coinPegged {
			ex = 1.0
			assetName1 = ticker
		}

	for i := 0; i < len(curvePoolData.volumes); i++ {
		fmt.Print(curvePoolData.volumes[i][0])
		v := 0 // float32((curvePoolData.volumes[i][0]).Int64())]
		fmt.Print(v)
		volumeUSD := float32(0.0) * histcombo.Price[i]
		balance := float32(0.0)
		fmt.Print(balance)
		pool_sz := int64(0)
		days_ago := 1
		t := BoD(time.Now()).Unix() - int64(86400*days_ago) + int64(i*86400)
		xx := append_hist_volume_record_to_database("Curve", assetName0, assetName1, t, int64(volumeUSD), pool_sz)
		if xx == " " {
		}
	}

	//	var Date []int64    // `default0:"10099999999999" json:"date"` // `default0:"Mon Jan 2 15:04:05 MST 2006" json:"date"`
	//	var Price []float32 // `default0:"420.69" json:"price"`
}
*/

// database add function
/*
func addCurvePoolDataToCurveDatabase(assetName string, volume float32,
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
	curvedata := Database.Collection("Curve Pool Data")

	new_curvedata, err := curvedata.InsertOne(ctx, bson.D{
		{Key: "AssetName", Value: assetName},
		{Key: "Volume", Value: volume},
		{Key: "CurrentBalance", Value: currentBalance},
		{Key: "Timestamp", Value: timestamp},
	})

	if err != nil {
		log.Fatal(err)
	}

	newID := new_curvedata.InsertedID
	hexID := newID.(primitive.ObjectID).Hex()

	return hexID

}
*/
