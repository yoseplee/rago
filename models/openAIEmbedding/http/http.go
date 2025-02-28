package http

import (
	"context"
	"github.com/openai/openai-go"
	"github.com/yoseplee/rago/infra"
	"github.com/yoseplee/rago/models/openAIEmbedding"
)

type OpenAIHTTPAdapter struct {
}

func (o *OpenAIHTTPAdapter) GenerateEmbeddings(documents []string) ([]openAIEmbedding.Embedding, error) {
	embeddings, err := infra.OpenAIClient.Embeddings.New(
		context.TODO(),
		openai.EmbeddingNewParams{
			Input: openai.F[openai.EmbeddingNewParamsInputUnion](openai.EmbeddingNewParamsInputArrayOfStrings(documents)),
			Model: openai.F(openai.EmbeddingModelTextEmbedding3Small),
		},
	)
	if err != nil {
		panic(err)
	}

	var result []openAIEmbedding.Embedding
	for _, embedding := range embeddings.Data {
		result = append(result, openAIEmbedding.NewEmbedding(embedding.Embedding))
	}

	return result, nil
}
