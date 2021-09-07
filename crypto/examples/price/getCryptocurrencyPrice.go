package example

import (
	"fmt"
	"github.com/micro/micro-go/crypto"
)

// Returns the last traded price of a currency
func GetCryptocurrencyPrice() {
	cryptoService := crypto.NewCryptoService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := cryptoService.Price(crypto.PriceRequest{
		Symbol: "BTCUSD",
	})
	fmt.Println(rsp)
}
