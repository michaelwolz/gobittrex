package gobittrex

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"time"
	"crypto/hmac"
	"crypto/sha512"
	"strings"
	"encoding/hex"
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
	"buylimit":             1,
	"selllimit":            1,
	"cancel":               1,
	"getopenorders":        1,
	"getbalances":          2,
	"getbalance":           2,
	"getdepositaddress":    2,
	"withdraw":             2,
	"getorder":             2,
	"getorderhistory":      2,
	"getwithdrawalhistory": 2,
	"getdeposithistory":    2,
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
		err = errors.New(fmt.Sprintf("API-Call-Error: %s", apiResp.Message))
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

	if params != nil {
		for key, value := range *params {
			q.Set(key, value)
		}
	}

	req.URL.RawQuery = q.Encode()

	if authNeeded {
		nonce := time.Now().UnixNano()
		q.Set("apikey", ApiKey)
		q.Set("nonce", fmt.Sprintf("%d", nonce))
		req.URL.RawQuery = q.Encode()

		sign, err := generateAPISign(req.URL.String())
		if err != nil {
			return nil, err
		}
		req.Header.Add("apisign", sign)
	}
	return req, err
}

func generateAPISign(url string) (sign string, err error) {
	mac := hmac.New(sha512.New, []byte(ApiSecret))
	_, err = mac.Write([]byte(url))
	return hex.EncodeToString(mac.Sum(nil)), err
}

// *** Public APIs ***
// Following functions are callable without authentication

func GetMarkets() (markets []Market, err error) {
	apiResp, err := apiQuery("getmarkets", "GET", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &markets)
	return
}

func GetCurrencies() (currencies []Currency, err error) {
	apiResp, err := apiQuery("getcurrencies", "GET", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &currencies)
	return
}

func GetTicker(market string) (ticker Ticker, err error) {
	apiResp, err := apiQuery("getticker", "GET", &map[string]string{"market": strings.ToUpper(market)})
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &ticker)
	return
}

func GetMarketSummaries() (marketSummaries []MarketSummary, err error) {
	apiResp, err := apiQuery("getmarketsummaries", "GET", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &marketSummaries)
	return
}

func GetMarketSummary(market string) (marketSummary []MarketSummary, err error) {
	apiResp, err := apiQuery("getmarketsummary", "GET", &map[string]string{"market": strings.ToUpper(market)})
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &marketSummary)
	return
}

func GetOrderBook(market string, otype string) (orderBook OrderBook, err error) {
	apiResp, err := apiQuery("getorderbook", "GET", &map[string]string{"market": strings.ToUpper(market), "type": otype})
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &orderBook)
	return
}

func GetMarketHistory(market string) (marketHistory []MarketHistoryEvent, err error) {
	apiResp, err := apiQuery("getmarkethistory", "GET", &map[string]string{"market": strings.ToUpper(market)})
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &marketHistory)
	return
}

// *** Market APIs ***
// Following functions are only callable with authentication

// While orderType is sell or buy
func LimitOrder(orderType, market, quantity, rate string) (orderUuid string, err error) {
	params := map[string]string{
		"market":   strings.ToUpper(market),
		"quantity": fmt.Sprintf("%s", quantity),
		"rate":     fmt.Sprintf("%s", rate),
	}

	if orderType != "buy" && orderType != "sell" {
		err = errors.New(fmt.Sprintf("There is no ordertype: %s", orderType))
	}

	apiResp, err := apiQuery(fmt.Sprintf("%slimit", orderType), "GET", &params)

	if err != nil {
		return
	}

	var uuid UUID
	err = json.Unmarshal(apiResp.Result, &uuid)

	return uuid.Uuid, err
}

func Cancel(uuid string) (err error) {
	_, err = apiQuery("cancel", "GET", &map[string]string{"uuid": uuid})
	return
}

func GetOpenOrders(market string) (openOrders []OpenOrder, err error) {
	var params map[string]string
	if market != "" {
		params = make(map[string]string, 1)
		params["market"] = market

	}

	apiResp, err := apiQuery("getopenorders", "GET", &params)
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &openOrders)
	return
}

// *** Account APIs ***
// Following functions are only callable with authentication

func GetBalances() (balances []Balance, err error) {
	apiResp, err := apiQuery("getbalances", "GET", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &balances)
	return
}

func GetBalance(currency string) (balance Balance, err error) {
	apiResp, err := apiQuery("getbalance", "GET", &map[string]string{"currency": strings.ToUpper(currency)})
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &balance)
	return
}

func GetDepositAddress(currency string) (depositAddress DepositAddress, err error) {
	apiResp, err := apiQuery("getdepositaddress", "GET", &map[string]string{"currency": strings.ToUpper(currency)})
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &depositAddress)
	return
}

func Withdraw(currency string, quantity string, address string) (withdrawalUUID string, err error) {
	params := map[string]string{
		"currency": currency,
		"quantity": quantity,
		"address":  address,
	}

	apiResp, err := apiQuery("getbalances", "GET", &params)
	if err != nil {
		return
	}

	var uuid UUID
	err = json.Unmarshal(apiResp.Result, &uuid)
	return uuid.Uuid, err
}

func GetOrder(uuid string) (order Order, err error) {
	apiResp, err := apiQuery("getorder", "GET", &map[string]string{"uuid": uuid})
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &order)
	return
}

func GetOrderHistory(market string) (orderHistory []OrderHistoryEvent, err error) {
	apiResp, err := apiQuery("getorderhistory", "GET", &map[string]string{"market": strings.ToUpper(market)})
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &orderHistory)
	return
}

func GetWithdrawalHistory(currency string) (withdrawalHistory []DWHistoryEvent, err error) {
	apiResp, err := apiQuery("getwithdrawalhistory", "GET", &map[string]string{"currency": strings.ToUpper(currency)})
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &withdrawalHistory)
	return
}

func GetDepositHistory(currency string) (depositHistory []DWHistoryEvent, err error) {
	apiResp, err := apiQuery("getdeposithistory", "GET", &map[string]string{"currency": strings.ToUpper(currency)})
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResp.Result, &depositHistory)
	return
}
