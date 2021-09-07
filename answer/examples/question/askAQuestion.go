package example

import (
	"fmt"
	"github.com/micro/micro-go/answer"
)

// Ask a question and get an instant answer
func AskAquestion() {
	answerService := answer.NewAnswerService("YOUR_MICRO_TOKEN_HERE")
	rsp, _ := answerService.Question(answer.QuestionRequest{
		Query: "google",
	})
	fmt.Println(rsp)
}
