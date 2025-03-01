package v1

import (
	"bytes"
	"encoding/json"

	"github.com/yoseplee/rago/infra"
	models2 "github.com/yoseplee/rago/v1/models"
	openAIEmbedding2 "github.com/yoseplee/rago/v1/models/openAIEmbedding"
	"go.uber.org/zap"
)

// Ingester is an interface that defines the Ingest method.
// Ingest method is responsible for loading, modifying, generating embedding, and store the data.
type Ingester interface {
	Ingest() error
}

type DefaultIngester struct {
	OpenAIEmbeddingAdapter openAIEmbedding2.Adapter
	Loader
}

func NewDefaultIngester(loader Loader) *DefaultIngester {
	return &DefaultIngester{
		OpenAIEmbeddingAdapter: openAIEmbedding2.NewAdapter(),
		Loader:                 loader,
	}
}

type openSearchDocument struct {
	Embedding models2.Vector    `json:"embedding"`
	Dimension models2.Dimension `json:"dimension"`
	Content   string            `json:"content"`
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
