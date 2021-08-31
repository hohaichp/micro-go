package m3o

import (
	"github.com/m3o/m3o-go/client"
)

func NewClient(token string) *Client {
	return &Client{
		token: token,

		AnswerService:     NewAnswerService(token),
		CacheService:      NewCacheService(token),
		CryptoService:     NewCryptoService(token),
		CurrencyService:   NewCurrencyService(token),
		DbService:         NewDbService(token),
		EmailService:      NewEmailService(token),
		EmojiService:      NewEmojiService(token),
		FileService:       NewFileService(token),
		ForexService:      NewForexService(token),
		GeocodingService:  NewGeocodingService(token),
		HelloworldService: NewHelloworldService(token),
		IdService:         NewIdService(token),
		ImageService:      NewImageService(token),
		IpService:         NewIpService(token),
		LocationService:   NewLocationService(token),
		OtpService:        NewOtpService(token),
		RoutingService:    NewRoutingService(token),
		RssService:        NewRssService(token),
		SentimentService:  NewSentimentService(token),
		SmsService:        NewSmsService(token),
		StockService:      NewStockService(token),
		StreamService:     NewStreamService(token),
		ThumbnailService:  NewThumbnailService(token),
		TimeService:       NewTimeService(token),
		UrlService:        NewUrlService(token),
		UserService:       NewUserService(token),
		WeatherService:    NewWeatherService(token),
	}
}

type Client struct {
	token string

	AnswerService     *AnswerService
	CacheService      *CacheService
	CryptoService     *CryptoService
	CurrencyService   *CurrencyService
	DbService         *DbService
	EmailService      *EmailService
	EmojiService      *EmojiService
	FileService       *FileService
	ForexService      *ForexService
	GeocodingService  *GeocodingService
	HelloworldService *HelloworldService
	IdService         *IdService
	ImageService      *ImageService
	IpService         *IpService
	LocationService   *LocationService
	OtpService        *OtpService
	RoutingService    *RoutingService
	RssService        *RssService
	SentimentService  *SentimentService
	SmsService        *SmsService
	StockService      *StockService
	StreamService     *StreamService
	ThumbnailService  *ThumbnailService
	TimeService       *TimeService
	UrlService        *UrlService
	UserService       *UserService
	WeatherService    *WeatherService
}

