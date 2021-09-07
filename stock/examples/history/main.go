package main

import (
	"fmt"
	"github.com/micro/micro-go/stock"
)

func main() {
	stockService := stock.NewStockService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := stockService.History(stock.HistoryRequest{
		Date:  "2020-10-01",
		Stock: "AAPL",
	})
	fmt.Println(rsp)
}
