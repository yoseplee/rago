package openAIEmbedding

import "github.com/yoseplee/rago/models"

type Embedding struct {
	v models.Vector
}

func NewEmbedding(v models.Vector) Embedding {
	return Embedding{v: v}
}

func (e Embedding) Dimension() models.Dimension {
	return models.Dimension(len(e.v))
}

func (e Embedding) Vector() models.Vector {
	return e.v
}
