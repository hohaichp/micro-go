package emoji

import (
	"github.com/m3o/m3o-go/client"
)

func NewEmojiService(token string) *EmojiService {
	return &EmojiService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type EmojiService struct {
	client *client.Client
}

func (t *EmojiService) Find(request FindRequest) (*FindResponse, error) {
	rsp := &FindResponse{}
	return rsp, t.client.Call("emoji", "Find", request, rsp)
}

func (t *EmojiService) Flag(request FlagRequest) (*FlagResponse, error) {
	rsp := &FlagResponse{}
	return rsp, t.client.Call("emoji", "Flag", request, rsp)
}

func (t *EmojiService) Print(request PrintRequest) (*PrintResponse, error) {
	rsp := &PrintResponse{}
	return rsp, t.client.Call("emoji", "Print", request, rsp)
}

func (t *EmojiService) Send(request SendRequest) (*SendResponse, error) {
	rsp := &SendResponse{}
	return rsp, t.client.Call("emoji", "Send", request, rsp)
}

type FindRequest struct {
	Alias string `json:"alias"`
}

type FindResponse struct {
	Emoji string `json:"emoji"`
}

type FlagRequest struct {
	Code string `json:"code"`
}

type FlagResponse struct {
	Flag string `json:"flag"`
}

type PrintRequest struct {
	Text string `json:"text"`
}

type PrintResponse struct {
	Text string `json:"text"`
}

type SendRequest struct {
	From    string `json:"from"`
	Message string `json:"message"`
	To      string `json:"to"`
}

type SendResponse struct {
	Success bool `json:"success"`
}
