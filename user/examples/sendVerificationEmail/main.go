package main

import (
	"fmt"
	"github.com/micro/micro-go/user"
)

func main() {
	userService := user.NewUserService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := userService.SendVerificationEmail(user.SendVerificationEmailRequest{
		Email:              "joe@example.com",
		FailureRedirectUrl: "https://m3o.com/verification-failed",
		FromName:           "Awesome Dot Com",
		RedirectUrl:        "https://m3o.com",
		Subject:            "Email verification",
		TextContent: `Hi there,

Please verify your email by clicking this link: $micro_verification_link`,
	})
	fmt.Println(rsp)
}
