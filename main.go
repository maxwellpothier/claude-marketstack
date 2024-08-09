package main

import (
	"fmt"
	"log"

	"maxpothier.com/go/api/v2/claude"
	"maxpothier.com/go/api/v2/marketapi"
)

var ticker = "AAPL"

func main() {
	stockResponse, err := marketapi.GetStockData(ticker)
	if err != nil {
		log.Println(err)
	}

	claudeResponse, err := claude.GetInfoBreakdown(stockResponse)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(claudeResponse)
}