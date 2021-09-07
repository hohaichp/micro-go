package main

import (
	"fmt"
	"github.com/micro/micro-go/forex"
)

func main() {
	forexService := forex.NewForexService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := forexService.Price(forex.PriceRequest{
		Symbol: "GBPUSD",
	})
	fmt.Println(rsp)
}
