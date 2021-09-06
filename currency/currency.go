package currency

import (
	"github.com/m3o/m3o-go/client"
)

func NewCurrencyService(token string) *CurrencyService {
	return &CurrencyService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type CurrencyService struct {
	client *client.Client
}

func (t *CurrencyService) Codes(request CodesRequest) (*CodesResponse, error) {
	rsp := &CodesResponse{}
	return rsp, t.client.Call("currency", "Codes", request, rsp)
}

func (t *CurrencyService) Convert(request ConvertRequest) (*ConvertResponse, error) {
	rsp := &ConvertResponse{}
	return rsp, t.client.Call("currency", "Convert", request, rsp)
}

func (t *CurrencyService) History(request HistoryRequest) (*HistoryResponse, error) {
	rsp := &HistoryResponse{}
	return rsp, t.client.Call("currency", "History", request, rsp)
}

func (t *CurrencyService) Rates(request RatesRequest) (*RatesResponse, error) {
	rsp := &RatesResponse{}
	return rsp, t.client.Call("currency", "Rates", request, rsp)
}

type Code struct {
	Currency string `json:"currency"`
	Name     string `json:"name"`
}

type CodesRequest struct {
}

type CodesResponse struct {
	Codes []Code `json:"codes"`
}

type ConvertRequest struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}

type ConvertResponse struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	Rate   float64 `json:"rate"`
	To     string  `json:"to"`
}

type HistoryRequest struct {
	Code string `json:"code"`
	Date string `json:"date"`
}

type HistoryResponse struct {
	Code  string             `json:"code"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

type RatesRequest struct {
	Code string `json:"code"`
}

type RatesResponse struct {
	Code  string             `json:"code"`
	Rates map[string]float64 `json:"rates"`
}
