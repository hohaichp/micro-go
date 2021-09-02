package crypto

import(
	"github.com/m3o/m3o-go/client"
)

func NewCryptoService(token string) *CryptoService {
	return &CryptoService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type CryptoService struct {
	client *client.Client
}


func (t *CryptoService) History(request HistoryRequest) (*HistoryResponse, error) {
	rsp := &HistoryResponse{}
	return rsp, t.client.Call("crypto", "History", request, rsp)
}

func (t *CryptoService) News(request NewsRequest) (*NewsResponse, error) {
	rsp := &NewsResponse{}
	return rsp, t.client.Call("crypto", "News", request, rsp)
}

func (t *CryptoService) Price(request PriceRequest) (*PriceResponse, error) {
	rsp := &PriceResponse{}
	return rsp, t.client.Call("crypto", "Price", request, rsp)
}

func (t *CryptoService) Quote(request QuoteRequest) (*QuoteResponse, error) {
	rsp := &QuoteResponse{}
	return rsp, t.client.Call("crypto", "Quote", request, rsp)
}




type Article struct {
  Date string `json:"date"`
  Description string `json:"description"`
  Source string `json:"source"`
  Title string `json:"title"`
  Url string `json:"url"`
}

type HistoryRequest struct {
  Symbol string `json:"symbol"`
}

type HistoryResponse struct {
  Close float64 `json:"close"`
  Date string `json:"date"`
  High float64 `json:"high"`
  Low float64 `json:"low"`
  Open float64 `json:"open"`
  Symbol string `json:"symbol"`
  Volume float64 `json:"volume"`
}

type NewsRequest struct {
  Symbol string `json:"symbol"`
}

type NewsResponse struct {
  Articles []Article `json:"articles"`
  Symbol string `json:"symbol"`
}

type PriceRequest struct {
  Symbol string `json:"symbol"`
}

type PriceResponse struct {
  Price float64 `json:"price"`
  Symbol string `json:"symbol"`
}

type QuoteRequest struct {
  Symbol string `json:"symbol"`
}

type QuoteResponse struct {
  AskPrice float64 `json:"askPrice"`
  AskSize float64 `json:"askSize"`
  BidPrice float64 `json:"bidPrice"`
  BidSize float64 `json:"bidSize"`
  Symbol string `json:"symbol"`
  Timestamp string `json:"timestamp"`
}

