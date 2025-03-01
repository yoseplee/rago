package openAIEmbedding

import (
	models2 "github.com/yoseplee/rago/v1/models"
)

type Embedding struct {
	v models2.Vector
}

func NewEmbedding(v models2.Vector) Embedding {
	return Embedding{v: v}
}

func (e Embedding) Dimension() models2.Dimension {
	return models2.Dimension(len(e.v))
}

func (e Embedding) Vector() models2.Vector {
	return e.v
}
