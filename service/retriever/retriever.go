package retriever

import (
	"github.com/yoseplee/rago/models"
)

type Retriever interface {
	Retrieve(embedding models.Embedding) error
}
