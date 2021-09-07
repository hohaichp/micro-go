package main

import (
	"fmt"
	"github.com/micro/micro-go/helloworld"
)

func main() {
	helloworldService := helloworld.NewHelloworldService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := helloworldService.Call(helloworld.CallRequest{
		Name: "John",
	})
	fmt.Println(rsp)
}
