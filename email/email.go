package email

import (
	"github.com/m3o/m3o-go/client"
)

func NewEmailService(token string) *EmailService {
	return &EmailService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type EmailService struct {
	client *client.Client
}

func (t *EmailService) Send(request SendRequest) (*SendResponse, error) {
	rsp := &SendResponse{}
	return rsp, t.client.Call("email", "Send", request, rsp)
}

type SendRequest struct {
	From     string `json:"from"`
	HtmlBody string `json:"htmlBody"`
	ReplyTo  string `json:"replyTo"`
	Subject  string `json:"subject"`
	TextBody string `json:"textBody"`
	To       string `json:"to"`
}

type SendResponse struct {
}
