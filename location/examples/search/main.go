package main

import (
	"fmt"
	"github.com/micro/micro-go/location"
)

func main() {
	locationService := location.NewLocationService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := locationService.Search(location.SearchRequest{
		Center: location.Point{
			Latitude:  51.511061,
			Longitude: -0.120022,
		},
		NumEntities: 10,
		Radius:      100,
		Type:        "bike",
	})
	fmt.Println(rsp)
}
