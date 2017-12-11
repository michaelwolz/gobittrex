package gobittrex

import (
	"testing"
	"fmt"
)

// *** Public APIs ***

func TestGetMarkets(t *testing.T) {
	markets, err := GetMarkets()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(markets, "\n")
	}
}

func TestGetCurrencies(t *testing.T) {
	currencies, err := GetCurrencies()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(currencies, "\n")
	}
}

func TestGetTicker(t *testing.T) {
	ticker, err := GetTicker("BTC-LTC")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(ticker, "\n")
	}
}

func TestGetMarketSummaries(t *testing.T) {
	marketSummaries, err := GetMarketSummaries()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(marketSummaries, "\n")
	}
}

func TestGetMarketSummary(t *testing.T) {
	marketSummary, err := GetMarketSummary("BTC-LTC")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(marketSummary, "\n")
	}
}

func TestGetOrderBook(t *testing.T) {
	orderBook, err := GetOrderBook("BTC-LTC", "both")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(orderBook, "\n")
	}
}

func TestGetMarketHistory(t *testing.T) {
	marketHistory, err := GetMarketHistory("BTC-LTC")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(marketHistory, "\n")
	}
}

// *** Market APIs ***

/*
func TestBuyLimit(t *testing.T) {}

func TestSellLimit(t *testing.T) {}

func TestCancel(t *testing.T) {}

func TestGetOpenOrders(t *testing.T) {}
*/

// *** Account APIs ***

/*
func TestGetBalances(t *testing.T) {}

func TestGetBalance(t *testing.T) {}

func TestGetDepositAddress(t *testing.T) {}

func TestWithdraw(t *testing.T) {}

func TestGetOrder(t *testing.T) {}

func TestGetOrderHistory(t *testing.T) {}

func TestGetWithdrawalHistory(t *testing.T) {}

func TestGetDepositHistory(t *testing.T) {}
*/
