package example

import (
	"fmt"
	"github.com/micro/micro-go/file"
)

func ListFiles() {
	fileService := file.NewFileService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := fileService.List(file.ListRequest{
		Project: "examples",
	})
	fmt.Println(rsp)
}
