package example

import (
	"fmt"
	"github.com/micro/micro-go/currency"
)

// Convert from USD to GBP
func ConvertUsdToGbp() {
	currencyService := currency.NewCurrencyService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := currencyService.Convert(currency.ConvertRequest{
		From: "USD",
		To:   "GBP",
	})
	fmt.Println(rsp)
}
