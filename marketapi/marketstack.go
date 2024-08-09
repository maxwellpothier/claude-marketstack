package marketapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"maxpothier.com/go/api/v2/model"
)

var apiUrl string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}

func init() {
	apiKey := os.Getenv("MARKETSTACK_API_KEY")
	if apiKey == "" {
		panic("MARKETSTACK_API_KEY is not set")
	}
	apiUrl = fmt.Sprintf("http://api.marketstack.com/v1/eod?access_key=%s&symbols=%%s", apiKey)

}
func GetStockData(symbol string) (*model.StockData, error) {
	upSymbol := strings.ToUpper(symbol)
	res, err := http.Get(fmt.Sprintf(apiUrl, upSymbol))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	stock := model.StockData{Ticker: symbol, Data: string(body)}

	return &stock, nil
}