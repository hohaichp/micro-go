package example

import (
	"fmt"
	"github.com/micro/micro-go/routing"
)

// Get the GPS points for a route
func GpsPointsForAroute() {
	routingService := routing.NewRoutingService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := routingService.Route(routing.RouteRequest{
		Destination: routing.Point{
			Latitude:  52.529407,
			Longitude: 13.397634,
		},
		Origin: routing.Point{
			Latitude:  52.517037,
			Longitude: 13.38886,
		},
	})
	fmt.Println(rsp)
}
