package gobittrex

import (
	"testing"
	"fmt"
)

func TestGetMarkets(t *testing.T) {
	GetMarkets()
}

func TestGetTicker(t *testing.T) {
	GetTicker("BTC-LTC")
}

func TestGetOrderBook(t *testing.T) {
	ob, _ := GetOrderBook("BTC-LTC", "both")
	fmt.Println(ob)
}
