package gobittrex

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

type APIResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

type Market struct {
	MarketCurrency     string
	BaseCurrency       string
	MarketCurrencyLong string
	BaseCurrencyLong   string
	MinTradeSize       decimal.Decimal
	MarketName         string
	IsActive           bool
	Created            string
}

type MarketSummary struct {
	MarketName        string
	High              decimal.Decimal
	Low               decimal.Decimal
	Volume            decimal.Decimal
	Last              decimal.Decimal
	BaseVolume        decimal.Decimal
	TimeStamp         string
	Bid               decimal.Decimal
	Ask               decimal.Decimal
	OpenBuyOrders     int
	OpenSellOrders    int
	PrevDay           decimal.Decimal
	Created           string
	DisplayMarketName string
}

type Currency struct {
	Currency        string
	CurrencyLong    string
	MinConfirmation int
	TxFee           decimal.Decimal
	IsActive        bool
	CoinType        string
	BaseAddress     string
}

type Ticker struct {
	Bid  decimal.Decimal
	Ask  decimal.Decimal
	Last decimal.Decimal
}

type OrderBookEntry struct {
	Quantity decimal.Decimal
	Rate     decimal.Decimal
}

type OrderBook struct {
	Buy  OrderBookEntry `json:"buy"`
	Sell OrderBookEntry `json:"sell"`
}

type MarketHistoryEvent struct {
	Id        int
	TimeStamp string
	Quantity  decimal.Decimal
	Price     decimal.Decimal
	Total     decimal.Decimal
	FillType  string
	OrderType string
}
