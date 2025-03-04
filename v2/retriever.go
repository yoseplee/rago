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

	for _, searchResult := range searchResults {
		for _, r := range searchResult {
			logger.Debug(
				"search result",
				[]logger.LogField[any]{
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

	var results []Documents
	for _, searchResult := range searchResults {
		var similarDocuments Documents
		for _, r := range searchResult {
			similarDocuments = append(similarDocuments, r.Document)
		}
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
