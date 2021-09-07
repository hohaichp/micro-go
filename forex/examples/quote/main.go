package main

import (
	"fmt"
	"github.com/micro/micro-go/forex"
)

func main() {
	forexService := forex.NewForexService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := forexService.Quote(forex.QuoteRequest{
		Symbol: "GBPUSD",
	})
	fmt.Println(rsp)
}
