package openAIEmbedding

import "github.com/yoseplee/rago/models"

type Embedding struct {
	D models.Dimension
	V models.Vector
}

func (e Embedding) Dimension() models.Dimension {
	return e.D
}

func (e Embedding) Vector() models.Vector {
	return e.V
}
