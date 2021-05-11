package db

/*
func getCurveData0() {

	// Connecting to client
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e009cbb4a2bd4c28a3174ac7884f4b42")
	if err != nil {
		log.Fatal(err)
	}

	// Creaitng a contract instance
	var curveRegistryAddress = common.HexToAddress("0x7D86446dDb609eD0F5f8684AcF30380a356b2B4c")
	provider, err := curveRegistry.NewMain(curveRegistryAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	var pools []CurvePoolData

	//pools = getCurveDataI(client, provider, pools, 3, true)
	fmt.Print("chk998")
	fmt.Println("Got day one")

	for i := 0; i < 31; i++ {

		fmt.Println("pool address:")
		fmt.Println(pools[i].poolAddress)
		fmt.Println("Fees collected:")
		for j := 0; j < 2; j++ {
			fmt.Println(pools[i].fees[j])
		}
		fmt.Println("Addresses of coins in the pool:")
		fmt.Println(pools[i].assetAddresses)
		fmt.Println("Decimals for coins in the pool:")
		fmt.Println(pools[i].assetDecimals)
		/*
			fmt.Println("Normalised volumes:")

			for j := 0; j < 8; j++{

				normalisedVolume := new(big.Float).SetInt(pools[i].volumes[j])
				normalisedVolume = negPow(normalisedVolume, pools[i].assetDecimals[j].Int64())

			}

	}

	zero := big.NewFloat(0)

	for i := 0; i < 4; i++ {
		var normalsied_fees []*big.Float
		var normalsied_balances []*big.Float
		current_coin_balances, err := provider.GetBalances(&bind.CallOpts{}, pools[i].poolAddress)

		if err != nil {
			log.Fatal(err)
		}

		for j := 0; j < 8; j++ {
			balance := new(big.Float).SetInt(current_coin_balances[j])
			balance = negPow(balance, pools[i].assetDecimals[j].Int64())
			normalsied_balances = append(normalsied_balances, balance)
		}

		for j := 0; j < 8; j++ {

			fee := negPow(pools[i].fees[0][j], pools[i].assetDecimals[j].Int64())
			normalsied_fees = append(normalsied_fees, fee)
		}

		pools[i].normalsiedBalances = normalsied_balances

		fmt.Println("Returns:")
		for j := 0; j < 8; j++ {
			if pools[i].normalsiedBalances[j].Cmp(zero) > 0 {
				returns := new(big.Float).Quo(normalsied_fees[j], pools[i].normalsiedBalances[j])
				fmt.Println(returns)
			}

		}

	}
}
*/
/*
func getCurveDataI(client *ethclient.Client, provider *curveRegistry.Main, pools []CurvePoolData, daysAgo int, first_turn bool) []CurvePoolData {
	fmt.Print("In CurvedataI")
	number_of_pools := big.NewInt(32)
	oldest_block := getOldestBlock(client, daysAgo)
	latest_block := getOldestBlock(client, daysAgo-1)
	count_pools := 0
	var one = big.NewInt(1)
	start := big.NewInt(1)
	end := big.NewInt(0).Sub(number_of_pools, big.NewInt(1))

	if first_turn {
		// Getting data from pools		// Getting data for the first pool
		pool_address, err := provider.PoolList(&bind.CallOpts{}, big.NewInt(0))

		if err != nil {
			log.Fatal(err)
		}

		// Addresses of underlying coins in the pool
		coin_addresses, err := provider.GetCoins(&bind.CallOpts{}, pool_address)
		if err != nil {
			log.Fatal(err)
		}

		// Getting the number of decimal spaces for undelying coins in the pool
		coin_decimals, err := provider.GetDecimals(&bind.CallOpts{}, pool_address)
		if err != nil {
			log.Fatal(err)
		}

		// Getting current pool balances

		// Getting swap volumes and fees and balances
		volumes, fees := curveGetPoolVolume(pool_address, client)
		var volumes_array []*[8]*big.Int
		volumes_array = append(volumes_array, volumes)
		var fees_array []*[8]*big.Float
		fees_array = append(fees_array, fees)

		// Appending a list of pool data structs
		pools = append(pools, CurvePoolData{poolAddress: pool_address, assetAddresses: coin_addresses,
			volumes: volumes_array, fees: fees_array, assetDecimals: coin_decimals})

		count_pools++
		/*
			fmt.Println("pool address:")
			fmt.Println(pool_address)
			fmt.Println("Fees collected:")
			fmt.Println(pools[count_pools].fees)
			fmt.Println("Addresses of coins in the pool:")
			fmt.Println(pools[count_pools].assetAddresses)
			fmt.Println("Decimals for coins in the pool:")
			fmt.Println(pools[count_pools].assetDecimals)


		// Getting data for the rest of the pools

		// i must be a new int so that it does not overwrite start
		for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, one) {

			pool_address, err = provider.PoolList(&bind.CallOpts{}, i)
			fmt.Println(pool_address)
			if err != nil {
				log.Fatal(err)
			}

			coin_addresses, err := provider.GetCoins(&bind.CallOpts{}, pool_address)
			if err != nil {
				log.Fatal(err)
			}

			// Get decimals for underlying tokens

			coin_decimals, err := provider.GetDecimals(&bind.CallOpts{}, pool_address)

			if err != nil {
				log.Fatal(err)
			}

			// Getting volumes and fees

			volumes, fees := curveGetPoolVolume(pool_address, oldest_block, latest_block, client)

			var volumes_array []*[8]*big.Int
			volumes_array = append(volumes_array, volumes)
			var fees_array []*[8]*big.Float
			fees_array = append(fees_array, fees)

			pools = append(pools, CurvePoolData{poolAddress: pool_address, assetAddresses: coin_addresses,
				volumes: volumes_array, fees: fees_array, assetDecimals: coin_decimals})

			count_pools++

		}

	} else {

		fmt.Println("Got into else")

		pool_address, err := provider.PoolList(&bind.CallOpts{}, big.NewInt(0))

		if err != nil {
			log.Fatal(err)
		}

		volumes, fees := curveGetPoolVolume(pool_address, oldest_block, latest_block, client)
		fmt.Println("Got first fees")
		pools[count_pools].volumes = append(pools[count_pools].volumes, volumes)
		pools[count_pools].fees = append(pools[count_pools].fees, fees)
		count_pools++
		fmt.Println("Upended first pool:")
		fmt.Println(pools[count_pools].volumes)

		for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, one) {
			fmt.Println("looping:")
			fmt.Println(i)
			pool_address, err = provider.PoolList(&bind.CallOpts{}, i)
			if err != nil {
				log.Fatal(err)
			}

			volumes, fees := curveGetPoolVolume(pool_address, client)
			pools[count_pools].volumes = append(pools[count_pools].volumes, volumes)
			pools[count_pools].fees = append(pools[count_pools].fees, fees)
			count_pools++

		}

	}

	return pools
}
*/
/*
func getOldestBlock(client *ethclient.Client, daysAgo int) *big.Int {

	var current_block *big.Int
	var oldest_block *big.Int
	current_block = big.NewInt(0)

	// Get current block
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	current_block = header.Number

	//2)  Find oldest block in our lookup date range
	oldest_block = new(big.Int).Set(current_block)

	now := time.Now()

	//timeonehourago := uint64(now.Add(-2*time.Hour).Unix())
	//timeonemonthago := uint64((now.AddDate(0, 0, -1)).Unix())
	timeonemonthago := uint64(now.Unix()) - 24*60*60*uint64(daysAgo)

	var j int64
	j = 0

	for {
		j -= 500
		oldest_block.Add(oldest_block, big.NewInt(j))

		block, err := client.BlockByNumber(context.Background(), oldest_block)
		if err != nil {
			log.Fatal(err)
		}

		if block.Time() < timeonemonthago {

			break
		}
	}

	return oldest_block
}
*/

