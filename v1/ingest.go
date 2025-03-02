package v1

import (
	"github.com/yoseplee/rago/infra/opensearch"
	models2 "github.com/yoseplee/rago/v1/models"
	openAIEmbedding2 "github.com/yoseplee/rago/v1/models/openAIEmbedding"
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
		o := opensearch.Document{
			Embedding: embedding.Vector(),
			Dimension: int(embedding.Dimension()),
			Content:   documents[i],
		}

		if err := opensearch.IndexDocument("sim-search-test", o); err != nil {
			return err
		}
	}

	return nil
}
