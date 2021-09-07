package main

import (
	"fmt"
	"github.com/micro/micro-go/stock"
)

func main() {
	stockService := stock.NewStockService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := stockService.Price(stock.PriceRequest{
		Symbol: "AAPL",
	})
	fmt.Println(rsp)
}
