package example

import (
	"fmt"
	"github.com/micro/micro-go/rss"
)

// List the saved rss feeds
func ListRssFeeds() {
	rssService := rss.NewRssService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := rssService.List(rss.ListRequest{})
	fmt.Println(rsp)
}
