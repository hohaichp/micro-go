package main

import (
	"fmt"
	"github.com/micro/micro-go/rss"
)

func main() {
	rssService := rss.NewRssService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := rssService.List(rss.ListRequest{})
	fmt.Println(rsp)
}
