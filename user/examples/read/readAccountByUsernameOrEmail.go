package example

import (
	"fmt"
	"github.com/micro/micro-go/user"
)

func ReadAccountByUsernameOrEmail() {
	userService := user.NewUserService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := userService.Read(user.ReadRequest{
		Username: "usrname-1",
	})
	fmt.Println(rsp)
}
