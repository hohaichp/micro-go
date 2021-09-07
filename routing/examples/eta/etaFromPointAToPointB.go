package example

import (
	"fmt"
	"github.com/micro/micro-go/routing"
)

// Get an estimated time of arrival for two points
func EtaFromPointAtoPointB() {
	routingService := routing.NewRoutingService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := routingService.Eta(routing.EtaRequest{
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
