package v2

import "github.com/yoseplee/rago/infra/opensearch"

type KnowledgeSearchable interface {
	Search(collectionName string, embeddings Embeddings, topK int) ([]Documents, error)
}

type KnowledgeAddable interface {
	Add(collectionName string, embeddings Embeddings) error
}

type KnowledgeBase interface {
	KnowledgeAddable
	KnowledgeSearchable
}

type OpenSearchKnowledgeBase struct{}

func (o OpenSearchKnowledgeBase) Add(collectionName string, embeddings Embeddings) error {
	for _, e := range embeddings.Embeddings {
		document := opensearch.Document{
			Embedding: e,
			Dimension: int(embeddings.Dimension),
			Content:   "",
		}

		if err := opensearch.Index(collectionName, document); err != nil {
			return err
		}

	}
	return nil
}

func (o OpenSearchKnowledgeBase) Search(collectionName string, embeddings Embeddings, topK int) ([]Documents, error) {
	//TODO implement me
	panic("implement me")
}
