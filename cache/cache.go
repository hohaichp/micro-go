package cache

import(
	"github.com/m3o/m3o-go/client"
)

func NewCacheService(token string) *CacheService {
	return &CacheService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type CacheService struct {
	client *client.Client
}


func (t *CacheService) Decrement(request DecrementRequest) (*DecrementResponse, error) {
	rsp := &DecrementResponse{}
	return rsp, t.client.Call("cache", "Decrement", request, rsp)
}

func (t *CacheService) Delete(request DeleteRequest) (*DeleteResponse, error) {
	rsp := &DeleteResponse{}
	return rsp, t.client.Call("cache", "Delete", request, rsp)
}

func (t *CacheService) Get(request GetRequest) (*GetResponse, error) {
	rsp := &GetResponse{}
	return rsp, t.client.Call("cache", "Get", request, rsp)
}

func (t *CacheService) Increment(request IncrementRequest) (*IncrementResponse, error) {
	rsp := &IncrementResponse{}
	return rsp, t.client.Call("cache", "Increment", request, rsp)
}

func (t *CacheService) Set(request SetRequest) (*SetResponse, error) {
	rsp := &SetResponse{}
	return rsp, t.client.Call("cache", "Set", request, rsp)
}




type DecrementRequest struct {
  Key string `json:"key"`
  Value int64 `json:"value"`
}

type DecrementResponse struct {
  Key string `json:"key"`
  Value int64 `json:"value"`
}

type DeleteRequest struct {
  Key string `json:"key"`
}

type DeleteResponse struct {
  Status string `json:"status"`
}

type GetRequest struct {
  Key string `json:"key"`
}

type GetResponse struct {
  Key string `json:"key"`
  Ttl int64 `json:"ttl"`
  Value string `json:"value"`
}

type IncrementRequest struct {
  Key string `json:"key"`
  Value int64 `json:"value"`
}

type IncrementResponse struct {
  Key string `json:"key"`
  Value int64 `json:"value"`
}

type SetRequest struct {
  Key string `json:"key"`
  Ttl int64 `json:"ttl"`
  Value string `json:"value"`
}

type SetResponse struct {
  Status string `json:"status"`
}

