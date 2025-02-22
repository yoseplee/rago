package main

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"rago/config"
)

func main() {
	client := openai.NewClient(
		option.WithAPIKey(config.Config.ApiKey),
		option.WithMaxRetries(config.Config.MaxRetries),
	)

	if chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
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
