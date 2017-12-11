package gobittrex

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

const (
	ApiURL     = "https://bittrex.com/api/"
	APIVersion = "v1.1"
)

// functionType 0: public; 1: market; 2: account
var functionType = map[string]uint8{
	"getmarkets":         0,
	"getcurrencies":      0,
	"getticker":          0,
	"getmarketsummaries": 0,
	"getmarketsummary":   0,
	"getorderbook":       0,
	"getmarkethistory":   0,
}

//const apiKey = GET FROM CONFIG FILE

func apiQuery(method string, params map[string]string) (apiResp APIResponse, err error) {
	var queryURL = generateQueryURL(method, params)

	response, err := http.Get(queryURL)
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &apiResp)
	if err = handleAPIError(apiResp); err != nil {
		return
	}
	return
}

func generateQueryURL(method string, params map[string]string) (queryURL string) {
	queryURL = ApiURL + APIVersion
	switch functionType[method] {
	case 0:
		queryURL += "/public/"
	case 1:
		queryURL += "/market/"
	case 2:
		queryURL += "/account/"
	}
	queryURL += method

	if len(params) > 0 {
		queryURL += "?"

		for key, value := range params {
			queryURL += key + "=" + value + "&"
		}

		queryURL = queryURL[:len(queryURL)-1]
	}

	fmt.Println("Calling: ", queryURL)

	return
}

func handleAPIError(apiResp APIResponse) error {
	if !apiResp.Success {
		return errors.New(apiResp.Message)
	}
	return nil
}

// *** Public APIs ***
// Following functions are callable without authentication

func GetMarkets() (markets []Market, err error) {
	apiResp, err := apiQuery("getmarkets", map[string]string{})
	err = json.Unmarshal(apiResp.Result, &markets)
	return
}

func GetCurrencies() (currencies []Currency, err error) {
	apiResp, err := apiQuery("getcurrencies", map[string]string{})
	err = json.Unmarshal(apiResp.Result, &currencies)
	return
}

func GetTicker(market string) (ticker []Ticker, err error) {
	apiResp, err := apiQuery("getticker", map[string]string{"market": market})
	err = json.Unmarshal(apiResp.Result, &ticker)
	return
}

func GetMarketSummaries() (marketSummaries []MarketSummary, err error) {
	apiResp, err := apiQuery("getmarketsummaries", map[string]string{})
	err = json.Unmarshal(apiResp.Result, &marketSummaries)
	return
}

func GetMarketSummary(market string) (marketSummary MarketSummary, err error) {
	apiResp, err := apiQuery("getmarketsummary", map[string]string{"market": market})
	err = json.Unmarshal(apiResp.Result, &marketSummary)
	return
}

func GetOrderBook(market string, otype string) (orderBook OrderBook, err error) {
	apiResp, err := apiQuery("getorderbook", map[string]string{"market": market, "type": otype})
	err = json.Unmarshal(apiResp.Result, &orderBook)
	return
}

func GetMarketHistory(market string) (marketHistory []MarketHistoryEvent, err error) {
	apiResp, err := apiQuery("getmarkethistory", map[string]string{"market": market})
	err = json.Unmarshal(apiResp.Result, &marketHistory)
	return
}

// *** Market APIs ***
// Following functions are only callable with authentication

// *** Account APIs ***
// Following functions are only callable with authentication
