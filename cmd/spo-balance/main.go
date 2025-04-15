package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/blockfrost/blockfrost-go"
	"github.com/davidt58/spo-balance/configs"
	"github.com/davidt58/spo-balance/internal/database"
	"github.com/davidt58/spo-balance/internal/lbank"
	"github.com/davidt58/spo-balance/internal/models"
)

func main() {

	var (
		sum    float64 = 0
		config configs.Config
		err    error
	)

	configFile := flag.String("config", "example.config.yaml", "Path to yaml config file")
	blockFrostServer := flag.String("blockfrost-address", "http://185.170.115.3:3033", "Blockfrost API address")
	flag.Parse()

	config, err = configs.LoadConfigFromYAML(*configFile)
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}

	blockfrostClient := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{
			Server: *blockFrostServer,
		},
	)

	err = database.Initialize("./data.db")

	// var pHistory []blockfrost.PoolHistory

	for _, pool := range config.Pools {
		// pHistory, err := blockfrostClient.PoolHistory(context.TODO(), poolID, blockfrost.APIQueryParams{})
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// for _, ph := range pHistory {
		// 	parsedRewards, _ := strconv.ParseInt(ph.Rewards, 10, 64)
		// 	fmt.Printf("Epoch: %d\n\tRewards: %d\n", ph.Epoch, parsedRewards/1000000)
		// }
		poolInfo, err := blockfrostClient.Pool(context.TODO(), pool.PoolID)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(poolInfo.LivePledge)
		parsedAmount, _ := strconv.ParseInt(poolInfo.LivePledge, 10, 64)
		sum += float64(parsedAmount)
	}

	sum = sum / 1000000

	fmt.Printf("Total Pledge: %f\n", sum)

	lbankClient := lbank.NewClient()

	lastPrice, err := database.GetLastPrice()

	if err != nil {
		log.Fatal("Failed to get last price:", err)
	}

	if time.Since(lastPrice.CreatedAt) > 2*time.Hour {
		price, err := lbankClient.GetPrice("ap3x_usdt")
		if err != nil {
			log.Fatal("Failed to fetch price:", err)
		}

		if err != nil {
			log.Fatal("Failed to parse price:", err)
		}

		lastPrice, err = database.StorePrice(models.Price{Price: price.Data[0].Price})

		if err != nil {
			log.Fatal("Failed to store price in database:", err)
			panic(err)
		}
	} else {
		fmt.Println("Using cached price")
	}

	convertedPrice, _ := strconv.ParseFloat(lastPrice.Price, 64)

	fmt.Printf("Price: %f\n", convertedPrice)
	fmt.Printf("Total APEX: %f\n", sum)

	TOTAL := int(sum * convertedPrice)

	fmt.Printf("Total Value: $%d\n", TOTAL)

	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// fmt.Printf("Address: %s\n", *address)
	// fmt.Printf("Balance: %s\n", balance.String())
}
