package example

import (
	"fmt"
	"github.com/micro/micro-go/id"
)

func GenerateAsnowflakeId() {
	idService := id.NewIdService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := idService.Generate(id.GenerateRequest{
		Type: "snowflake",
	})
	fmt.Println(rsp)
}
