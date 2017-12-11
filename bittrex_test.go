package gobittrex

import (
	"testing"
	"fmt"
)

func TestGetMarkets(t *testing.T) {
	markets, err := GetMarkets()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(markets)
	}
}

func TestGetCurrencies(t *testing.T) {
	currencies, err := GetCurrencies()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(currencies)
	}
}

func TestGetTicker(t *testing.T) {
	ticker, err := GetTicker("BTC-LTC")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(ticker)
	}
}

func TestGetMarketSummaries(t *testing.T) {
	marketSummaries, err := GetMarketSummaries()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(marketSummaries)
	}
}

func TestGetMarketSummary(t *testing.T) {
	marketSummary, err := GetMarketSummary("BTC-LTC")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(marketSummary)
	}
}

func TestGetOrderBook(t *testing.T) {
	orderBook, err := GetOrderBook("BTC-LTC", "both")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(orderBook)
	}
}

func TestGetMarketHistory(t *testing.T) {
	marketHistory, err := GetMarketHistory("BTC-LTC")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(marketHistory)
	}
}
