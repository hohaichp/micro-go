package example

import (
	"fmt"
	"github.com/micro/micro-go/stock"
)

// Returns historic stock data for a given date
func GetHistoricData() {
	stockService := stock.NewStockService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := stockService.History(stock.HistoryRequest{
		Date:  "2020-10-01",
		Stock: "AAPL",
	})
	fmt.Println(rsp)
}
