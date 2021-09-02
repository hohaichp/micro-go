package location

import(
	"github.com/m3o/m3o-go/client"
)

func NewLocationService(token string) *LocationService {
	return &LocationService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type LocationService struct {
	client *client.Client
}


func (t *LocationService) Read(request ReadRequest) (*ReadResponse, error) {
	rsp := &ReadResponse{}
	return rsp, t.client.Call("location", "Read", request, rsp)
}

func (t *LocationService) Save(request SaveRequest) (*SaveResponse, error) {
	rsp := &SaveResponse{}
	return rsp, t.client.Call("location", "Save", request, rsp)
}

func (t *LocationService) Search(request SearchRequest) (*SearchResponse, error) {
	rsp := &SearchResponse{}
	return rsp, t.client.Call("location", "Search", request, rsp)
}




type Entity struct {
  Id string `json:"id"`
  Location Point `json:"location"`
  Type string `json:"type"`
}

type Point struct {
  Latitude float64 `json:"latitude"`
  Longitude float64 `json:"longitude"`
  Timestamp int64 `json:"timestamp"`
}

type ReadRequest struct {
  Id string `json:"id"`
}

type ReadResponse struct {
  Entity Entity `json:"entity"`
}

type SaveRequest struct {
  Entity Entity `json:"entity"`
}

type SaveResponse struct {
}

type SearchRequest struct {
  Center Point `json:"center"`
  NumEntities int64 `json:"numEntities"`
  Radius float64 `json:"radius"`
  Type string `json:"type"`
}

type SearchResponse struct {
  Entities []Entity `json:"entities"`
}

