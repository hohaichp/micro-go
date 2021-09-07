package main

import (
	"fmt"
	"github.com/micro/micro-go/url"
)

func main() {
	urlService := url.NewUrlService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := urlService.List(url.ListRequest{})
	fmt.Println(rsp)
}
