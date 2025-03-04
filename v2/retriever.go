package v2

import "github.com/yoseplee/rago/infra/logger"

type Retriever interface {
	Retrieve(document Document) (Documents, error)
}

type DefaultRetriever struct {
	CollectionName string
	TopK           int
	EmbeddingGenerator
	KnowledgeSearchable
}

type CandidateGenerator interface { // To Vector Search Engine
	Generate(embeddings Embeddings) ([]Documents, error)
}

func (d DefaultRetriever) Retrieve(document Document) ([]Documents, error) {
	inputEmbeddings, embeddingGenerateErr := d.EmbeddingGenerator.Generate([]Document{document})
	if embeddingGenerateErr != nil {
		return nil, embeddingGenerateErr
	}

	logger.Debug(
		"input embedding",
		[]logger.LogField[any]{
			{
				"documents",
				document,
			},
			{
				"embeddings",
				inputEmbeddings,
			},
		},
	)

	searchResults, searchErr := d.KnowledgeSearchable.Search(d.CollectionName, inputEmbeddings, d.TopK)
	if searchErr != nil {
		return nil, searchErr
	}

	var results []Documents
	for _, similarDocuments := range searchResults {
		results = append(results, similarDocuments)
	}

	logger.Debug(
		"search result",
		[]logger.LogField[any]{
			{
				"document",
				document,
			},
			{
				"results",
				results,
			},
		},
	)

	return results, nil
}
