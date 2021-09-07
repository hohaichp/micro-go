package example

import (
	"fmt"
	"github.com/micro/micro-go/user"
)

func ReadAccountByEmail() {
	userService := user.NewUserService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := userService.Read(user.ReadRequest{
		Email: "joe@example.com",
	})
	fmt.Println(rsp)
}
