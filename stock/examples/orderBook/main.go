package main

import (
	"fmt"
	"github.com/micro/micro-go/stock"
)

func main() {
	stockService := stock.NewStockService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := stockService.OrderBook(stock.OrderBookRequest{
		Date:  "2020-10-01",
		End:   "2020-10-01T11:00:00Z",
		Limit: 3,
		Start: "2020-10-01T10:00:00Z",
		Stock: "AAPL",
	})
	fmt.Println(rsp)
}
