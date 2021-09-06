package forex

import (
	"github.com/m3o/m3o-go/client"
)

func NewForexService(token string) *ForexService {
	return &ForexService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type ForexService struct {
	client *client.Client
}

func (t *ForexService) History(request HistoryRequest) (*HistoryResponse, error) {
	rsp := &HistoryResponse{}
	return rsp, t.client.Call("forex", "History", request, rsp)
}

func (t *ForexService) Price(request PriceRequest) (*PriceResponse, error) {
	rsp := &PriceResponse{}
	return rsp, t.client.Call("forex", "Price", request, rsp)
}

func (t *ForexService) Quote(request QuoteRequest) (*QuoteResponse, error) {
	rsp := &QuoteResponse{}
	return rsp, t.client.Call("forex", "Quote", request, rsp)
}

type HistoryRequest struct {
	Symbol string `json:"symbol"`
}

type HistoryResponse struct {
	Close  float64 `json:"close"`
	Date   string  `json:"date"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Open   float64 `json:"open"`
	Symbol string  `json:"symbol"`
	Volume float64 `json:"volume"`
}

type PriceRequest struct {
	Symbol string `json:"symbol"`
}

type PriceResponse struct {
	Price  float64 `json:"price"`
	Symbol string  `json:"symbol"`
}

type QuoteRequest struct {
	Symbol string `json:"symbol"`
}

type QuoteResponse struct {
	AskPrice  float64 `json:"askPrice"`
	BidPrice  float64 `json:"bidPrice"`
	Symbol    string  `json:"symbol"`
	Timestamp string  `json:"timestamp"`
}
