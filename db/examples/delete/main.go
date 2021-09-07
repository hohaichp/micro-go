package main

import (
	"fmt"
	"github.com/micro/micro-go/db"
)

func main() {
	dbService := db.NewDbService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := dbService.Delete(db.DeleteRequest{
		Id:    "1",
		Table: "users",
	})
	fmt.Println(rsp)
}