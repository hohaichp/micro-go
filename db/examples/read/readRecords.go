package example

import (
	"fmt"
	"github.com/micro/micro-go/db"
)

func ReadRecords() {
	dbService := db.NewDbService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := dbService.Read(db.ReadRequest{
		Query: "age == 43",
		Table: "users",
	})
	fmt.Println(rsp)
}
