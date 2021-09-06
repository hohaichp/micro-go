package time

import (
	"github.com/m3o/m3o-go/client"
)

func NewTimeService(token string) *TimeService {
	return &TimeService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type TimeService struct {
	client *client.Client
}

func (t *TimeService) Now(request NowRequest) (*NowResponse, error) {
	rsp := &NowResponse{}
	return rsp, t.client.Call("time", "Now", request, rsp)
}

func (t *TimeService) Zone(request ZoneRequest) (*ZoneResponse, error) {
	rsp := &ZoneResponse{}
	return rsp, t.client.Call("time", "Zone", request, rsp)
}

type NowRequest struct {
	Location string `json:"location"`
}

type NowResponse struct {
	Localtime string `json:"localtime"`
	Location  string `json:"location"`
	Timestamp string `json:"timestamp"`
	Timezone  string `json:"timezone"`
	Unix      int64  `json:"unix"`
}

type ZoneRequest struct {
	Location string `json:"location"`
}

type ZoneResponse struct {
	Abbreviation string  `json:"abbreviation"`
	Country      string  `json:"country"`
	Dst          bool    `json:"dst"`
	Latitude     float64 `json:"latitude"`
	Localtime    string  `json:"localtime"`
	Location     string  `json:"location"`
	Longitude    float64 `json:"longitude"`
	Region       string  `json:"region"`
	Timezone     string  `json:"timezone"`
}
