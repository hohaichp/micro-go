package main

import (
	"fmt"
	"github.com/micro/micro-go/location"
)

func main() {
	locationService := location.NewLocationService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := locationService.Save(location.SaveRequest{
		Entity: location.Entity{
			Id: "1",
			Location: location.Point{
				Latitude:  51.511061,
				Longitude: -0.120022,
				Timestamp: 1622802761,
			},
			Type: "bike",
		},
	})
	fmt.Println(rsp)
}
