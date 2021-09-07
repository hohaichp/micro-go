package example

import (
	"fmt"
	"github.com/micro/micro-go/id"
)

func GenerateAuniqueId() {
	idService := id.NewIdService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := idService.Generate(id.GenerateRequest{
		Type: "uuid",
	})
	fmt.Println(rsp)
}
