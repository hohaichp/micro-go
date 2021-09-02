package sentiment

import(
	"github.com/m3o/m3o-go/client"
)

func NewSentimentService(token string) *SentimentService {
	return &SentimentService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type SentimentService struct {
	client *client.Client
}


func (t *SentimentService) Analyze(request AnalyzeRequest) (*AnalyzeResponse, error) {
	rsp := &AnalyzeResponse{}
	return rsp, t.client.Call("sentiment", "Analyze", request, rsp)
}




type AnalyzeRequest struct {
  Lang string `json:"lang"`
  Text string `json:"text"`
}

type AnalyzeResponse struct {
  Score float64 `json:"score"`
}

