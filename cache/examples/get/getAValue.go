package example

import (
	"fmt"
	"github.com/micro/micro-go/cache"
)

// Get returns a value from the cache
func GetAvalue() {
	cacheService := cache.NewCacheService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := cacheService.Get(cache.GetRequest{
		Key: "foo",
	})
	fmt.Println(rsp)
}
