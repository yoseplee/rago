package main

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/yoseplee/rago"
	"github.com/yoseplee/rago/infra"
)

func main() {
	defer infra.Logger.Sync()

	ingester := rago.NewDefaultIngester(rago.JSONLoader{FilePath: "sample_shop_items_all.json"})
	if err := ingester.Ingest(); err != nil {
		panic(err)
	}
}

func chatCompletionExample() {
	if chatCompletion, err := infra.OpenAIClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
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
