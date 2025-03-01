package v2

import (
	"errors"
	"fmt"
)

var (
	IngestErr = errors.New("error occurred while ingesting documents")
)

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
		return fmt.Errorf("%w: %v", IngestErr, loadErr)
	}

	modifiedDocuments, modifyErr := d.DocumentModifiers.ApplyAll(loadedDocuments)
	if modifyErr != nil {
		return fmt.Errorf("%w: %v", IngestErr, modifyErr)
	}

	embeddings, embeddingGenerateErr := d.EmbeddingGenerator.Generate(modifiedDocuments)
	if embeddingGenerateErr != nil {
		return fmt.Errorf("%w: %v", IngestErr, embeddingGenerateErr)
	}

	if storeErr := d.KnowledgeAddable.Add(d.CollectionName, embeddings); storeErr != nil {
		return fmt.Errorf("%w: %v", IngestErr, storeErr)
	}

	return nil
}
