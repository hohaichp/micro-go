package example

import (
	"fmt"
	"github.com/micro/micro-go/stream"
)

// Publish a message to a topic on the stream
func PublishAmessage() {
	streamService := stream.NewStreamService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := streamService.Publish(stream.PublishRequest{
		Message: map[string]interface{}{
			"user": "john",
			"id":   "1",
			"type": "signup",
		},
		Topic: "events",
	})
	fmt.Println(rsp)
}
