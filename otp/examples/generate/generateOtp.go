package example

import (
	"fmt"
	"github.com/micro/micro-go/otp"
)

// Generate a one time password for a unique id, email or user
func GenerateOtp() {
	otpService := otp.NewOtpService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := otpService.Generate(otp.GenerateRequest{
		Id: "asim@example.com",
	})
	fmt.Println(rsp)
}
