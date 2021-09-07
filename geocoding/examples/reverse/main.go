package main

import (
	"fmt"
	"github.com/micro/micro-go/geocoding"
)

func main() {
	geocodingService := geocoding.NewGeocodingService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := geocodingService.Reverse(geocoding.ReverseRequest{
		Latitude:  51.5123064,
		Longitude: -0.1216235,
	})
	fmt.Println(rsp)
}
