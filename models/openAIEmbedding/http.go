package openAIEmbedding

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/yoseplee/rago/infra"
	"github.com/yoseplee/rago/models"
)

type OpenAIHTTPAdapter struct {
}

func NewAdapter() *OpenAIHTTPAdapter {
	return &OpenAIHTTPAdapter{}
}

func (o *OpenAIHTTPAdapter) GenerateEmbeddings(documents []string) ([]models.Embedding, error) {
	embeddings, err := infra.OpenAIClient.Embeddings.New(
		context.TODO(),
		openai.EmbeddingNewParams{
			Input: openai.F[openai.EmbeddingNewParamsInputUnion](openai.EmbeddingNewParamsInputArrayOfStrings(documents)),
			Model: openai.F(openai.EmbeddingModelTextEmbedding3Large),
		},
	)
	if err != nil {
		panic(err)
	}

	var result []models.Embedding
	for _, embedding := range embeddings.Data {
		result = append(result, NewEmbedding(embedding.Embedding))
	}

	return result, nil
}
