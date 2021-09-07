package main

import (
	"fmt"
	"github.com/micro/micro-go/otp"
)

func main() {
	otpService := otp.NewOtpService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := otpService.Validate(otp.ValidateRequest{
		Code: "656211",
		Id:   "asim@example.com",
	})
	fmt.Println(rsp)
}
