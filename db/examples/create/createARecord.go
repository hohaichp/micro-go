package example

import (
	"fmt"
	"github.com/micro/micro-go/db"
)

func CreateArecord() {
	dbService := db.NewDbService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := dbService.Create(db.CreateRequest{
		Record: map[string]interface{}{
			"age":      42,
			"isActive": true,
			"id":       "1",
			"name":     "Jane",
		},
		Table: "users",
	})
	fmt.Println(rsp)
}
