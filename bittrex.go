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
	"strings"
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

var client = http.Client{
	Timeout: time.Second * 30,
}

func apiQuery(function string, httpMethod string, params *map[string]string) (apiResp APIResponse, err error) {
	req, err := generateRequest(&function, &httpMethod, params)
	if err != nil {
		return
	}

	fmt.Println("API Call: ", req.URL.String())
	res, err := client.Do(req)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &apiResp)
	if !apiResp.Success {
		err = errors.New(apiResp.Message)
	}
	return
}

func generateRequest(function *string, httpMethod *string, params *map[string]string) (client *http.Request, err error) {
	reqURL := ApiURL + APIVersion
	var authNeeded bool

	switch functionType[*function] {
	case 0:
		reqURL += "/public/"
	case 1:
		reqURL += "/market/"
		authNeeded = true
	case 2:
		reqURL += "/account/"
		authNeeded = true
	}
	reqURL += *function

	req, err := http.NewRequest(*httpMethod, reqURL, nil)
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/json")
	q := req.URL.Query()

	if params != nil && len(*params) > 0 {
		for key, value := range *params {
			q.Set(key, value)
		}
	}

	if authNeeded {
		nonce := time.Now().UnixNano()
		q.Set("apikey", ApiKey)
		q.Set("nonce", string(nonce))

		sign, err := generateAPISign(req.URL.String())
		if err != nil {
			return nil, err
		}
		req.Header.Add("apisign", sign)
	}

	req.URL.RawQuery = q.Encode()
	return req, err
}

func generateAPISign(url string) (sign string, err error) {
	mac := hmac.New(sha512.New, []byte(ApiSecret))
	_, err = mac.Write([]byte(url))
	return string(mac.Sum(nil)), err
}

// *** Public APIs ***
// Following functions are callable without authentication

func GetMarkets() (markets []Market, err error) {
	apiResp, err := apiQuery("getmarkets", "GET", nil)
	err = json.Unmarshal(apiResp.Result, &markets)
	return
}

func GetCurrencies() (currencies []Currency, err error) {
	apiResp, err := apiQuery("getcurrencies", "GET", nil)
	err = json.Unmarshal(apiResp.Result, &currencies)
	return
}

func GetTicker(market string) (ticker Ticker, err error) {
	apiResp, err := apiQuery("getticker", "GET", &map[string]string{"market": strings.ToUpper(market)})
	err = json.Unmarshal(apiResp.Result, &ticker)
	return
}

func GetMarketSummaries() (marketSummaries []MarketSummary, err error) {
	apiResp, err := apiQuery("getmarketsummaries", "GET", nil)
	err = json.Unmarshal(apiResp.Result, &marketSummaries)
	return
}

func GetMarketSummary(market string) (marketSummary []MarketSummary, err error) {
	apiResp, err := apiQuery("getmarketsummary", "GET", &map[string]string{"market": strings.ToUpper(market)})
	err = json.Unmarshal(apiResp.Result, &marketSummary)
	return
}

func GetOrderBook(market string, otype string) (orderBook OrderBook, err error) {
	apiResp, err := apiQuery("getorderbook", "GET", &map[string]string{"market": strings.ToUpper(market), "type": otype})
	err = json.Unmarshal(apiResp.Result, &orderBook)
	return
}

func GetMarketHistory(market string) (marketHistory []MarketHistoryEvent, err error) {
	apiResp, err := apiQuery("getmarkethistory", "GET", &map[string]string{"market": strings.ToUpper(market)})
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
