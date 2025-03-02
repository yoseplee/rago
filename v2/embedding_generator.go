package v2

import (
	"context"
	"errors"

	"github.com/openai/openai-go"
	"github.com/yoseplee/rago/infra"
	"github.com/yoseplee/rago/infra/logger"
)

var (
	EmbeddingEmptyErr          = errors.New("embedding is empty")
	EmbeddingGenerateFailedErr = errors.New("embedding generator failed")
)

type EmbeddingGenerator interface {
	Generate(documents Documents) (Embeddings, error)
}

type OpenAIEmbeddingGenerator struct {
	ModelName
	Dimension
}

func (o OpenAIEmbeddingGenerator) Generate(documents Documents) (Embeddings, error) {
	embeddings, err := infra.OpenAIClient.Embeddings.New(
		context.TODO(),
		openai.EmbeddingNewParams{
			Input: openai.F[openai.EmbeddingNewParamsInputUnion](openai.EmbeddingNewParamsInputArrayOfStrings(documents.AsStrings())),
			Model: openai.F(openai.EmbeddingModelTextEmbedding3Large),
		},
	)
	if err != nil {
		panic(err)
	}

	var result []Embedding
	for _, embedding := range embeddings.Data {
		e := Embedding(embedding.Embedding)
		if e.AllZero() {
			logger.Warn(
				"possible incorrect embedding found",
				[]logger.LogField[any]{
					{
						"reason",
						"vector is consists of zeros",
					},
					{
						"index",
						embedding.Index,
					},
				},
			)
		}
		result = append(result, e)
	}

	return Embeddings{
		ModelName:  ModelName(embeddings.Model),
		Dimension:  0,
		Embeddings: nil,
	}, nil
}