func NewAnswerService(token string) *AnswerService {
	return &AnswerService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type AnswerService struct {
	client *client.Client
}

func (t *AnswerService) Question(request AnswerQuestionRequest) (*AnswerQuestionResponse, error) {
	rsp := &AnswerQuestionResponse{}
	return rsp, t.client.Call("answer", "Question", request, rsp)
}

func NewCacheService(token string) *CacheService {
	return &CacheService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type CacheService struct {
	client *client.Client
}

func (t *CacheService) Decrement(request CacheDecrementRequest) (*CacheDecrementResponse, error) {
	rsp := &CacheDecrementResponse{}
	return rsp, t.client.Call("cache", "Decrement", request, rsp)
}

func (t *CacheService) Delete(request CacheDeleteRequest) (*CacheDeleteResponse, error) {
	rsp := &CacheDeleteResponse{}
	return rsp, t.client.Call("cache", "Delete", request, rsp)
}

func (t *CacheService) Get(request CacheGetRequest) (*CacheGetResponse, error) {
	rsp := &CacheGetResponse{}
	return rsp, t.client.Call("cache", "Get", request, rsp)
}

func (t *CacheService) Increment(request CacheIncrementRequest) (*CacheIncrementResponse, error) {
	rsp := &CacheIncrementResponse{}
	return rsp, t.client.Call("cache", "Increment", request, rsp)
}

func (t *CacheService) Set(request CacheSetRequest) (*CacheSetResponse, error) {
	rsp := &CacheSetResponse{}
	return rsp, t.client.Call("cache", "Set", request, rsp)
}

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

func (t *CryptoService) History(request CryptoHistoryRequest) (*CryptoHistoryResponse, error) {
	rsp := &CryptoHistoryResponse{}
	return rsp, t.client.Call("crypto", "History", request, rsp)
}

func (t *CryptoService) News(request CryptoNewsRequest) (*CryptoNewsResponse, error) {
	rsp := &CryptoNewsResponse{}
	return rsp, t.client.Call("crypto", "News", request, rsp)
}

func (t *CryptoService) Price(request CryptoPriceRequest) (*CryptoPriceResponse, error) {
	rsp := &CryptoPriceResponse{}
	return rsp, t.client.Call("crypto", "Price", request, rsp)
}

func (t *CryptoService) Quote(request CryptoQuoteRequest) (*CryptoQuoteResponse, error) {
	rsp := &CryptoQuoteResponse{}
	return rsp, t.client.Call("crypto", "Quote", request, rsp)
}

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

func (t *CurrencyService) Codes(request CurrencyCodesRequest) (*CurrencyCodesResponse, error) {
	rsp := &CurrencyCodesResponse{}
	return rsp, t.client.Call("currency", "Codes", request, rsp)
}

func (t *CurrencyService) Convert(request CurrencyConvertRequest) (*CurrencyConvertResponse, error) {
	rsp := &CurrencyConvertResponse{}
	return rsp, t.client.Call("currency", "Convert", request, rsp)
}

func (t *CurrencyService) History(request CurrencyHistoryRequest) (*CurrencyHistoryResponse, error) {
	rsp := &CurrencyHistoryResponse{}
	return rsp, t.client.Call("currency", "History", request, rsp)
}

func (t *CurrencyService) Rates(request CurrencyRatesRequest) (*CurrencyRatesResponse, error) {
	rsp := &CurrencyRatesResponse{}
	return rsp, t.client.Call("currency", "Rates", request, rsp)
}

func NewDbService(token string) *DbService {
	return &DbService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type DbService struct {
	client *client.Client
}

func (t *DbService) Create(request DbCreateRequest) (*DbCreateResponse, error) {
	rsp := &DbCreateResponse{}
	return rsp, t.client.Call("db", "Create", request, rsp)
}

func (t *DbService) Delete(request DbDeleteRequest) (*DbDeleteResponse, error) {
	rsp := &DbDeleteResponse{}
	return rsp, t.client.Call("db", "Delete", request, rsp)
}

func (t *DbService) Read(request DbReadRequest) (*DbReadResponse, error) {
	rsp := &DbReadResponse{}
	return rsp, t.client.Call("db", "Read", request, rsp)
}

func (t *DbService) Truncate(request DbTruncateRequest) (*DbTruncateResponse, error) {
	rsp := &DbTruncateResponse{}
	return rsp, t.client.Call("db", "Truncate", request, rsp)
}

func (t *DbService) Update(request DbUpdateRequest) (*DbUpdateResponse, error) {
	rsp := &DbUpdateResponse{}
	return rsp, t.client.Call("db", "Update", request, rsp)
}

func NewEmailService(token string) *EmailService {
	return &EmailService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type EmailService struct {
	client *client.Client
}

func (t *EmailService) Send(request EmailSendRequest) (*EmailSendResponse, error) {
	rsp := &EmailSendResponse{}
	return rsp, t.client.Call("email", "Send", request, rsp)
}

func NewEmojiService(token string) *EmojiService {
	return &EmojiService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type EmojiService struct {
	client *client.Client
}

func (t *EmojiService) Find(request EmojiFindRequest) (*EmojiFindResponse, error) {
	rsp := &EmojiFindResponse{}
	return rsp, t.client.Call("emoji", "Find", request, rsp)
}

func (t *EmojiService) Flag(request EmojiFlagRequest) (*EmojiFlagResponse, error) {
	rsp := &EmojiFlagResponse{}
	return rsp, t.client.Call("emoji", "Flag", request, rsp)
}

func (t *EmojiService) Print(request EmojiPrintRequest) (*EmojiPrintResponse, error) {
	rsp := &EmojiPrintResponse{}
	return rsp, t.client.Call("emoji", "Print", request, rsp)
}

func (t *EmojiService) Send(request EmojiSendRequest) (*EmojiSendResponse, error) {
	rsp := &EmojiSendResponse{}
	return rsp, t.client.Call("emoji", "Send", request, rsp)
}

func NewFileService(token string) *FileService {
	return &FileService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type FileService struct {
	client *client.Client
}

func (t *FileService) List(request FileListRequest) (*FileListResponse, error) {
	rsp := &FileListResponse{}
	return rsp, t.client.Call("file", "List", request, rsp)
}

func (t *FileService) Save(request FileSaveRequest) (*FileSaveResponse, error) {
	rsp := &FileSaveResponse{}
	return rsp, t.client.Call("file", "Save", request, rsp)
}

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

func (t *ForexService) History(request ForexHistoryRequest) (*ForexHistoryResponse, error) {
	rsp := &ForexHistoryResponse{}
	return rsp, t.client.Call("forex", "History", request, rsp)
}

func (t *ForexService) Price(request ForexPriceRequest) (*ForexPriceResponse, error) {
	rsp := &ForexPriceResponse{}
	return rsp, t.client.Call("forex", "Price", request, rsp)
}

func (t *ForexService) Quote(request ForexQuoteRequest) (*ForexQuoteResponse, error) {
	rsp := &ForexQuoteResponse{}
	return rsp, t.client.Call("forex", "Quote", request, rsp)
}

func NewGeocodingService(token string) *GeocodingService {
	return &GeocodingService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type GeocodingService struct {
	client *client.Client
}

func (t *GeocodingService) Lookup(request GeocodingLookupRequest) (*GeocodingLookupResponse, error) {
	rsp := &GeocodingLookupResponse{}
	return rsp, t.client.Call("geocoding", "Lookup", request, rsp)
}

func (t *GeocodingService) Reverse(request GeocodingReverseRequest) (*GeocodingReverseResponse, error) {
	rsp := &GeocodingReverseResponse{}
	return rsp, t.client.Call("geocoding", "Reverse", request, rsp)
}

func NewHelloworldService(token string) *HelloworldService {
	return &HelloworldService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type HelloworldService struct {
	client *client.Client
}

func (t *HelloworldService) Call(request HelloworldCallRequest) (*HelloworldCallResponse, error) {
	rsp := &HelloworldCallResponse{}
	return rsp, t.client.Call("helloworld", "Call", request, rsp)
}

func (t *HelloworldService) Stream(request HelloworldStreamRequest) (*HelloworldStreamResponse, error) {
	rsp := &HelloworldStreamResponse{}
	return rsp, t.client.Call("helloworld", "Stream", request, rsp)
}

func NewIdService(token string) *IdService {
	return &IdService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type IdService struct {
	client *client.Client
}

func (t *IdService) Generate(request IdGenerateRequest) (*IdGenerateResponse, error) {
	rsp := &IdGenerateResponse{}
	return rsp, t.client.Call("id", "Generate", request, rsp)
}

func (t *IdService) Types(request IdTypesRequest) (*IdTypesResponse, error) {
	rsp := &IdTypesResponse{}
	return rsp, t.client.Call("id", "Types", request, rsp)
}

func NewImageService(token string) *ImageService {
	return &ImageService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type ImageService struct {
	client *client.Client
}

func (t *ImageService) Convert(request ImageConvertRequest) (*ImageConvertResponse, error) {
	rsp := &ImageConvertResponse{}
	return rsp, t.client.Call("image", "Convert", request, rsp)
}

func (t *ImageService) Resize(request ImageResizeRequest) (*ImageResizeResponse, error) {
	rsp := &ImageResizeResponse{}
	return rsp, t.client.Call("image", "Resize", request, rsp)
}

func (t *ImageService) Upload(request ImageUploadRequest) (*ImageUploadResponse, error) {
	rsp := &ImageUploadResponse{}
	return rsp, t.client.Call("image", "Upload", request, rsp)
}

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

func (t *IpService) Lookup(request IpLookupRequest) (*IpLookupResponse, error) {
	rsp := &IpLookupResponse{}
	return rsp, t.client.Call("ip", "Lookup", request, rsp)
}

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

func (t *LocationService) Read(request LocationReadRequest) (*LocationReadResponse, error) {
	rsp := &LocationReadResponse{}
	return rsp, t.client.Call("location", "Read", request, rsp)
}

func (t *LocationService) Save(request LocationSaveRequest) (*LocationSaveResponse, error) {
	rsp := &LocationSaveResponse{}
	return rsp, t.client.Call("location", "Save", request, rsp)
}

func (t *LocationService) Search(request LocationSearchRequest) (*LocationSearchResponse, error) {
	rsp := &LocationSearchResponse{}
	return rsp, t.client.Call("location", "Search", request, rsp)
}

func NewOtpService(token string) *OtpService {
	return &OtpService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type OtpService struct {
	client *client.Client
}

func (t *OtpService) Generate(request OtpGenerateRequest) (*OtpGenerateResponse, error) {
	rsp := &OtpGenerateResponse{}
	return rsp, t.client.Call("otp", "Generate", request, rsp)
}

func (t *OtpService) Validate(request OtpValidateRequest) (*OtpValidateResponse, error) {
	rsp := &OtpValidateResponse{}
	return rsp, t.client.Call("otp", "Validate", request, rsp)
}

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

func (t *RoutingService) Directions(request RoutingDirectionsRequest) (*RoutingDirectionsResponse, error) {
	rsp := &RoutingDirectionsResponse{}
	return rsp, t.client.Call("routing", "Directions", request, rsp)
}

func (t *RoutingService) Eta(request RoutingEtaRequest) (*RoutingEtaResponse, error) {
	rsp := &RoutingEtaResponse{}
	return rsp, t.client.Call("routing", "Eta", request, rsp)
}

func (t *RoutingService) Route(request RoutingRouteRequest) (*RoutingRouteResponse, error) {
	rsp := &RoutingRouteResponse{}
	return rsp, t.client.Call("routing", "Route", request, rsp)
}

func NewRssService(token string) *RssService {
	return &RssService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type RssService struct {
	client *client.Client
}

func (t *RssService) Add(request RssAddRequest) (*RssAddResponse, error) {
	rsp := &RssAddResponse{}
	return rsp, t.client.Call("rss", "Add", request, rsp)
}

func (t *RssService) Feed(request RssFeedRequest) (*RssFeedResponse, error) {
	rsp := &RssFeedResponse{}
	return rsp, t.client.Call("rss", "Feed", request, rsp)
}

func (t *RssService) List(request RssListRequest) (*RssListResponse, error) {
	rsp := &RssListResponse{}
	return rsp, t.client.Call("rss", "List", request, rsp)
}

func (t *RssService) Remove(request RssRemoveRequest) (*RssRemoveResponse, error) {
	rsp := &RssRemoveResponse{}
	return rsp, t.client.Call("rss", "Remove", request, rsp)
}

func NewSentimentService(token string) *SentimentService {
	return &SentimentService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type SentimentService struct {
	client *client.Client
}

func (t *SentimentService) Analyze(request SentimentAnalyzeRequest) (*SentimentAnalyzeResponse, error) {
	rsp := &SentimentAnalyzeResponse{}
	return rsp, t.client.Call("sentiment", "Analyze", request, rsp)
}

func NewSmsService(token string) *SmsService {
	return &SmsService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type SmsService struct {
	client *client.Client
}

func (t *SmsService) Send(request SmsSendRequest) (*SmsSendResponse, error) {
	rsp := &SmsSendResponse{}
	return rsp, t.client.Call("sms", "Send", request, rsp)
}

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

func (t *StockService) History(request StockHistoryRequest) (*StockHistoryResponse, error) {
	rsp := &StockHistoryResponse{}
	return rsp, t.client.Call("stock", "History", request, rsp)
}

func (t *StockService) OrderBook(request StockOrderBookRequest) (*StockOrderBookResponse, error) {
	rsp := &StockOrderBookResponse{}
	return rsp, t.client.Call("stock", "OrderBook", request, rsp)
}

func (t *StockService) Price(request StockPriceRequest) (*StockPriceResponse, error) {
	rsp := &StockPriceResponse{}
	return rsp, t.client.Call("stock", "Price", request, rsp)
}

func (t *StockService) Quote(request StockQuoteRequest) (*StockQuoteResponse, error) {
	rsp := &StockQuoteResponse{}
	return rsp, t.client.Call("stock", "Quote", request, rsp)
}

func NewStreamService(token string) *StreamService {
	return &StreamService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type StreamService struct {
	client *client.Client
}

func (t *StreamService) Publish(request StreamPublishRequest) (*StreamPublishResponse, error) {
	rsp := &StreamPublishResponse{}
	return rsp, t.client.Call("stream", "Publish", request, rsp)
}

func (t *StreamService) Subscribe(request StreamSubscribeRequest) (*StreamSubscribeResponse, error) {
	rsp := &StreamSubscribeResponse{}
	return rsp, t.client.Call("stream", "Subscribe", request, rsp)
}

func NewThumbnailService(token string) *ThumbnailService {
	return &ThumbnailService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type ThumbnailService struct {
	client *client.Client
}

func (t *ThumbnailService) Screenshot(request ThumbnailScreenshotRequest) (*ThumbnailScreenshotResponse, error) {
	rsp := &ThumbnailScreenshotResponse{}
	return rsp, t.client.Call("thumbnail", "Screenshot", request, rsp)
}

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

func (t *TimeService) Now(request TimeNowRequest) (*TimeNowResponse, error) {
	rsp := &TimeNowResponse{}
	return rsp, t.client.Call("time", "Now", request, rsp)
}

func (t *TimeService) Zone(request TimeZoneRequest) (*TimeZoneResponse, error) {
	rsp := &TimeZoneResponse{}
	return rsp, t.client.Call("time", "Zone", request, rsp)
}

func NewUrlService(token string) *UrlService {
	return &UrlService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type UrlService struct {
	client *client.Client
}

func (t *UrlService) List(request UrlListRequest) (*UrlListResponse, error) {
	rsp := &UrlListResponse{}
	return rsp, t.client.Call("url", "List", request, rsp)
}

func (t *UrlService) Proxy(request UrlProxyRequest) (*UrlProxyResponse, error) {
	rsp := &UrlProxyResponse{}
	return rsp, t.client.Call("url", "Proxy", request, rsp)
}

func (t *UrlService) Shorten(request UrlShortenRequest) (*UrlShortenResponse, error) {
	rsp := &UrlShortenResponse{}
	return rsp, t.client.Call("url", "Shorten", request, rsp)
}

func NewUserService(token string) *UserService {
	return &UserService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type UserService struct {
	client *client.Client
}

func (t *UserService) Create(request UserCreateRequest) (*UserCreateResponse, error) {
	rsp := &UserCreateResponse{}
	return rsp, t.client.Call("user", "Create", request, rsp)
}

func (t *UserService) Delete(request UserDeleteRequest) (*UserDeleteResponse, error) {
	rsp := &UserDeleteResponse{}
	return rsp, t.client.Call("user", "Delete", request, rsp)
}

func (t *UserService) Login(request UserLoginRequest) (*UserLoginResponse, error) {
	rsp := &UserLoginResponse{}
	return rsp, t.client.Call("user", "Login", request, rsp)
}

func (t *UserService) Logout(request UserLogoutRequest) (*UserLogoutResponse, error) {
	rsp := &UserLogoutResponse{}
	return rsp, t.client.Call("user", "Logout", request, rsp)
}

func (t *UserService) Read(request UserReadRequest) (*UserReadResponse, error) {
	rsp := &UserReadResponse{}
	return rsp, t.client.Call("user", "Read", request, rsp)
}

func (t *UserService) ReadSession(request UserReadSessionRequest) (*UserReadSessionResponse, error) {
	rsp := &UserReadSessionResponse{}
	return rsp, t.client.Call("user", "ReadSession", request, rsp)
}

func (t *UserService) SendVerificationEmail(request UserSendVerificationEmailRequest) (*UserSendVerificationEmailResponse, error) {
	rsp := &UserSendVerificationEmailResponse{}
	return rsp, t.client.Call("user", "SendVerificationEmail", request, rsp)
}

func (t *UserService) UpdatePassword(request UserUpdatePasswordRequest) (*UserUpdatePasswordResponse, error) {
	rsp := &UserUpdatePasswordResponse{}
	return rsp, t.client.Call("user", "UpdatePassword", request, rsp)
}

func (t *UserService) Update(request UserUpdateRequest) (*UserUpdateResponse, error) {
	rsp := &UserUpdateResponse{}
	return rsp, t.client.Call("user", "Update", request, rsp)
}

func (t *UserService) VerifyEmail(request UserVerifyEmailRequest) (*UserVerifyEmailResponse, error) {
	rsp := &UserVerifyEmailResponse{}
	return rsp, t.client.Call("user", "VerifyEmail", request, rsp)
}

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

func (t *WeatherService) Forecast(request WeatherForecastRequest) (*WeatherForecastResponse, error) {
	rsp := &WeatherForecastResponse{}
	return rsp, t.client.Call("weather", "Forecast", request, rsp)
}

func (t *WeatherService) Now(request WeatherNowRequest) (*WeatherNowResponse, error) {
	rsp := &WeatherNowResponse{}
	return rsp, t.client.Call("weather", "Now", request, rsp)
}

type AnswerQuestionRequest struct {
	Query string `json:"query"`
}

type AnswerQuestionResponse struct {
	Answer string `json:"answer"`
	Image  string `json:"image"`
	Url    string `json:"url"`
}

type CacheDecrementRequest struct {
	Key   string `json:"key"`
	Value int64  `json:"value"`
}

type CacheDecrementResponse struct {
	Key   string `json:"key"`
	Value int64  `json:"value"`
}

type CacheDeleteRequest struct {
	Key string `json:"key"`
}

type CacheDeleteResponse struct {
	Status string `json:"status"`
}

type CacheGetRequest struct {
	Key string `json:"key"`
}

type CacheGetResponse struct {
	Key   string `json:"key"`
	Ttl   int64  `json:"ttl"`
	Value string `json:"value"`
}

type CacheIncrementRequest struct {
	Key   string `json:"key"`
	Value int64  `json:"value"`
}

type CacheIncrementResponse struct {
	Key   string `json:"key"`
	Value int64  `json:"value"`
}

type CacheSetRequest struct {
	Key   string `json:"key"`
	Ttl   int64  `json:"ttl"`
	Value string `json:"value"`
}

type CacheSetResponse struct {
	Status string `json:"status"`
}

type CryptoArticle struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	Source      string `json:"source"`
	Title       string `json:"title"`
	Url         string `json:"url"`
}

type CryptoHistoryRequest struct {
	Symbol string `json:"symbol"`
}

type CryptoHistoryResponse struct {
	Close  float64 `json:"close"`
	Date   string  `json:"date"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Open   float64 `json:"open"`
	Symbol string  `json:"symbol"`
	Volume float64 `json:"volume"`
}

type CryptoNewsRequest struct {
	Symbol string `json:"symbol"`
}

type CryptoNewsResponse struct {
	Articles []CryptoArticle `json:"articles"`
	Symbol   string          `json:"symbol"`
}

type CryptoPriceRequest struct {
	Symbol string `json:"symbol"`
}

type CryptoPriceResponse struct {
	Price  float64 `json:"price"`
	Symbol string  `json:"symbol"`
}

type CryptoQuoteRequest struct {
	Symbol string `json:"symbol"`
}

type CryptoQuoteResponse struct {
	AskPrice  float64 `json:"askPrice"`
	AskSize   float64 `json:"askSize"`
	BidPrice  float64 `json:"bidPrice"`
	BidSize   float64 `json:"bidSize"`
	Symbol    string  `json:"symbol"`
	Timestamp string  `json:"timestamp"`
}

type CurrencyCode struct {
	Currency string `json:"currency"`
	Name     string `json:"name"`
}

type CurrencyCodesRequest struct {
}

type CurrencyCodesResponse struct {
	Codes []CurrencyCode `json:"codes"`
}

type CurrencyConvertRequest struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}

type CurrencyConvertResponse struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	Rate   float64 `json:"rate"`
	To     string  `json:"to"`
}

type CurrencyHistoryRequest struct {
	Code string `json:"code"`
	Date string `json:"date"`
}

type CurrencyHistoryResponse struct {
	Code  string                 `json:"code"`
	Date  string                 `json:"date"`
	Rates map[string]interface{} `json:"rates"`
}

type CurrencyRatesRequest struct {
	Code string `json:"code"`
}

type CurrencyRatesResponse struct {
	Code  string                 `json:"code"`
	Rates map[string]interface{} `json:"rates"`
}

type DbCreateRequest struct {
	Record map[string]interface{} `json:"record"`
	Table  string                 `json:"table"`
}

type DbCreateResponse struct {
	Id string `json:"id"`
}

type DbDeleteRequest struct {
	Id    string `json:"id"`
	Table string `json:"table"`
}

type DbDeleteResponse struct {
}

type DbReadRequest struct {
	Id      string `json:"id"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
	Order   string `json:"order"`
	OrderBy string `json:"orderBy"`
	Query   string `json:"query"`
	Table   string `json:"table"`
}

type DbReadResponse struct {
	Records map[string]interface{} `json:"records"`
}

type DbTruncateRequest struct {
	Table string `json:"table"`
}

type DbTruncateResponse struct {
	Table string `json:"table"`
}

type DbUpdateRequest struct {
	Id     string                 `json:"id"`
	Record map[string]interface{} `json:"record"`
	Table  string                 `json:"table"`
}

type DbUpdateResponse struct {
}

type EmailSendRequest struct {
	From     string `json:"from"`
	HtmlBody string `json:"htmlBody"`
	ReplyTo  string `json:"replyTo"`
	Subject  string `json:"subject"`
	TextBody string `json:"textBody"`
	To       string `json:"to"`
}

type EmailSendResponse struct {
}

type EmojiFindRequest struct {
	Alias string `json:"alias"`
}

type EmojiFindResponse struct {
	Emoji string `json:"emoji"`
}

type EmojiFlagRequest struct {
	Code string `json:"code"`
}

type EmojiFlagResponse struct {
	Flag string `json:"flag"`
}

type EmojiPrintRequest struct {
	Text string `json:"text"`
}

type EmojiPrintResponse struct {
	Text string `json:"text"`
}

type EmojiSendRequest struct {
	From    string `json:"from"`
	Message string `json:"message"`
	To      string `json:"to"`
}

type EmojiSendResponse struct {
	Success bool `json:"success"`
}

type FileFile struct {
	Created      int64  `json:"created"`
	FileContents string `json:"fileContents"`
	IsDirectory  bool   `json:"isDirectory"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Project      string `json:"project"`
	Updated      int64  `json:"updated"`
}

type FileListRequest struct {
	Path    string `json:"path"`
	Project string `json:"project"`
}

type FileListResponse struct {
	File []FileFile `json:"file"`
}

type FileSaveRequest struct {
	File []FileFile `json:"file"`
}

type FileSaveResponse struct {
}

type ForexHistoryRequest struct {
	Symbol string `json:"symbol"`
}

type ForexHistoryResponse struct {
	Close  float64 `json:"close"`
	Date   string  `json:"date"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Open   float64 `json:"open"`
	Symbol string  `json:"symbol"`
	Volume float64 `json:"volume"`
}

type ForexPriceRequest struct {
	Symbol string `json:"symbol"`
}

type ForexPriceResponse struct {
	Price  float64 `json:"price"`
	Symbol string  `json:"symbol"`
}

type ForexQuoteRequest struct {
	Symbol string `json:"symbol"`
}

type ForexQuoteResponse struct {
	AskPrice  float64 `json:"askPrice"`
	BidPrice  float64 `json:"bidPrice"`
	Symbol    string  `json:"symbol"`
	Timestamp string  `json:"timestamp"`
}

type GeocodingAddress struct {
	City     string `json:"city"`
	Country  string `json:"country"`
	LineOne  string `json:"lineOne"`
	LineTwo  string `json:"lineTwo"`
	Postcode string `json:"postcode"`
}

type GeocodingLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeocodingLookupRequest struct {
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Postcode string `json:"postcode"`
}

type GeocodingLookupResponse struct {
	Address  GeocodingAddress  `json:"address"`
	Location GeocodingLocation `json:"location"`
}

type GeocodingReverseRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeocodingReverseResponse struct {
	Address  GeocodingAddress  `json:"address"`
	Location GeocodingLocation `json:"location"`
}

type HelloworldCallRequest struct {
	Name string `json:"name"`
}

type HelloworldCallResponse struct {
	Message string `json:"message"`
}

type HelloworldStreamRequest struct {
	Messages int64  `json:"messages"`
	Name     string `json:"name"`
}

type HelloworldStreamResponse struct {
	Message string `json:"message"`
}

type IdGenerateRequest struct {
	Type string `json:"type"`
}

type IdGenerateResponse struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type IdTypesRequest struct {
}

type IdTypesResponse struct {
	Types []string `json:"types"`
}

type ImageConvertRequest struct {
	Base64    string `json:"base64"`
	Name      string `json:"name"`
	OutputUrl bool   `json:"outputUrl"`
	Url       string `json:"url"`
}

type ImageConvertResponse struct {
	Base64 string `json:"base64"`
	Url    string `json:"url"`
}

type ImageCropOptions struct {
	Anchor string `json:"anchor"`
	Height int32  `json:"height"`
	Width  int32  `json:"width"`
}

type ImagePoint struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type ImageRectangle struct {
	Max ImagePoint `json:"max"`
	Min ImagePoint `json:"min"`
}

type ImageResizeRequest struct {
	Base64      string           `json:"base64"`
	CropOptions ImageCropOptions `json:"cropOptions"`
	Height      int64            `json:"height"`
	Name        string           `json:"name"`
	OutputUrl   bool             `json:"outputUrl"`
	Url         string           `json:"url"`
	Width       int64            `json:"width"`
}

type ImageResizeResponse struct {
	Base64 string `json:"base64"`
	Url    string `json:"url"`
}

type ImageUploadRequest struct {
	Base64 string `json:"base64"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}

type ImageUploadResponse struct {
	Url string `json:"url"`
}

type IpLookupRequest struct {
	Ip string `json:"ip"`
}

type IpLookupResponse struct {
	Asn       int32   `json:"asn"`
	City      string  `json:"city"`
	Continent string  `json:"continent"`
	Country   string  `json:"country"`
	Ip        string  `json:"ip"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
}

type LocationEntity struct {
	Id       string        `json:"id"`
	Location LocationPoint `json:"location"`
	Type     string        `json:"type"`
}

type LocationPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

type LocationReadRequest struct {
	Id string `json:"id"`
}

type LocationReadResponse struct {
	Entity LocationEntity `json:"entity"`
}

type LocationSaveRequest struct {
	Entity LocationEntity `json:"entity"`
}

type LocationSaveResponse struct {
}

type LocationSearchRequest struct {
	Center      LocationPoint `json:"center"`
	NumEntities int64         `json:"numEntities"`
	Radius      float64       `json:"radius"`
	Type        string        `json:"type"`
}

type LocationSearchResponse struct {
	Entities []LocationEntity `json:"entities"`
}

type OtpGenerateRequest struct {
	Expiry int64  `json:"expiry"`
	Id     string `json:"id"`
	Size   int64  `json:"size"`
}

type OtpGenerateResponse struct {
	Code string `json:"code"`
}

type OtpValidateRequest struct {
	Code string `json:"code"`
	Id   string `json:"id"`
}

type OtpValidateResponse struct {
	Success bool `json:"success"`
}

type RoutingDirection struct {
	Distance      float64               `json:"distance"`
	Duration      float64               `json:"duration"`
	Instruction   string                `json:"instruction"`
	Intersections []RoutingIntersection `json:"intersections"`
	Maneuver      RoutingManeuver       `json:"maneuver"`
	Name          string                `json:"name"`
	Reference     string                `json:"reference"`
}

type RoutingDirectionsRequest struct {
	Destination RoutingPoint `json:"destination"`
	Origin      RoutingPoint `json:"origin"`
}

type RoutingDirectionsResponse struct {
	Directions []RoutingDirection `json:"directions"`
	Distance   float64            `json:"distance"`
	Duration   float64            `json:"duration"`
	Waypoints  []RoutingWaypoint  `json:"waypoints"`
}

type RoutingEtaRequest struct {
	Destination RoutingPoint `json:"destination"`
	Origin      RoutingPoint `json:"origin"`
	Speed       float64      `json:"speed"`
	Type        string       `json:"type"`
}

type RoutingEtaResponse struct {
	Duration float64 `json:"duration"`
}

type RoutingIntersection struct {
	Bearings []float64    `json:"bearings"`
	Location RoutingPoint `json:"location"`
}

type RoutingManeuver struct {
	Action        string       `json:"action"`
	BearingAfter  float64      `json:"bearingAfter"`
	BearingBefore float64      `json:"bearingBefore"`
	Direction     string       `json:"direction"`
	Location      RoutingPoint `json:"location"`
}

type RoutingPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RoutingRouteRequest struct {
	Destination RoutingPoint `json:"destination"`
	Origin      RoutingPoint `json:"origin"`
}

type RoutingRouteResponse struct {
	Distance  float64           `json:"distance"`
	Duration  float64           `json:"duration"`
	Waypoints []RoutingWaypoint `json:"waypoints"`
}

type RoutingWaypoint struct {
	Location RoutingPoint `json:"location"`
	Name     string       `json:"name"`
}

type RssAddRequest struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	Url      string `json:"url"`
}

type RssAddResponse struct {
}

type RssEntry struct {
	Content string `json:"content"`
	Date    string `json:"date"`
	Feed    string `json:"feed"`
	Id      string `json:"id"`
	Link    string `json:"link"`
	Summary string `json:"summary"`
	Title   string `json:"title"`
}

type RssFeed struct {
	Category string `json:"category"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
}

type RssFeedRequest struct {
	Limit  int64  `json:"limit"`
	Name   string `json:"name"`
	Offset int64  `json:"offset"`
}

type RssFeedResponse struct {
	Entries []RssEntry `json:"entries"`
}

type RssListRequest struct {
}

type RssListResponse struct {
	Feeds []RssFeed `json:"feeds"`
}

type RssRemoveRequest struct {
	Name string `json:"name"`
}

type RssRemoveResponse struct {
}

type SentimentAnalyzeRequest struct {
	Lang string `json:"lang"`
	Text string `json:"text"`
}

type SentimentAnalyzeResponse struct {
	Score float64 `json:"score"`
}

type SmsSendRequest struct {
	From    string `json:"from"`
	Message string `json:"message"`
	To      string `json:"to"`
}

type SmsSendResponse struct {
	Info   string `json:"info"`
	Status string `json:"status"`
}

type StockHistoryRequest struct {
	Date  string `json:"date"`
	Stock string `json:"stock"`
}

type StockHistoryResponse struct {
	Close  float64 `json:"close"`
	Date   string  `json:"date"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Open   float64 `json:"open"`
	Symbol string  `json:"symbol"`
	Volume int32   `json:"volume"`
}

type StockOrder struct {
	AskPrice  float64 `json:"askPrice"`
	AskSize   int32   `json:"askSize"`
	BidPrice  float64 `json:"bidPrice"`
	BidSize   int32   `json:"bidSize"`
	Timestamp string  `json:"timestamp"`
}

type StockOrderBookRequest struct {
	Date  string `json:"date"`
	End   string `json:"end"`
	Limit int32  `json:"limit"`
	Start string `json:"start"`
	Stock string `json:"stock"`
}

type StockOrderBookResponse struct {
	Date   string       `json:"date"`
	Orders []StockOrder `json:"orders"`
	Symbol string       `json:"symbol"`
}

type StockPriceRequest struct {
	Symbol string `json:"symbol"`
}

type StockPriceResponse struct {
	Price  float64 `json:"price"`
	Symbol string  `json:"symbol"`
}

type StockQuoteRequest struct {
	Symbol string `json:"symbol"`
}

type StockQuoteResponse struct {
	AskPrice  float64 `json:"askPrice"`
	AskSize   int32   `json:"askSize"`
	BidPrice  float64 `json:"bidPrice"`
	BidSize   int32   `json:"bidSize"`
	Symbol    string  `json:"symbol"`
	Timestamp string  `json:"timestamp"`
}

type StreamPublishRequest struct {
	Message map[string]interface{} `json:"message"`
	Topic   string                 `json:"topic"`
}

type StreamPublishResponse struct {
}

type StreamSubscribeRequest struct {
	Topic string `json:"topic"`
}

type StreamSubscribeResponse struct {
	Message map[string]interface{} `json:"message"`
	Topic   string                 `json:"topic"`
}

type ThumbnailScreenshotRequest struct {
	Height int32  `json:"height"`
	Url    string `json:"url"`
	Width  int32  `json:"width"`
}

type ThumbnailScreenshotResponse struct {
	ImageUrl string `json:"imageUrl"`
}

type TimeNowRequest struct {
	Location string `json:"location"`
}

type TimeNowResponse struct {
	Localtime string `json:"localtime"`
	Location  string `json:"location"`
	Timestamp string `json:"timestamp"`
	Timezone  string `json:"timezone"`
	Unix      int64  `json:"unix"`
}

type TimeZoneRequest struct {
	Location string `json:"location"`
}

type TimeZoneResponse struct {
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

type UrlListRequest struct {
	ShortUrl string `json:"shortUrl"`
}

type UrlListResponse struct {
	UrlPairs UrlURLPair `json:"urlPairs"`
}

type UrlProxyRequest struct {
	ShortUrl string `json:"shortUrl"`
}

type UrlProxyResponse struct {
	DestinationUrl string `json:"destinationUrl"`
}

type UrlShortenRequest struct {
	DestinationUrl string `json:"destinationUrl"`
}

type UrlShortenResponse struct {
	ShortUrl string `json:"shortUrl"`
}

type UrlURLPair struct {
	Created        int64  `json:"created"`
	DestinationUrl string `json:"destinationUrl"`
	HitCount       int64  `json:"hitCount"`
	Owner          string `json:"owner"`
	ShortUrl       string `json:"shortUrl"`
}

type UserAccount struct {
	Created          int64                  `json:"created"`
	Email            string                 `json:"email"`
	Id               string                 `json:"id"`
	Profile          map[string]interface{} `json:"profile"`
	Updated          int64                  `json:"updated"`
	Username         string                 `json:"username"`
	VerificationDate int64                  `json:"verificationDate"`
	Verified         bool                   `json:"verified"`
}

type UserCreateRequest struct {
	Email    string                 `json:"email"`
	Id       string                 `json:"id"`
	Password string                 `json:"password"`
	Profile  map[string]interface{} `json:"profile"`
	Username string                 `json:"username"`
}

type UserCreateResponse struct {
	Account UserAccount `json:"account"`
}

type UserDeleteRequest struct {
	Id string `json:"id"`
}

type UserDeleteResponse struct {
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserLoginResponse struct {
	Session UserSession `json:"session"`
}

type UserLogoutRequest struct {
	SessionId string `json:"sessionId"`
}

type UserLogoutResponse struct {
}

type UserReadRequest struct {
	Email    string `json:"email"`
	Id       string `json:"id"`
	Username string `json:"username"`
}

type UserReadResponse struct {
	Account UserAccount `json:"account"`
}

type UserReadSessionRequest struct {
	SessionId string `json:"sessionId"`
}

type UserReadSessionResponse struct {
	Session UserSession `json:"session"`
}

type UserSendVerificationEmailRequest struct {
	Email              string `json:"email"`
	FailureRedirectUrl string `json:"failureRedirectUrl"`
	FromName           string `json:"fromName"`
	RedirectUrl        string `json:"redirectUrl"`
	Subject            string `json:"subject"`
	TextContent        string `json:"textContent"`
}

type UserSendVerificationEmailResponse struct {
}

type UserSession struct {
	Created int64  `json:"created"`
	Expires int64  `json:"expires"`
	Id      string `json:"id"`
	UserId  string `json:"userId"`
}

type UserUpdatePasswordRequest struct {
	ConfirmPassword string `json:"confirmPassword"`
	NewPassword     string `json:"newPassword"`
	OldPassword     string `json:"oldPassword"`
	UserId          string `json:"userId"`
}

type UserUpdatePasswordResponse struct {
}

type UserUpdateRequest struct {
	Email    string                 `json:"email"`
	Id       string                 `json:"id"`
	Profile  map[string]interface{} `json:"profile"`
	Username string                 `json:"username"`
}

type UserUpdateResponse struct {
}

type UserVerifyEmailRequest struct {
	Token string `json:"token"`
}

type UserVerifyEmailResponse struct {
}

type WeatherForecast struct {
	AvgTempC     float64 `json:"avgTempC"`
	AvgTempF     float64 `json:"avgTempF"`
	ChanceOfRain int32   `json:"chanceOfRain"`
	Condition    string  `json:"condition"`
	Date         string  `json:"date"`
	IconUrl      string  `json:"iconUrl"`
	MaxTempC     float64 `json:"maxTempC"`
	MaxTempF     float64 `json:"maxTempF"`
	MinTempC     float64 `json:"minTempC"`
	MinTempF     float64 `json:"minTempF"`
	Sunrise      string  `json:"sunrise"`
	Sunset       string  `json:"sunset"`
	WillItRain   bool    `json:"willItRain"`
}

type WeatherForecastRequest struct {
	Days     int32  `json:"days"`
	Location string `json:"location"`
}

type WeatherForecastResponse struct {
	Country   string            `json:"country"`
	Forecast  []WeatherForecast `json:"forecast"`
	Latitude  float64           `json:"latitude"`
	LocalTime string            `json:"localTime"`
	Location  string            `json:"location"`
	Longitude float64           `json:"longitude"`
	Region    string            `json:"region"`
	Timezone  string            `json:"timezone"`
}

type WeatherNowRequest struct {
	Location string `json:"location"`
}

type WeatherNowResponse struct {
	Cloud         int32   `json:"cloud"`
	Condition     string  `json:"condition"`
	Country       string  `json:"country"`
	Daytime       bool    `json:"daytime"`
	FeelsLikeC    float64 `json:"feelsLikeC"`
	FeelsLikeF    float64 `json:"feelsLikeF"`
	Humidity      int32   `json:"humidity"`
	IconUrl       string  `json:"iconUrl"`
	Latitude      float64 `json:"latitude"`
	LocalTime     string  `json:"localTime"`
	Location      string  `json:"location"`
	Longitude     float64 `json:"longitude"`
	Region        string  `json:"region"`
	TempC         float64 `json:"tempC"`
	TempF         float64 `json:"tempF"`
	Timezone      string  `json:"timezone"`
	WindDegree    int32   `json:"windDegree"`
	WindDirection string  `json:"windDirection"`
	WindKph       float64 `json:"windKph"`
	WindMph       float64 `json:"windMph"`
}
