package claude

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"maxpothier.com/go/api/v2/model"
)

var baseUrl = "https://api.anthropic.com/v1/messages"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}

func GetInfoBreakdown(stock *model.StockData) (string, error) {
	apiKey := os.Getenv("CLAUDE_API_KEY")

	if apiKey == "" {
		return "", errors.New("CLAUDE_API_KEY is not set")
	}

	requestBody := map[string]interface{}{
		"model": 		"claude-3-5-sonnet-20240620",
		"max_tokens": 	1024,
		"messages": 	[]map[string]string{
			{
				"role": "user",
				"content": fmt.Sprintf("Tell me about %s in two sentences. Here is the raw data from the stock market API for this stock:\n\n%s", stock.Ticker, stock.Data),
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}

	req, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	var response struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	if len(response.Content) > 0 {
		return response.Content[0].Text, nil
	} else {
		return "", errors.New("no content in the response")
	}
}