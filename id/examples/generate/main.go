package main

import (
	"fmt"
	"github.com/micro/micro-go/id"
)

func main() {
	idService := id.NewIdService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := idService.Generate(id.GenerateRequest{
		Type: "bigflake",
	})
	fmt.Println(rsp)
}
