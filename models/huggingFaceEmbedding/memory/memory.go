package memory

import (
	"github.com/yoseplee/rago/models/huggingFaceEmbedding"
)

type HuggingFaceMemoryAdapter struct{}

func (h *HuggingFaceMemoryAdapter) GenerateEmbedding(document string) (error, huggingFaceEmbedding.Embedding) {
	//TODO implement me
	panic("implement me")
}
