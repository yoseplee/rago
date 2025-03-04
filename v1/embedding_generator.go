package v1

import (
	"errors"

	"github.com/yoseplee/rago/infra/logger"
	"github.com/yoseplee/rago/infra/openai"
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
	openai.EmbeddingGeneratable
}

func (o OpenAIEmbeddingGenerator) Generate(documents Documents) (Embeddings, error) {
	embeddings, err := o.Embedding(string(o.ModelName), documents.AsStrings())
	if err != nil {
		panic(err)
	}

	var result []Embedding
	var dimension Dimension
	for _, embedding := range embeddings.Data {
		e := Embedding(embedding.Embedding)
		if dimension == 0 {
			dimension = e.Dimension()
		}

		if dimension != 0 && e.Dimension() != dimension {
			logger.Warn(
				"possible incorrect embedding found",
				[]logger.F[any]{
					{
						"reason",
						"dimension mismatch",
					},
					{
						"index",
						embedding.Index,
					},
				},
			)
		}
		if e.AllZero() {
			logger.Warn(
				"possible incorrect embedding found",
				[]logger.F[any]{
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
		Dimension:  dimension,
		Embeddings: result,
	}, nil
}
