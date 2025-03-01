package openAIEmbedding

import (
	"github.com/yoseplee/rago/v1/models"
)

type Adapter interface {
	GenerateEmbeddings(document []string) ([]models.Embedding, error)
}
