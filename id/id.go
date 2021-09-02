package id

import(
	"github.com/m3o/m3o-go/client"
)

func NewIdService(token string) *IdService {
	return &IdService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type IdService struct {
	client *client.Client
}


func (t *IdService) Generate(request GenerateRequest) (*GenerateResponse, error) {
	rsp := &GenerateResponse{}
	return rsp, t.client.Call("id", "Generate", request, rsp)
}

func (t *IdService) Types(request TypesRequest) (*TypesResponse, error) {
	rsp := &TypesResponse{}
	return rsp, t.client.Call("id", "Types", request, rsp)
}




type GenerateRequest struct {
  Type string `json:"type"`
}

type GenerateResponse struct {
  Id string `json:"id"`
  Type string `json:"type"`
}

type TypesRequest struct {
}

type TypesResponse struct {
  Types []string `json:"types"`
}
