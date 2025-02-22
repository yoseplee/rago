package ingest

import (
	"github.com/yoseplee/rago/models"
)

// OpenSearchIngester is a struct that implements Ingester interface
// This struct is used to ingest raw data into OpenSearch (OpenSearch generates embeddings in a nutshell).
type OpenSearchIngester struct {
}

func (o *OpenSearchIngester) Ingest(embedding models.Embedding) error {
	//TODO implement me
	panic("implement me")
}
