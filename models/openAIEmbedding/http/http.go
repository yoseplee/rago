package http

import (
	"github.com/yoseplee/rago/models/openAIEmbedding"
)

type OpenAIHTTPAdapter struct {
}

func (o *OpenAIHTTPAdapter) GenerateEmbedding(document string) (openAIEmbedding.Embedding, error) {
	//embeddings, err := infra.OpenAIClient.Embeddings.New(
	//	context.TODO(),
	//	openai.EmbeddingNewParams{
	//		Input: openai.F[openai.EmbeddingNewParamsInputUnion](openai.EmbeddingNewParamsInputArrayOfStrings{
	//			document,
	//		}),
	//		Model: openai.F(openai.EmbeddingModelTextEmbedding3Small),
	//	},
	//)
	//if err != nil {
	//	panic(err)
	//}

	return openAIEmbedding.Embedding{}, nil
}
