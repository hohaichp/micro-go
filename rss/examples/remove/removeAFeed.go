package example

import (
	"fmt"
	"github.com/micro/micro-go/rss"
)

// Remove an rss feed from the crawler
func RemoveAfeed() {
	rssService := rss.NewRssService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := rssService.Remove(rss.RemoveRequest{
		Name: "bbc",
	})
	fmt.Println(rsp)
}
