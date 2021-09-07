package example

import (
	"fmt"
	"github.com/micro/micro-go/currency"
)

// Returns the rates for USD (sample response)
func GetRatesForUsd() {
	currencyService := currency.NewCurrencyService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := currencyService.Rates(currency.RatesRequest{
		Code: "USD",
	})
	fmt.Println(rsp)
}
