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

type Balance struct {
	Currency      string
	Balance       decimal.Decimal
	Available     decimal.Decimal
	Pending       decimal.Decimal
	CryptoAddress string
	Requested     bool
	Uuid          string
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

type DepositAddress struct {
	Currency string
	Address  string
}

// Deposit- or WithdrawalEvent
type DWHistoryEvent struct {
	PaymentUuid    string
	Currency       string
	Amount         decimal.Decimal
	Address        string
	Opened         string
	Authorized     bool
	PendingPayment bool
	TxCost         decimal.Decimal
	TxId           string
	Canceled       bool
	InvalidAddress bool
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

type MarketHistoryEvent struct {
	Id        int
	TimeStamp string
	Quantity  decimal.Decimal
	Price     decimal.Decimal
	Total     decimal.Decimal
	FillType  string
	OrderType string
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

type OpenOrder struct {
	Uuid              string
	OrderUuid         string
	Exchange          string
	OrderType         string
	Quantity          decimal.Decimal
	QuantityRemaining decimal.Decimal
	Limit             decimal.Decimal
	CommissionPaid    decimal.Decimal
	Price             decimal.Decimal
	PricePerUnit      decimal.Decimal
	Opened            string
	Closed            string
	CancelInitiated   bool
	ImmediateOrCancel bool
	IsConditional     bool
	Condition         string
	ConditionTarget   string
}

type Order struct {
	AccountId                  string
	OrderUuid                  string
	Exchange                   string
	Type                       string
	Quantity                   decimal.Decimal
	QuantityRemaining          decimal.Decimal
	Limit                      decimal.Decimal
	Reserved                   decimal.Decimal
	ReserveRemaining           decimal.Decimal
	CommissionReserved         decimal.Decimal
	CommissionReserveRemaining decimal.Decimal
	CommissionPaid             decimal.Decimal
	Price                      decimal.Decimal
	PricePerUnit               decimal.Decimal
	Opened                     string
	Closed                     string
	IsOpen                     bool
	Sentinel                   string
	CancelInitiated            bool
	ImmediateOrCancel          bool
	IsConditional              bool
	Condition                  string
	ConditionTarget            string
}

type OrderBook struct {
	Buy  []OrderBookEntry `json:"buy"`
	Sell []OrderBookEntry `json:"sell"`
}

type OrderBookEntry struct {
	Quantity decimal.Decimal
	Rate     decimal.Decimal
}

type OrderHistoryEvent struct {
	OrderUuid         string
	Exchange          string
	TimeStamp         string
	OrderType         string
	Limit             decimal.Decimal
	Quantity          decimal.Decimal
	QuantityRemaining decimal.Decimal
	Commission        decimal.Decimal
	Price             decimal.Decimal
	PricePerUnit      decimal.Decimal
	IsConditional     bool
	Condition         string
	ConditionTarget   string
	ImmediateOrCancel bool
}

type OrderUUID struct {
	Uuid string
}

type Ticker struct {
	Bid  decimal.Decimal
	Ask  decimal.Decimal
	Last decimal.Decimal
}
