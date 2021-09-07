package main

import (
	"fmt"
	"github.com/micro/micro-go/forex"
)

func main() {
	forexService := forex.NewForexService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := forexService.History(forex.HistoryRequest{
		Symbol: "GBPUSD",
	})
	fmt.Println(rsp)
}
