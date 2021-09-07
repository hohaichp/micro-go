package main

import (
	"fmt"
	"github.com/micro/micro-go/file"
)

func main() {
	fileService := file.NewFileService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := fileService.Read(file.ReadRequest{
		Path:    "/document/text-files/file.txt",
		Project: "examples",
	})
	fmt.Println(rsp)
}
