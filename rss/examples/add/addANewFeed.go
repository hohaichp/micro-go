package example

import (
	"fmt"
	"github.com/micro/micro-go/rss"
)

// Add a new rss feed to the crawler
func AddAnewFeed() {
	rssService := rss.NewRssService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := rssService.Add(rss.AddRequest{
		Category: "news",
		Name:     "bbc",
		Url:      "http://feeds.bbci.co.uk/news/rss.xml",
	})
	fmt.Println(rsp)
}
