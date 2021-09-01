# M3O Go Client

A typesafe generated Go client for M3O services

```go
package main

import (
	"fmt"
	"os"

	"github.com/micro/micro-go"
)

func main() {
	c := micro.NewClient(os.Getenv("MICRO_TOKEN"))
	rsp, err := c.HelloworldService.Call(micro.HelloworldCallRequest{
		Name: "Janos",
	})
	fmt.Println(rsp, err)
}
```

## Beta warning

This package is highly experimental. You might notice request and response types not matching the APIs found on m3o.
In that case let us know.
