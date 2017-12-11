package gobittrex

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"github.com/shopspring/decimal"
	"time"
	"crypto/hmac"
	"crypto/sha512"
)

const (
	ApiURL     = "https://bittrex.com/api/"
	APIVersion = "v1.1"
)

// functionType 0: public; 1: market; 2: account
var functionType = map[string]uint8{
	"getmarkets":           0,
	"getcurrencies":        0,
	"getticker":            0,
	"getmarketsummaries":   0,
	"getmarketsummary":     0,
	"getorderbook":         0,
	"getmarkethistory":     0,
	"BuyLimit":             1,
	"SellLimit":            1,
	"Cancel":               1,
	"GetOpenOrders":        1,
	"GetBalances":          2,
	"GetBalance":           2,
	"GetDepositAddress":    2,
	"Withdraw":             2,
	"GetOrder":             2,
	"GetOrderHistory":      2,
	"GetWithdrawalHistory": 2,
	"GetDepositHistory":    2,
}

var ApiKey string
var ApiSecret string
var nonce time.Time

func apiQuery(method string, params *map[string]string) (apiResp APIResponse, err error) {
	var queryURL = generateQueryURL(&method, params)

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
	if !apiResp.Success {
		err = errors.New(apiResp.Message)
	}
	return
}

func generateQueryURL(method *string, params *map[string]string) (queryURL string) {
	queryURL = ApiURL + APIVersion
	var authNeeded bool
	switch functionType[*method] {
	case 0:
		queryURL += "/public/"
	case 1:
		queryURL += "/market/"
		authNeeded = true
	case 2:
		queryURL += "/account/"
		authNeeded = true
	}
	queryURL += *method

	if params != nil && len(*params) > 0 {
		queryURL += "?"
		for key, value := range *params {
			queryURL += key + "=" + value + "&"
		}
	}

	if authNeeded {
		// TODO: Authentication
		nonce := time.Now().UnixNano()
		queryURL += "apikey=" + ApiKey + "&nonce=" + string(nonce)
		sign := hmac.New(sha512.New, []byte(ApiSecret))
		_, err := sign.Write([]byte(queryURL))
		if err != nil {
			return // TODO: Add Error handling
		}
	} else {
		queryURL = queryURL[:len(queryURL)-1]
	}

	fmt.Println("Calling: ", queryURL)
	return
}

// *** Public APIs ***
// Following functions are callable without authentication

func GetMarkets() (markets []Market, err error) {
	apiResp, err := apiQuery("getmarkets", nil)
	err = json.Unmarshal(apiResp.Result, &markets)
	return
}

func GetCurrencies() (currencies []Currency, err error) {
	apiResp, err := apiQuery("getcurrencies", nil)
	err = json.Unmarshal(apiResp.Result, &currencies)
	return
}

func GetTicker(market string) (ticker Ticker, err error) {
	apiResp, err := apiQuery("getticker", &map[string]string{"market": market})
	err = json.Unmarshal(apiResp.Result, &ticker)
	return
}

func GetMarketSummaries() (marketSummaries []MarketSummary, err error) {
	apiResp, err := apiQuery("getmarketsummaries", nil)
	err = json.Unmarshal(apiResp.Result, &marketSummaries)
	return
}

func GetMarketSummary(market string) (marketSummary MarketSummary, err error) {
	apiResp, err := apiQuery("getmarketsummary", &map[string]string{"market": market})
	err = json.Unmarshal(apiResp.Result, &marketSummary)
	return
}

func GetOrderBook(market string, otype string) (orderBook OrderBook, err error) {
	apiResp, err := apiQuery("getorderbook", &map[string]string{"market": market, "type": otype})
	err = json.Unmarshal(apiResp.Result, &orderBook)
	return
}

func GetMarketHistory(market string) (marketHistory []MarketHistoryEvent, err error) {
	apiResp, err := apiQuery("getmarkethistory", &map[string]string{"market": market})
	err = json.Unmarshal(apiResp.Result, &marketHistory)
	return
}

// *** Market APIs ***
// Following functions are only callable with authentication

func BuyLimit() {}

func SellLimit() {}

func Cancel() {}

func GetOpenOrders() {}

// *** Account APIs ***
// Following functions are only callable with authentication

func GetBalances() {}

func GetBalance(currency string) {}

func GetDepositAddress(currency string) {}

func Withdraw(currency string, quantity decimal.Decimal, address string) {}

func GetOrder(uuid string) {}

func GetOrderHistory(market string) {}

func GetWithdrawalHistory(currency string) {}

func GetDepositHistory(curreny string) {}
