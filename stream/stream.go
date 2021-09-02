package stream

import(
	"github.com/m3o/m3o-go/client"
)

func NewStreamService(token string) *StreamService {
	return &StreamService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type StreamService struct {
	client *client.Client
}


func (t *StreamService) Publish(request PublishRequest) (*PublishResponse, error) {
	rsp := &PublishResponse{}
	return rsp, t.client.Call("stream", "Publish", request, rsp)
}

func (t *StreamService) Subscribe(request SubscribeRequest) (*SubscribeResponse, error) {
	rsp := &SubscribeResponse{}
	return rsp, t.client.Call("stream", "Subscribe", request, rsp)
}




type PublishRequest struct {
  Message map[string]interface{} `json:"message"`
  Topic string `json:"topic"`
}

type PublishResponse struct {
}

type SubscribeRequest struct {
  Topic string `json:"topic"`
}

type SubscribeResponse struct {
  Message map[string]interface{} `json:"message"`
  Topic string `json:"topic"`
}

