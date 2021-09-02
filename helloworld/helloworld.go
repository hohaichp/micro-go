package helloworld

import(
	"github.com/m3o/m3o-go/client"
)

func NewHelloworldService(token string) *HelloworldService {
	return &HelloworldService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type HelloworldService struct {
	client *client.Client
}


func (t *HelloworldService) Call(request CallRequest) (*CallResponse, error) {
	rsp := &CallResponse{}
	return rsp, t.client.Call("helloworld", "Call", request, rsp)
}

func (t *HelloworldService) Stream(request StreamRequest) (*StreamResponse, error) {
	rsp := &StreamResponse{}
	return rsp, t.client.Call("helloworld", "Stream", request, rsp)
}




type CallRequest struct {
  Name string `json:"name"`
}

type CallResponse struct {
  Message string `json:"message"`
}

type StreamRequest struct {
  Messages int64 `json:"messages"`
  Name string `json:"name"`
}

type StreamResponse struct {
  Message string `json:"message"`
}

