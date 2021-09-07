package example

import (
	"fmt"
	"github.com/micro/micro-go/cache"
)

// Increment a value by 2
func IncrementAvalue() {
	cacheService := cache.NewCacheService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := cacheService.Increment(cache.IncrementRequest{
		Key:   "counter",
		Value: 2,
	})
	fmt.Println(rsp)
}
