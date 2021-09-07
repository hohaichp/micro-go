package example

import (
	"fmt"
	"github.com/micro/micro-go/crypto"
)

// Returns the last quote for a currency including bid and ask prices
func GetAcryptocurrencyQuote() {
	cryptoService := crypto.NewCryptoService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := cryptoService.Quote(crypto.QuoteRequest{
		Symbol: "BTCUSD",
	})
	fmt.Println(rsp)
}
