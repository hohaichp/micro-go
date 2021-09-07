package main

import (
	"fmt"
	"github.com/micro/micro-go/stock"
)

func main() {
	stockService := stock.NewStockService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := stockService.Quote(stock.QuoteRequest{
		Symbol: "AAPL",
	})
	fmt.Println(rsp)
}
