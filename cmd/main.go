package main

import "github.com/yoseplee/rago"

func main() {
	ingester := rago.NewDefaultIngester()
	if err := ingester.Ingest(); err != nil {
		panic(err)
	}

	//if chatCompletion, err := OpenAIClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
	//	Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
	//		openai.UserMessage("Describe who you are."),
	//	}),
	//	Model: openai.F(openai.ChatModelGPT3_5Turbo),
	//}); err != nil {
	//	panic(err.Error())
	//} else {
	//	fmt.Println(chatCompletion.Choices[0].Message.Content)
	//}
}
