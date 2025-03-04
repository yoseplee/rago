package v2

import "github.com/yoseplee/rago/infra/opensearch"

type KnowledgeSearchable interface {
	Search(collectionName string, embeddings Embeddings, topK int) ([]Documents, error)
}

type KnowledgeAddable interface {
	Add(collectionName string, embeddings Embeddings, contents Documents) error
}

type KnowledgeBase interface {
	KnowledgeAddable
	KnowledgeSearchable
}

type OpenSearchKnowledgeBase struct{}

func (o OpenSearchKnowledgeBase) Add(collectionName string, embeddings Embeddings, contents Documents) error {
	for i, e := range embeddings.Embeddings {
		document := opensearch.Document{
			Embedding: e,
			Dimension: int(embeddings.Dimension),
			Content:   contents[i],
		}

		if err := opensearch.Index(collectionName, document); err != nil {
			return err
		}

	}
	return nil
}

func (o OpenSearchKnowledgeBase) Search(collectionName string, embeddings Embeddings, topK int) ([]Documents, error) {
	var results []Documents
	for _, e := range embeddings.Embeddings {
		queryResult, err := opensearch.Search([]string{collectionName}, opensearch.NewKNNQuery(e, topK))
		if err != nil {
			return nil, err
		}
		var searchResult Documents
		for _, hit := range queryResult.Hits.Hits {
			searchResult = append(searchResult, Document(hit.Source.Content))
		}

		results = append(results, searchResult)
	}
	return results, nil
}
