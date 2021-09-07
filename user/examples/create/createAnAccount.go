package example

import (
	"fmt"
	"github.com/micro/micro-go/user"
)

func CreateAnAccount() {
	userService := user.NewUserService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := userService.Create(user.CreateRequest{
		Email:    "joe@example.com",
		Id:       "usrid-1",
		Password: "mySecretPass123",
		Username: "usrname-1",
	})
	fmt.Println(rsp)
}
