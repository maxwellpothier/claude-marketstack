package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"maxpothier.com/go/api/v2/claude"
	"maxpothier.com/go/api/v2/marketapi"
)

func getAnalysis(w http.ResponseWriter, r *http.Request) {
	stockTicker := r.URL.Query().Get("ticker")
	upperTicker := strings.ToUpper(stockTicker)
	stockResponse, err := marketapi.GetStockData(upperTicker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	claudeResponse, err := claude.GetInfoBreakdown(stockResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(claudeResponse)

}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/analysis", getAnalysis)

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}