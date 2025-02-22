package retriever

import (
	"github.com/yoseplee/rago/models"
)

// OpenSearchRetriever is a struct that implements Retriever interface.
// This struct is used to retrieve similar documents from OpenSearch.
type OpenSearchRetriever struct {
}

func (o *OpenSearchRetriever) Retrieve(embedding models.Embedding) error {
	//TODO implement me
	panic("implement me")
}
