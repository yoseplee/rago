package main

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
)

func main() {
	client := NewClient()

	if chatCompletion, err := client.OpenAIClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Describe who you are."),
		}),
		Model: openai.F(openai.ChatModelGPT3_5Turbo),
	}); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(chatCompletion.Choices[0].Message.Content)
	}
}
