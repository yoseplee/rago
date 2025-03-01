package openAIEmbedding

import "github.com/yoseplee/rago/models"

type Adapter interface {
	GenerateEmbeddings(document []string) ([]models.Embedding, error)
}
