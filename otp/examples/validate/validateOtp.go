package example

import (
	"fmt"
	"github.com/micro/micro-go/otp"
)

// Validate the one time password created via generate
func ValidateOtp() {
	otpService := otp.NewOtpService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := otpService.Validate(otp.ValidateRequest{
		Code: "656211",
		Id:   "asim@example.com",
	})
	fmt.Println(rsp)
}
