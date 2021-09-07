package main

import (
	"fmt"
	"github.com/micro/micro-go/url"
)

func main() {
	urlService := url.NewUrlService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := urlService.Shorten(url.ShortenRequest{})
	fmt.Println(rsp)
}
