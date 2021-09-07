package example

import (
	"fmt"
	"github.com/micro/micro-go/stock"
)

// Returns historic order book for a given date
func OrderBookHistory() {
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
