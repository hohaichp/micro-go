package main

import (
	"fmt"
	"github.com/micro/micro-go/location"
)

func main() {
	locationService := location.NewLocationService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := locationService.Read(location.ReadRequest{
		Id: "1",
	})
	fmt.Println(rsp)
}
