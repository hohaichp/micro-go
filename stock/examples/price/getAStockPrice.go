package example

import (
	"fmt"
	"github.com/micro/micro-go/stock"
)

// Returns the last traded price of a stock
func GetAstockPrice() {
	stockService := stock.NewStockService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := stockService.Price(stock.PriceRequest{
		Symbol: "AAPL",
	})
	fmt.Println(rsp)
}
