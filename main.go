package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"maxpothier.com/go/api/v2/marketapi"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	apiUrl := marketapi.GetAPIUrl()
	res, err := http.Get(fmt.Sprintf(apiUrl, "AAPL"))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the entire response body as a string
	fmt.Println(string(body))
}