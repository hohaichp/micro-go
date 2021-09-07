package example

import (
	"fmt"
	"github.com/micro/micro-go/cache"
)

// Decrement a value by 2
func DecrementAvalue() {
	cacheService := cache.NewCacheService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := cacheService.Decrement(cache.DecrementRequest{
		Key:   "counter",
		Value: 2,
	})
	fmt.Println(rsp)
}
