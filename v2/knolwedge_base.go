package v2

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
	//TODO implement me
	panic("implement me")
}

func (o OpenSearchKnowledgeBase) Search(collectionName string, embeddings Embeddings, topK int) ([]Documents, error) {
	//TODO implement me
	panic("implement me")
}
