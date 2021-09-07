package main

import (
	"fmt"
	"github.com/micro/micro-go/ip"
)

func main() {
	ipService := ip.NewIpService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := ipService.Lookup(ip.LookupRequest{
		Ip: "93.148.214.31",
	})
	fmt.Println(rsp)
}
