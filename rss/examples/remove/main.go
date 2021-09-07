package main

import (
	"fmt"
	"github.com/micro/micro-go/rss"
)

func main() {
	rssService := rss.NewRssService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := rssService.Remove(rss.RemoveRequest{
		Name: "bbc",
	})
	fmt.Println(rsp)
}
