package balance

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"spo-data/configs"
	"spo-data/internal/database"
	"spo-data/internal/lbank"
	"spo-data/internal/models"

	"github.com/blockfrost/blockfrost-go"
)

func CalculateBalance(config *configs.Config, blockfrostClient blockfrost.APIClient) (float64, error) {

	var sum float64 = 0

	for _, pool := range config.Pools {
		poolInfo, err := blockfrostClient.Pool(context.TODO(), pool.PoolID)
		if err != nil {
			fmt.Println(err)
		}
		parsedAmount, _ := strconv.ParseInt(poolInfo.LivePledge, 10, 64)
		fmt.Printf("Pool name: %s, pledge %d\n", pool.Name, parsedAmount/1000000)
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

	return convertedPrice * sum, nil
}
