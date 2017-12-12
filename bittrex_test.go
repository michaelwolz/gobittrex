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
	ticker, err := GetTicker("BTC-DOGE")
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
	marketSummary, err := GetMarketSummary("BTC-DOGE")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(marketSummary, "\n")
	}
}

func TestGetOrderBook(t *testing.T) {
	orderBook, err := GetOrderBook("BTC-DOGE", "both")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(orderBook, "\n")
	}
}

func TestGetMarketHistory(t *testing.T) {
	marketHistory, err := GetMarketHistory("BTC-DOGE")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(marketHistory, "\n")
	}
}

// *** Market APIs ***

// Buy Order
/*func TestLimitOrder(t *testing.T) {
	buyOrderUuid, err := LimitOrder("buy", "BTC-DOGE", "5000.0", "0.00000010")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Printf("Buy-uuid: %s\n\n", buyOrderUuid)
	}
}*/

//Sell Order
//func TestLimitOrder2(t *testing.T) {}

/*func TestCancel(t *testing.T) {
	err := Cancel("8eb46010-24cf-4bbb-a48b-5dad0793043e")
	if err != nil {
		fmt.Println("WARNING!!! ORDER NOT CANCELED!\n")
		t.Errorf(err.Error())
	} else {
		fmt.Println("Buy-Order canceled!\n")
	}
}
*/

func TestGetOpenOrders(t *testing.T) {
	openOrders, err := GetOpenOrders("")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(openOrders, "\n")
	}
}

// *** Account APIs ***

func TestGetBalances(t *testing.T) {
	balances, err := GetBalances()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(balances, "\n")
	}
}

func TestGetBalance(t *testing.T) {
	balance, err := GetBalance("BTC")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(balance, "\n")
	}
}


func TestGetDepositAddress(t *testing.T) {
	depositAddress, err := GetDepositAddress("LTC")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(depositAddress, "\n")
	}
}
/*
func TestWithdraw(t *testing.T) {}
*/

func TestGetOrder(t *testing.T) {
	order, err := GetOrder("aa5eb312-1b5b-49fb-b2b6-2669d773926c")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(order, "\n")
	}
}


func TestGetOrderHistory(t *testing.T) {
	orderHistory, err := GetOrderHistory("BTC-LTC")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(orderHistory, "\n")
	}
}


func TestGetWithdrawalHistory(t *testing.T) {
	withdrawalHistory, err := GetWithdrawalHistory("ETH")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(withdrawalHistory, "\n")
	}
}

func TestGetDepositHistory(t *testing.T) {
	depositHistory, err := GetDepositHistory("ETH")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(depositHistory, "\n")
	}
}

