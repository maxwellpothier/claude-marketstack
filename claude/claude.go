package claude

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"maxpothier.com/go/api/v2/model"
)

var apiUrl string
var baseUrl = "https://api.anthropic.com/v1/messages"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}

func init() {
	apiKey := os.Getenv("CLAUDE_API_KEY")
	if apiKey == "" {
		panic("CLAUDE_API_KEY is not set")
	}
	apiUrl = fmt.Sprintf("http://api.marketstack.com/v1/eod?access_key=%s&symbols=%%s", apiKey)
}

func GetInfoBreakdown(stock *model.StockData) {
	apiKey := os.Getenv("CLAUDE_API_KEY")

	if apiKey == "" {
		panic("CLAUDE_API_KEY is not set")
	}

	requestBody := map[string]interface{}{
		"model": 		"claude-3-5-sonnet-20240620",
		"max_tokens": 	1024,
		"messages": 	[]map[string]string{
			{
				"role": "user",
				"content": fmt.Sprintf("Tell me about %s in two sentences", stock.Ticker),
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		panic(fmt.Errorf("error marshaling JSON: %v", err))
	}

	req, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(fmt.Errorf("error creating request: %v", err))
	}

	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Errorf("error sending request: %v", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("error reading response: %v", err))
	}

	var response struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(fmt.Errorf("error unmarshaling JSON: %v", err))
	}

	if len(response.Content) > 0 {
		fmt.Println(response.Content[0].Text)
	} else {
		fmt.Println("No content in the response")
	}
}