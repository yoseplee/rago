package v2

// Ingester is responsible for loading, modifying, embedding-generating, and storing to Vector Search Engine.
type Ingester interface {
	Ingest() error
}

type DefaultIngester struct {
	CollectionName string
	DocumentLoader
	DocumentModifiers
	EmbeddingGenerator
	KnowledgeAddable
}

func (d DefaultIngester) Ingest() error {
	loadedDocuments, loadErr := d.DocumentLoader.Load()
	if loadErr != nil {
		return loadErr
	}

	modifiedDocuments, err := d.DocumentModifiers.ApplyAll(loadedDocuments)
	if err != nil {
		return err
	}

	embeddings, embeddingGenerateErr := d.EmbeddingGenerator.Generate(modifiedDocuments)
	if embeddingGenerateErr != nil {
		return embeddingGenerateErr
	}

	if storeErr := d.KnowledgeAddable.Add(d.CollectionName, embeddings); storeErr != nil {
		return storeErr
	}

	return nil
}
