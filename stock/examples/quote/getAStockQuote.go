package example

import (
	"fmt"
	"github.com/micro/micro-go/stock"
)

// Returns the last quote for a stock including bid and ask prices
func GetAstockQuote() {
	stockService := stock.NewStockService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := stockService.Quote(stock.QuoteRequest{
		Symbol: "AAPL",
	})
	fmt.Println(rsp)
}
