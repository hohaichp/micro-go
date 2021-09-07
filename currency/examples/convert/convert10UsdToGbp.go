package example

import (
	"fmt"
	"github.com/micro/micro-go/currency"
)

// Convert $10 from USD to GBP
func Convert10usdToGbp() {
	currencyService := currency.NewCurrencyService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := currencyService.Convert(currency.ConvertRequest{
		Amount: 10,
		From:   "USD",
		To:     "GBP",
	})
	fmt.Println(rsp)
}
