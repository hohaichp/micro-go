package main

import (
	"fmt"
	"github.com/micro/micro-go/currency"
)

func main() {
	currencyService := currency.NewCurrencyService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := currencyService.Convert(currency.ConvertRequest{
		Amount: 10,
		From:   "USD",
		To:     "GBP",
	})
	fmt.Println(rsp)
}
