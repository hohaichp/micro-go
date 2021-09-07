package example

import (
	"fmt"
	"github.com/micro/micro-go/rss"
)

// Read an rss feed by name
func ReadAfeed() {
	rssService := rss.NewRssService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := rssService.Feed(rss.FeedRequest{
		Name: "bbc",
	})
	fmt.Println(rsp)
}
