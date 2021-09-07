package example

import (
	"fmt"
	"github.com/micro/micro-go/db"
)

func UpdateArecord() {
	dbService := db.NewDbService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := dbService.Update(db.UpdateRequest{
		Record: map[string]interface{}{
			"id":  "1",
			"age": 43,
		},
		Table: "users",
	})
	fmt.Println(rsp)
}
