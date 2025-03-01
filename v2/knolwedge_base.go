package v2

type KnowledgeSearchable interface {
	Search(embeddings Embeddings, topK int) ([]Documents, error)
}

type KnowledgeAddable interface {
	Add(collectionName string, embeddings Embeddings) error
}

type KnowledgeBase interface {
	KnowledgeAddable
	KnowledgeSearchable
}
