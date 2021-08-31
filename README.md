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