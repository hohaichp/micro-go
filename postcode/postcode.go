package postcode

import (
	"github.com/m3o/m3o-go/client"
)

func NewPostcodeService(token string) *PostcodeService {
	return &PostcodeService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type PostcodeService struct {
	client *client.Client
}

func (t *PostcodeService) Lookup(request LookupRequest) (*LookupResponse, error) {
	rsp := &LookupResponse{}
	return rsp, t.client.Call("postcode", "Lookup", request, rsp)
}

func (t *PostcodeService) Random(request RandomRequest) (*RandomResponse, error) {
	rsp := &RandomResponse{}
	return rsp, t.client.Call("postcode", "Random", request, rsp)
}

func (t *PostcodeService) Validate(request ValidateRequest) (*ValidateResponse, error) {
	rsp := &ValidateResponse{}
	return rsp, t.client.Call("postcode", "Validate", request, rsp)
}

type LookupRequest struct {
	Postcode string `json:"postcode"`
}

type LookupResponse struct {
	Country   string  `json:"country"`
	District  string  `json:"district"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Postcode  string  `json:"postcode"`
	Region    string  `json:"region"`
	Ward      string  `json:"ward"`
}

type RandomRequest struct {
}

type RandomResponse struct {
	Country   string  `json:"country"`
	District  string  `json:"district"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Postcode  string  `json:"postcode"`
	Region    string  `json:"region"`
	Ward      string  `json:"ward"`
}

type ValidateRequest struct {
	Postcode string `json:"postcode"`
}

type ValidateResponse struct {
	Valid bool `json:"valid"`
}
