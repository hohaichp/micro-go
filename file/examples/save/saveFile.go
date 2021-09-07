package example

import (
	"fmt"
	"github.com/micro/micro-go/file"
)

func SaveFile() {
	fileService := file.NewFileService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := fileService.Save(file.SaveRequest{
		File: file.Record{
			Content: "file content example",
			Path:    "/document/text-files/file.txt",
			Project: "examples",
		},
	})
	fmt.Println(rsp)
}
