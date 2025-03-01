package rago

import (
	"bytes"
	"encoding/json"

	"github.com/yoseplee/rago/infra"
	"github.com/yoseplee/rago/models"
	"github.com/yoseplee/rago/models/openAIEmbedding"
	"go.uber.org/zap"
)

// Ingester is an interface that defines the Ingest method.
// Ingest method is responsible for loading, modifying, generating embedding, and store the data.
type Ingester interface {
	Ingest() error
}

type DefaultIngester struct {
	OpenAIEmbeddingAdapter openAIEmbedding.Adapter
	Loader
}

func NewDefaultIngester(loader Loader) *DefaultIngester {
	return &DefaultIngester{
		OpenAIEmbeddingAdapter: openAIEmbedding.NewAdapter(),
		Loader:                 loader,
	}
}

type openSearchDocument struct {
	Embedding models.Vector    `json:"embedding"`
	Dimension models.Dimension `json:"dimension"`
	Content   string           `json:"content"`
}

func (d *DefaultIngester) Ingest() error {
	documents, loadErr := d.Loader.Load()
	if loadErr != nil {
		return loadErr
	}

	embeddings, embeddingErr := d.OpenAIEmbeddingAdapter.GenerateEmbeddings(documents)
	if embeddingErr != nil {
		return embeddingErr
	}

	for i, embedding := range embeddings {
		infra.Logger.Debug(
			"Ingesting embedding to index",
			zap.Int("index", i),
			zap.Int("dimension", int(embedding.Dimension())),
		)

		document := openSearchDocument{
			Embedding: embedding.Vector(),
			Dimension: embedding.Dimension(),
			Content:   documents[i],
		}

		documentJSON, err := json.Marshal(document)
		if err != nil {
			return err
		}

		_, err = infra.OpenSearchClient.Index("sim-search-test", bytes.NewReader(documentJSON))
		if err != nil {
			return err
		}
	}

	return nil
}
