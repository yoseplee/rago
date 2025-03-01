package v2

type Retriever interface {
	Retrieve(document Document) (Documents, error)
}

type DefaultRetriever struct {
	topK int
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

	searchResults, searchErr := d.KnowledgeSearchable.Search(inputEmbeddings, d.topK)
	if searchErr != nil {
		return nil, searchErr
	}

	var results []Documents
	for _, similarDocuments := range searchResults {
		results = append(results, similarDocuments)
	}

	return results, nil
}
