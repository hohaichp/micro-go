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