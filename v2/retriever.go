package v2

import (
	"github.com/yoseplee/rago/infra/logger"
)

type Retriever interface {
	Retrieve(document Document) (Retrieved, error)
}

type Retrieved []SimilarKnowledgeSearchResults

type DefaultRetriever struct {
	TopK int
	EmbeddingGenerator
	KnowledgeSearchable
}

type CandidateGenerator interface { // To Vector Search Engine
	Generate(embeddings Embeddings) ([]Documents, error)
}

func (d DefaultRetriever) Retrieve(documents Documents) (Retrieved, error) {
	inputEmbeddings, embeddingGenerateErr := d.EmbeddingGenerator.Generate(documents)
	if embeddingGenerateErr != nil {
		return nil, embeddingGenerateErr
	}

	logger.Debug(
		"input embedding",
		[]logger.F[any]{
			{
				"documents",
				documents,
			},
			{
				"embeddings",
				inputEmbeddings,
			},
		},
	)

	searchResults, searchErr := d.KnowledgeSearchable.Search(inputEmbeddings, d.TopK)
	if searchErr != nil {
		return nil, searchErr
	}

	for _, searchResult := range searchResults {
		for _, r := range searchResult {
			logger.Debug(
				"search result",
				[]logger.F[any]{
					{
						"score",
						r.Score,
					},
					{
						"document",
						r.Document,
					},
				},
			)
		}
	}

	return searchResults, nil
}
