package example

import (
	"fmt"
	"github.com/micro/micro-go/currency"
)

// Returns the supported currency codes (example response)
func GetSupportedCodes() {
	currencyService := currency.NewCurrencyService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := currencyService.Codes(currency.CodesRequest{})
	fmt.Println(rsp)
}
