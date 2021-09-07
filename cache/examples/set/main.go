package main

import (
	"fmt"
	"github.com/micro/micro-go/cache"
)

func main() {
	cacheService := cache.NewCacheService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := cacheService.Set(cache.SetRequest{
		Key:   "foo",
		Value: "bar",
	})
	fmt.Println(rsp)
}
