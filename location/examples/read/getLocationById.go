package example

import (
	"fmt"
	"github.com/micro/micro-go/location"
)

// Lookup the location of an entity by ID
func GetLocationById() {
	locationService := location.NewLocationService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := locationService.Read(location.ReadRequest{
		Id: "1",
	})
	fmt.Println(rsp)
}
