package example

import (
	"fmt"
	"github.com/micro/micro-go/cache"
)

// Set allows you to set a value
func SetAvalue() {
	cacheService := cache.NewCacheService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := cacheService.Set(cache.SetRequest{
		Key:   "foo",
		Value: "bar",
	})
	fmt.Println(rsp)
}
