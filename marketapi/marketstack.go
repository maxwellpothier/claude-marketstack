package marketapi

import (
	"fmt"
	"os"
)

var apiUrl string

func init() {
	apiKey := os.Getenv("MARKETSTACK_API_KEY")
	if apiKey == "" {
		panic("MARKETSTACK_API_KEY is not set")
	}
	apiUrl = fmt.Sprintf("http://api.marketstack.com/v1/eod?access_key=%s&symbols=%%s", apiKey)

}

func GetAPIUrl() string {
	return apiUrl
}
// const apiUrl = "http://api.marketstack.com/v1/eod?access_key=0de8a225b80862a09ff3bc202c17c61a&symbols=%s"