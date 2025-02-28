package openAIEmbedding

import (
	"context"
	"github.com/openai/openai-go"
	"github.com/yoseplee/rago/infra"
)

type OpenAIHTTPAdapter struct {
}

func (o *OpenAIHTTPAdapter) GenerateEmbeddings(documents []string) ([]Embedding, error) {
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

	var result []Embedding
	for _, embedding := range embeddings.Data {
		result = append(result, NewEmbedding(embedding.Embedding))
	}

	return result, nil
}
