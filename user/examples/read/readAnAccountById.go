package example

import (
	"fmt"
	"github.com/micro/micro-go/user"
)

func ReadAnAccountById() {
	userService := user.NewUserService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := userService.Read(user.ReadRequest{
		Id: "usrid-1",
	})
	fmt.Println(rsp)
}
