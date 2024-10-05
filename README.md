# Stock Market Data Analysis API

This project is a Go-based API that provides stock market data analysis. It fetches stock data from a market API and uses Claude AI to analyze and present the information in a clear, digestible format.

## What the App Does

1. Fetches stock market data for a given stock symbol using the Marketstack API.
2. Processes the raw data through Claude AI to generate a comprehensive analysis.
3. Presents the analysis in a structured, easy-to-understand format.

## How to Use

### Prerequisites

-   Go 1.19 or later
-   Claude API key
-   Marketstack API key

### Setup

1. Clone the repository
2. Create a `.env` file in the root directory with the following content:

```
CLAUDE_API_KEY=your_claude_api_key
MARKETSTACK_API_KEY=your_marketstack_api_key
```

3. Install dependencies:
    - `go mod download`

### Running the Application

1. Start the server:

    - `go run main.go`

2. The server will start on `localhost:8080`.

3. To get an analysis for a stock, make a GET request to the `/analysis` endpoint with the `ticker` query parameter:
    - http://localhost:8080/analysis?ticker=AAPL

Replace `AAPL` with the desired stock symbol.

4. The API will return a JSON response containing the AI-generated analysis of the stock's performance and key metrics.

## API Endpoints

-   `GET /analysis?ticker={symbol}`: Retrieves a comprehensive analysis of the specified stock symbol.

## Project Structure

-   `main.go`: Entry point of the application, sets up the HTTP server and routes.
-   `claude/claude.go`: Handles communication with the Claude AI API for stock analysis.
-   `marketapi/marketstack.go`: Manages requests to the Marketstack API for raw stock data.
-   `model/data.go`: Defines the data structures used in the application.

## Note

This application uses external APIs (Marketstack and Claude) which may have usage limits or costs associated with them. Please refer to their respective documentation for more information.
