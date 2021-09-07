package example

import (
	"fmt"
	"github.com/micro/micro-go/forex"
)

// Returns the data for the previous close
func GetPreviousClose() {
	forexService := forex.NewForexService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := forexService.History(forex.HistoryRequest{
		Symbol: "GBPUSD",
	})
	fmt.Println(rsp)
}
