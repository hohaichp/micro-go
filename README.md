# Type safe m3o clients

```go
package main

import (
	"fmt"
	"os"

	mgo "github.com/micro/micro-go"
)

func main() {
	c := mgo.NewClient(os.Getenv("MICRO_TOKEN"))
	rsp, err := c.HelloworldService.Call(mgo.HelloworldCallRequest{
		Name: "Janos",
	})
	fmt.Println(rsp, err)
}
```

## Beta warning

This package is highly experimental. You might notice request and response types not matching the APIs found on m3o.
In that case let us know.

## Known issues

Map types are represented inaccurately, all maps are `map[string]interfaces{}` even if it's defined as `map[string]float64` or equivalent on m3o.com
This is being fixed.