package example

import (
	"fmt"
	"github.com/micro/micro-go/forex"
)

// Returns the last quote for a currency pair as bid and ask prices
func GetAfxQuote() {
	forexService := forex.NewForexService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := forexService.Quote(forex.QuoteRequest{
		Symbol: "GBPUSD",
	})
	fmt.Println(rsp)
}
