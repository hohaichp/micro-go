package example

import (
	"fmt"
	"github.com/micro/micro-go/weather"
)

// Get the weather forecast right at this moment
func GetCurrentWeather() {
	weatherService := weather.NewWeatherService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := weatherService.Now(weather.NowRequest{
		Location: "london",
	})
	fmt.Println(rsp)
}
