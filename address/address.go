package address

import (
	"github.com/m3o/m3o-go/client"
)

func NewAddressService(token string) *AddressService {
	return &AddressService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type AddressService struct {
	client *client.Client
}

func (t *AddressService) LookupPostcode(request LookupPostcodeRequest) (*LookupPostcodeResponse, error) {
	rsp := &LookupPostcodeResponse{}
	return rsp, t.client.Call("address", "LookupPostcode", request, rsp)
}

type LookupPostcodeRequest struct {
	Postcode string `json:"postcode"`
}

type LookupPostcodeResponse struct {
	Addresses []Record `json:"addresses"`
}

type Record struct {
	BuildingName string `json:"buildingName"`
	County       string `json:"county"`
	LineOne      string `json:"lineOne"`
	LineTwo      string `json:"lineTwo"`
	Locality     string `json:"locality"`
	Organisation string `json:"organisation"`
	Postcode     string `json:"postcode"`
	Premise      string `json:"premise"`
	Street       string `json:"street"`
	Summary      string `json:"summary"`
	Town         string `json:"town"`
}
