package example

import (
	"fmt"
	"github.com/micro/micro-go/cache"
)

// Delete a value from the cache
func DeleteAvalue() {
	cacheService := cache.NewCacheService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := cacheService.Delete(cache.DeleteRequest{
		Key: "foo",
	})
	fmt.Println(rsp)
}
