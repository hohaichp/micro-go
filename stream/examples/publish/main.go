package main

import (
	"fmt"
	"github.com/micro/micro-go/stream"
)

func main() {
	streamService := stream.NewStreamService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := streamService.Publish(stream.PublishRequest{
		Message: map[string]interface{}{
			"id":   "1",
			"type": "signup",
			"user": "john",
		},
		Topic: "events",
	})
	fmt.Println(rsp)
}
