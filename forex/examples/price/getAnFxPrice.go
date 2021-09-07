package example

import (
	"fmt"
	"github.com/micro/micro-go/forex"
)

// Returns the last traded price of a currency pair
func GetAnFxPrice() {
	forexService := forex.NewForexService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := forexService.Price(forex.PriceRequest{
		Symbol: "GBPUSD",
	})
	fmt.Println(rsp)
}
