package stock

import(
	"github.com/m3o/m3o-go/client"
)

func NewStockService(token string) *StockService {
	return &StockService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type StockService struct {
	client *client.Client
}


func (t *StockService) History(request HistoryRequest) (*HistoryResponse, error) {
	rsp := &HistoryResponse{}
	return rsp, t.client.Call("stock", "History", request, rsp)
}

func (t *StockService) OrderBook(request OrderBookRequest) (*OrderBookResponse, error) {
	rsp := &OrderBookResponse{}
	return rsp, t.client.Call("stock", "OrderBook", request, rsp)
}

func (t *StockService) Price(request PriceRequest) (*PriceResponse, error) {
	rsp := &PriceResponse{}
	return rsp, t.client.Call("stock", "Price", request, rsp)
}

func (t *StockService) Quote(request QuoteRequest) (*QuoteResponse, error) {
	rsp := &QuoteResponse{}
	return rsp, t.client.Call("stock", "Quote", request, rsp)
}




type HistoryRequest struct {
  Date string `json:"date"`
  Stock string `json:"stock"`
}

type HistoryResponse struct {
  Close float64 `json:"close"`
  Date string `json:"date"`
  High float64 `json:"high"`
  Low float64 `json:"low"`
  Open float64 `json:"open"`
  Symbol string `json:"symbol"`
  Volume int32 `json:"volume"`
}

type Order struct {
  AskPrice float64 `json:"askPrice"`
  AskSize int32 `json:"askSize"`
  BidPrice float64 `json:"bidPrice"`
  BidSize int32 `json:"bidSize"`
  Timestamp string `json:"timestamp"`
}

type OrderBookRequest struct {
  Date string `json:"date"`
  End string `json:"end"`
  Limit int32 `json:"limit"`
  Start string `json:"start"`
  Stock string `json:"stock"`
}

type OrderBookResponse struct {
  Date string `json:"date"`
  Orders []Order `json:"orders"`
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
  AskSize int32 `json:"askSize"`
  BidPrice float64 `json:"bidPrice"`
  BidSize int32 `json:"bidSize"`
  Symbol string `json:"symbol"`
  Timestamp string `json:"timestamp"`
}

