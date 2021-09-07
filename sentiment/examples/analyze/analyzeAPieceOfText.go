package example

import (
	"fmt"
	"github.com/micro/micro-go/sentiment"
)

// Analyze and score a piece of text
func AnalyzeApieceOfText() {
	sentimentService := sentiment.NewSentimentService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := sentimentService.Analyze(sentiment.AnalyzeRequest{
		Text: "this is amazing",
	})
	fmt.Println(rsp)
}
