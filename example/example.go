package main

import (
	"fmt"
	"github.com/michaelwolz/gobittrex"
)

func main() {
	gobittrex.ApiKey = ""
	gobittrex.ApiSecret = ""

	markets, err := gobittrex.GetMarkets()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, market := range markets {
			fmt.Printf("### %s ### \n", market.MarketName)
			fmt.Printf("MarketCurrency: %s\n", market.MarketCurrencyLong)
			fmt.Printf("BaseCurrency: %s\n\n", market.BaseCurrencyLong)
		}
	}
}