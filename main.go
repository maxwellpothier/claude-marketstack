package main

import (
	"fmt"
	"log"

	"maxpothier.com/go/api/v2/claude"
	"maxpothier.com/go/api/v2/marketapi"
)

var ticker = "AAPL"

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("Error loading .env file")
// 	}
// }

func main() {
	stockResponse, err := marketapi.GetStockData(ticker)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(stockResponse.Ticker)

	// test := claude.GetInfoBreakdown(stockResponse)
	claude.GetInfoBreakdown(stockResponse)
}