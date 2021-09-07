package main

import (
	"fmt"
	"github.com/micro/micro-go/weather"
)

func main() {
	weatherService := weather.NewWeatherService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := weatherService.Now(weather.NowRequest{
		Location: "london",
	})
	fmt.Println(rsp)
}
