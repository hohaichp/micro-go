package example

import (
	"fmt"
	"github.com/micro/micro-go/routing"
)

// Get a set of turn by turn directions and manuevers
func TurnByTurnDirections() {
	routingService := routing.NewRoutingService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := routingService.Directions(routing.DirectionsRequest{
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
