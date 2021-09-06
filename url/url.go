package url

import (
	"github.com/m3o/m3o-go/client"
)

func NewUrlService(token string) *UrlService {
	return &UrlService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type UrlService struct {
	client *client.Client
}

func (t *UrlService) List(request ListRequest) (*ListResponse, error) {
	rsp := &ListResponse{}
	return rsp, t.client.Call("url", "List", request, rsp)
}

func (t *UrlService) Proxy(request ProxyRequest) (*ProxyResponse, error) {
	rsp := &ProxyResponse{}
	return rsp, t.client.Call("url", "Proxy", request, rsp)
}

func (t *UrlService) Shorten(request ShortenRequest) (*ShortenResponse, error) {
	rsp := &ShortenResponse{}
	return rsp, t.client.Call("url", "Shorten", request, rsp)
}

type ListRequest struct {
	ShortUrl string `json:"shortUrl"`
}

type ListResponse struct {
	UrlPairs URLPair `json:"urlPairs"`
}

type ProxyRequest struct {
	ShortUrl string `json:"shortUrl"`
}

type ProxyResponse struct {
	DestinationUrl string `json:"destinationUrl"`
}

type ShortenRequest struct {
	DestinationUrl string `json:"destinationUrl"`
}

type ShortenResponse struct {
	ShortUrl string `json:"shortUrl"`
}

type URLPair struct {
	Created        int64  `json:"created"`
	DestinationUrl string `json:"destinationUrl"`
	HitCount       int64  `json:"hitCount"`
	Owner          string `json:"owner"`
	ShortUrl       string `json:"shortUrl"`
}
