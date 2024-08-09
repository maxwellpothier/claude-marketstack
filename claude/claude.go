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

var prompt = `
You are tasked with analyzing and presenting stock market data in a clear and informative manner. Your goal is to provide a comprehensive yet digestible overview of the stock's performance and key metrics.

The stock symbol you will be analyzing is:
<stock_symbol>%s</stock_symbol>

Here is the raw data from the stock market API for this stock:
<stock_data>
%s
</stock_data>

Please follow these instructions to analyze and present the data:

1. Carefully parse through all the provided data, identifying key metrics and information about the stock.

2. Focus on the following areas, if the information is available:
   - Current price and recent price changes
   - Trading volume
   - Market capitalization
   - Price-to-earnings (P/E) ratio
   - Dividend yield
   - 52-week high and low
   - Any notable technical indicators (e.g., moving averages, RSI)
   - Recent news or events affecting the stock

3. Organize the information into a clear and logical structure. Start with the most important or noteworthy facts about the stock.

4. Provide context for the numbers you present. For example, if discussing the P/E ratio, briefly explain what it means and how it compares to industry averages.

5. If there are any significant trends or patterns in the data, highlight these and provide a brief explanation of their potential implications.

6. Include a brief summary of the stock's recent performance and any factors that may be influencing it.

7. If possible, provide a short-term outlook based on the available data, but avoid making specific predictions about future stock prices.

8. Use clear, concise language that would be understandable to an investor with basic knowledge of the stock market.

9. Present your analysis in a well-structured format, using bullet points or short paragraphs for easy readability.

10. Conclude with a brief overall assessment of the stock's current status and key points for potential investors to consider.

Begin with a brief introduction, then present the detailed analysis, and end with a concise conclusion. Use appropriate subheadings to organize the information.

Remember, your goal is to provide a comprehensive yet easily digestible overview of the stock's current status and performance based solely on the provided data. Do not include any external information or personal opinions not supported by the given data.
`

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
				"content": fmt.Sprintf(prompt, stock.Ticker, stock.Data),
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