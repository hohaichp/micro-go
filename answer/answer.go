package answer

import (
	"github.com/m3o/m3o-go/client"
)

func NewAnswerService(token string) *AnswerService {
	return &AnswerService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type AnswerService struct {
	client *client.Client
}

func (t *AnswerService) Question(request QuestionRequest) (*QuestionResponse, error) {
	rsp := &QuestionResponse{}
	return rsp, t.client.Call("answer", "Question", request, rsp)
}

type QuestionRequest struct {
	Query string `json:"query"`
}

type QuestionResponse struct {
	Answer string `json:"answer"`
	Image  string `json:"image"`
	Url    string `json:"url"`
}
