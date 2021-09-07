package main

import (
	"fmt"
	"github.com/micro/micro-go/file"
)

func main() {
	fileService := file.NewFileService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := fileService.Delete(file.DeleteRequest{
		Path:    "/document/text-files/file.txt",
		Project: "examples",
	})
	fmt.Println(rsp)
}
