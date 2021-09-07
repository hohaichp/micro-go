package micro

import (
	"github.com/micro/micro-go/address"
	"github.com/micro/micro-go/answer"
	"github.com/micro/micro-go/cache"
	"github.com/micro/micro-go/crypto"
	"github.com/micro/micro-go/currency"
	"github.com/micro/micro-go/db"
	"github.com/micro/micro-go/email"
	"github.com/micro/micro-go/emoji"
	"github.com/micro/micro-go/file"
	"github.com/micro/micro-go/forex"
	"github.com/micro/micro-go/geocoding"
	"github.com/micro/micro-go/helloworld"
	"github.com/micro/micro-go/id"
	"github.com/micro/micro-go/image"
	"github.com/micro/micro-go/ip"
	"github.com/micro/micro-go/location"
	"github.com/micro/micro-go/otp"
	"github.com/micro/micro-go/postcode"
	"github.com/micro/micro-go/routing"
	"github.com/micro/micro-go/rss"
	"github.com/micro/micro-go/sentiment"
	"github.com/micro/micro-go/sms"
	"github.com/micro/micro-go/stock"
	"github.com/micro/micro-go/stream"
	"github.com/micro/micro-go/thumbnail"
	"github.com/micro/micro-go/time"
	"github.com/micro/micro-go/url"
	"github.com/micro/micro-go/user"
	"github.com/micro/micro-go/weather"
)

func NewClient(token string) *Client {
	return &Client{
		token: token,

		AddressService:    address.NewAddressService(token),
		AnswerService:     answer.NewAnswerService(token),
		CacheService:      cache.NewCacheService(token),
		CryptoService:     crypto.NewCryptoService(token),
		CurrencyService:   currency.NewCurrencyService(token),
		DbService:         db.NewDbService(token),
		EmailService:      email.NewEmailService(token),
		EmojiService:      emoji.NewEmojiService(token),
		FileService:       file.NewFileService(token),
		ForexService:      forex.NewForexService(token),
		GeocodingService:  geocoding.NewGeocodingService(token),
		HelloworldService: helloworld.NewHelloworldService(token),
		IdService:         id.NewIdService(token),
		ImageService:      image.NewImageService(token),
		IpService:         ip.NewIpService(token),
		LocationService:   location.NewLocationService(token),
		OtpService:        otp.NewOtpService(token),
		PostcodeService:   postcode.NewPostcodeService(token),
		RoutingService:    routing.NewRoutingService(token),
		RssService:        rss.NewRssService(token),
		SentimentService:  sentiment.NewSentimentService(token),
		SmsService:        sms.NewSmsService(token),
		StockService:      stock.NewStockService(token),
		StreamService:     stream.NewStreamService(token),
		ThumbnailService:  thumbnail.NewThumbnailService(token),
		TimeService:       time.NewTimeService(token),
		UrlService:        url.NewUrlService(token),
		UserService:       user.NewUserService(token),
		WeatherService:    weather.NewWeatherService(token),
	}
}

type Client struct {
	token string

	AddressService    *address.AddressService
	AnswerService     *answer.AnswerService
	CacheService      *cache.CacheService
	CryptoService     *crypto.CryptoService
	CurrencyService   *currency.CurrencyService
	DbService         *db.DbService
	EmailService      *email.EmailService
	EmojiService      *emoji.EmojiService
	FileService       *file.FileService
	ForexService      *forex.ForexService
	GeocodingService  *geocoding.GeocodingService
	HelloworldService *helloworld.HelloworldService
	IdService         *id.IdService
	ImageService      *image.ImageService
	IpService         *ip.IpService
	LocationService   *location.LocationService
	OtpService        *otp.OtpService
	PostcodeService   *postcode.PostcodeService
	RoutingService    *routing.RoutingService
	RssService        *rss.RssService
	SentimentService  *sentiment.SentimentService
	SmsService        *sms.SmsService
	StockService      *stock.StockService
	StreamService     *stream.StreamService
	ThumbnailService  *thumbnail.ThumbnailService
	TimeService       *time.TimeService
	UrlService        *url.UrlService
	UserService       *user.UserService
	WeatherService    *weather.WeatherService
}
