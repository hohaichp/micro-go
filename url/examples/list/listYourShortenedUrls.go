package example

import (
	"fmt"
	"github.com/micro/micro-go/url"
)

// List the token holder's shortened URLs
func ListYourShortenedUrls() {
	urlService := url.NewUrlService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := urlService.List(url.ListRequest{})
	fmt.Println(rsp)
}
