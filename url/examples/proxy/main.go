package main

import (
	"fmt"
	"github.com/micro/micro-go/url"
)

func main() {
	urlService := url.NewUrlService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := urlService.Proxy(url.ProxyRequest{})
	fmt.Println(rsp)
}