//Function to convert a big Float to a BigInt
/*
func FloatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	// Set precision if required.
	// bigval.SetPrec(64)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}
*/

/*

/*
	current_coin_balances, err := provider.GetUnderlyingBalances(pool_address)

	if err != nil {
		log.Fatal(err)
	}

	// Add decimal spaces to volumes and fees and balances

	for i := 0; i < 8; i++{
		normvolumes[i] = negPow(volumes[i], coin_decimals.Int64())
		normfees[i] = negPow(fees[i], coin_decimals.Int64())
		current_coin_balances[i] = negPow(current_coin_balances[i], coin_decimals.Int64())
	}

	// Calculate returns

	var returns []*big.Float

	for i := 0; i < 8; i++{
		returns = append(returns, Quo(normfees[i], current_coin_balances[i]) )
	}
*/

/*
	/*
	fmt.Println("Returns:")
	fmt.Println(pools[count_pools].returns)

	for i := 0; i < 8; i++{

		normalisedVolume := new(big.Float).SetInt(pools[count_pools].volumes[i])

		normalisedVolume = negPow(normalisedVolume, pools[count_pools].assetDecimals[i].Int64())

		fmt.Println("Orginal volume")
		fmt.Println(pools[count_pools].volumes[i])
		fmt.Println("Normalised volume")
		fmt.Println(normalisedVolume)



	}

*/
