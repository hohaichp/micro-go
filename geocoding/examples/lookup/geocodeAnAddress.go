package example

import (
	"fmt"
	"github.com/micro/micro-go/geocoding"
)

// Lookup returns a normalized address and gps location
func GeocodeAnAddress() {
	geocodingService := geocoding.NewGeocodingService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := geocodingService.Lookup(geocoding.LookupRequest{
		Address:  "10 russell st",
		City:     "london",
		Country:  "uk",
		Postcode: "wc2b",
	})
	fmt.Println(rsp)
}
