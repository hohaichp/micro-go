package sms

import(
	"github.com/m3o/m3o-go/client"
)

func NewSmsService(token string) *SmsService {
	return &SmsService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type SmsService struct {
	client *client.Client
}


func (t *SmsService) Send(request SendRequest) (*SendResponse, error) {
	rsp := &SendResponse{}
	return rsp, t.client.Call("sms", "Send", request, rsp)
}




type SendRequest struct {
  From string `json:"from"`
  Message string `json:"message"`
  To string `json:"to"`
}

type SendResponse struct {
  Info string `json:"info"`
  Status string `json:"status"`
}

