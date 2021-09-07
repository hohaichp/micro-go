package main

import (
	"fmt"
	"github.com/micro/micro-go/currency"
)

func main() {
	currencyService := currency.NewCurrencyService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := currencyService.Codes(currency.CodesRequest{})
	fmt.Println(rsp)
}
