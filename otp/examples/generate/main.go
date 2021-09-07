package main

import (
	"fmt"
	"github.com/micro/micro-go/otp"
)

func main() {
	otpService := otp.NewOtpService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := otpService.Generate(otp.GenerateRequest{
		Id: "asim@example.com",
	})
	fmt.Println(rsp)
}
