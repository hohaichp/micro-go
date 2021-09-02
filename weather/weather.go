package weather

import(
	"github.com/m3o/m3o-go/client"
)

func NewWeatherService(token string) *WeatherService {
	return &WeatherService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type WeatherService struct {
	client *client.Client
}


func (t *WeatherService) Forecast(request ForecastRequest) (*ForecastResponse, error) {
	rsp := &ForecastResponse{}
	return rsp, t.client.Call("weather", "Forecast", request, rsp)
}

func (t *WeatherService) Now(request NowRequest) (*NowResponse, error) {
	rsp := &NowResponse{}
	return rsp, t.client.Call("weather", "Now", request, rsp)
}




type Forecast struct {
  AvgTempC float64 `json:"avgTempC"`
  AvgTempF float64 `json:"avgTempF"`
  ChanceOfRain int32 `json:"chanceOfRain"`
  Condition string `json:"condition"`
  Date string `json:"date"`
  IconUrl string `json:"iconUrl"`
  MaxTempC float64 `json:"maxTempC"`
  MaxTempF float64 `json:"maxTempF"`
  MinTempC float64 `json:"minTempC"`
  MinTempF float64 `json:"minTempF"`
  Sunrise string `json:"sunrise"`
  Sunset string `json:"sunset"`
  WillItRain bool `json:"willItRain"`
}

type ForecastRequest struct {
  Days int32 `json:"days"`
  Location string `json:"location"`
}

type ForecastResponse struct {
  Country string `json:"country"`
  Forecast []Forecast `json:"forecast"`
  Latitude float64 `json:"latitude"`
  LocalTime string `json:"localTime"`
  Location string `json:"location"`
  Longitude float64 `json:"longitude"`
  Region string `json:"region"`
  Timezone string `json:"timezone"`
}

type NowRequest struct {
  Location string `json:"location"`
}

type NowResponse struct {
  Cloud int32 `json:"cloud"`
  Condition string `json:"condition"`
  Country string `json:"country"`
  Daytime bool `json:"daytime"`
  FeelsLikeC float64 `json:"feelsLikeC"`
  FeelsLikeF float64 `json:"feelsLikeF"`
  Humidity int32 `json:"humidity"`
  IconUrl string `json:"iconUrl"`
  Latitude float64 `json:"latitude"`
  LocalTime string `json:"localTime"`
  Location string `json:"location"`
  Longitude float64 `json:"longitude"`
  Region string `json:"region"`
  TempC float64 `json:"tempC"`
  TempF float64 `json:"tempF"`
  Timezone string `json:"timezone"`
  WindDegree int32 `json:"windDegree"`
  WindDirection string `json:"windDirection"`
  WindKph float64 `json:"windKph"`
  WindMph float64 `json:"windMph"`
}

