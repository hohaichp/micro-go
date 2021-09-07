package example

import (
	"fmt"
	"github.com/micro/micro-go/id"
)

func GenerateAshortId() {
	idService := id.NewIdService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := idService.Generate(id.GenerateRequest{
		Type: "shortid",
	})
	fmt.Println(rsp)
}
