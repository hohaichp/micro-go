package example

import (
	"fmt"
	"github.com/micro/micro-go/crypto"
)

// Returns the history for the previous close
func GetPreviousClose() {
	cryptoService := crypto.NewCryptoService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := cryptoService.History(crypto.HistoryRequest{
		Symbol: "BTCUSD",
	})
	fmt.Println(rsp)
}
