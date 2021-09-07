package main

import (
	"fmt"
	"github.com/micro/micro-go/answer"
)

func main() {
	answerService := answer.NewAnswerService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := answerService.Question(answer.QuestionRequest{
		Query: "google",
	})
	fmt.Println(rsp)
}
