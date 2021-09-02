package routing

import(
	"github.com/m3o/m3o-go/client"
)

func NewRoutingService(token string) *RoutingService {
	return &RoutingService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type RoutingService struct {
	client *client.Client
}


func (t *RoutingService) Directions(request DirectionsRequest) (*DirectionsResponse, error) {
	rsp := &DirectionsResponse{}
	return rsp, t.client.Call("routing", "Directions", request, rsp)
}

func (t *RoutingService) Eta(request EtaRequest) (*EtaResponse, error) {
	rsp := &EtaResponse{}
	return rsp, t.client.Call("routing", "Eta", request, rsp)
}

func (t *RoutingService) Route(request RouteRequest) (*RouteResponse, error) {
	rsp := &RouteResponse{}
	return rsp, t.client.Call("routing", "Route", request, rsp)
}




type Direction struct {
  Distance float64 `json:"distance"`
  Duration float64 `json:"duration"`
  Instruction string `json:"instruction"`
  Intersections []Intersection `json:"intersections"`
  Maneuver Maneuver `json:"maneuver"`
  Name string `json:"name"`
  Reference string `json:"reference"`
}

type DirectionsRequest struct {
  Destination Point `json:"destination"`
  Origin Point `json:"origin"`
}

type DirectionsResponse struct {
  Directions []Direction `json:"directions"`
  Distance float64 `json:"distance"`
  Duration float64 `json:"duration"`
  Waypoints []Waypoint `json:"waypoints"`
}

type EtaRequest struct {
  Destination Point `json:"destination"`
  Origin Point `json:"origin"`
  Speed float64 `json:"speed"`
  Type string `json:"type"`
}

type EtaResponse struct {
  Duration float64 `json:"duration"`
}

type Intersection struct {
  Bearings []float64 `json:"bearings"`
  Location Point `json:"location"`
}

type Maneuver struct {
  Action string `json:"action"`
  BearingAfter float64 `json:"bearingAfter"`
  BearingBefore float64 `json:"bearingBefore"`
  Direction string `json:"direction"`
  Location Point `json:"location"`
}

type Point struct {
  Latitude float64 `json:"latitude"`
  Longitude float64 `json:"longitude"`
}

type RouteRequest struct {
  Destination Point `json:"destination"`
  Origin Point `json:"origin"`
}

type RouteResponse struct {
  Distance float64 `json:"distance"`
  Duration float64 `json:"duration"`
  Waypoints []Waypoint `json:"waypoints"`
}

type Waypoint struct {
  Location Point `json:"location"`
  Name string `json:"name"`
}

