package main

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	. "github.com/yoseplee/rago/infra"
	"github.com/yoseplee/rago/models/openAIEmbedding"
)

func main() {
	oa := openAIEmbedding.OpenAIHTTPAdapter{}
	embeddings, err := oa.GenerateEmbeddings([]string{"my new world"})
	if err != nil {
		panic(err)
	}
	for i, embedding := range embeddings {
		fmt.Printf("[%d] embedding: %+v\n: ", i, embedding.Vector())
	}
	fmt.Println()

	//embeddings, err := OpenAIClient.Embeddings.New(
	//	context.TODO(),
	//	openai.EmbeddingNewParams{
	//		Input: openai.F[openai.EmbeddingNewParamsInputUnion](openai.EmbeddingNewParamsInputArrayOfStrings{
	//			"my new world",
	//		}),
	//		Model: openai.F(openai.EmbeddingModelTextEmbedding3Small),
	//	},
	//)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("model: ", embeddings.Model)
	//for i, embedding := range embeddings.Data {
	//	fmt.Printf("[%d] embedding: %v\n", i, embedding.v)
	//}
	//fmt.Println()
	//fmt.Printf("%+v\n", embeddings)

	if chatCompletion, err := OpenAIClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
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
