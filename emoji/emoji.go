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
	// the alias code e.g :beer:
	Alias string `json:"alias"`
}

type FindResponse struct {
	// the unicode emoji 🍺
	Emoji string `json:"emoji"`
}

type FlagRequest struct {
	// country code e.g GB
	Code string `json:"code"`
}

type FlagResponse struct {
	// the emoji flag
	Flag string `json:"flag"`
}

type PrintRequest struct {
	// text including any alias e.g let's grab a :beer:
	Text string `json:"text"`
}

type PrintResponse struct {
	// text with rendered emojis
	Text string `json:"text"`
}

type SendRequest struct {
	// the name of the sender from e.g Alice
	From string `json:"from"`
	// message to send including emoji aliases
	Message string `json:"message"`
	// phone number to send to (including international dialing code)
	To string `json:"to"`
}

type SendResponse struct {
	// whether or not it succeeded
	Success bool `json:"success"`
}
