package main

import (
	"strconv"

	"spo-data/internal/database"
	"spo-data/internal/models"
)

func main() {

	database.Initialize("test.db")

	for i := range 10 {
		price := models.Price{Price: strconv.FormatInt(int64(i), 10)}
		database.StorePrice(price)
	}

	// price := models.Price{Price: "1000"}

	// database.StorePrice(price)

	priceRet, err := database.GetLastPrice()
	if err != nil {
		panic(err)
	}
	println(priceRet.Price)
}
