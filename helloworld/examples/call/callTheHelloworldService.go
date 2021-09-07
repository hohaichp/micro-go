package example

import (
	"fmt"
	"github.com/micro/micro-go/helloworld"
)

// Helloworld is a simple example that takes a name and returns a personalised message
func CallTheHelloworldService() {
	helloworldService := helloworld.NewHelloworldService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := helloworldService.Call(helloworld.CallRequest{
		Name: "John",
	})
	fmt.Println(rsp)
}
