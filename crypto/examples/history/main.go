package main

import (
	"fmt"
	"github.com/micro/micro-go/crypto"
)

func main() {
	cryptoService := crypto.NewCryptoService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := cryptoService.History(crypto.HistoryRequest{
		Symbol: "BTCUSD",
	})
	fmt.Println(rsp)
}
