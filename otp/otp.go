package otp

import (
	"github.com/m3o/m3o-go/client"
)

func NewOtpService(token string) *OtpService {
	return &OtpService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type OtpService struct {
	client *client.Client
}

func (t *OtpService) Generate(request GenerateRequest) (*GenerateResponse, error) {
	rsp := &GenerateResponse{}
	return rsp, t.client.Call("otp", "Generate", request, rsp)
}

func (t *OtpService) Validate(request ValidateRequest) (*ValidateResponse, error) {
	rsp := &ValidateResponse{}
	return rsp, t.client.Call("otp", "Validate", request, rsp)
}

type GenerateRequest struct {
	Expiry int64  `json:"expiry"`
	Id     string `json:"id"`
	Size   int64  `json:"size"`
}

type GenerateResponse struct {
	Code string `json:"code"`
}

type ValidateRequest struct {
	Code string `json:"code"`
	Id   string `json:"id"`
}

type ValidateResponse struct {
	Success bool `json:"success"`
}
