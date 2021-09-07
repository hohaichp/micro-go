package main

import (
	"fmt"
	"github.com/micro/micro-go/sentiment"
)

func main() {
	sentimentService := sentiment.NewSentimentService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := sentimentService.Analyze(sentiment.AnalyzeRequest{
		Text: "this is amazing",
	})
	fmt.Println(rsp)
}
