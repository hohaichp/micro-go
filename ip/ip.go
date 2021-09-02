package ip

import(
	"github.com/m3o/m3o-go/client"
)

func NewIpService(token string) *IpService {
	return &IpService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type IpService struct {
	client *client.Client
}


func (t *IpService) Lookup(request LookupRequest) (*LookupResponse, error) {
	rsp := &LookupResponse{}
	return rsp, t.client.Call("ip", "Lookup", request, rsp)
}




type LookupRequest struct {
  Ip string `json:"ip"`
}

type LookupResponse struct {
  Asn int32 `json:"asn"`
  City string `json:"city"`
  Continent string `json:"continent"`
  Country string `json:"country"`
  Ip string `json:"ip"`
  Latitude float64 `json:"latitude"`
  Longitude float64 `json:"longitude"`
  Timezone string `json:"timezone"`
}

