package http

import (
	"github.com/yoseplee/rago/models/openAIEmbedding"
)

type OpenAIHTTPAdapter struct {
}

func (o *OpenAIHTTPAdapter) GenerateEmbedding(document string) (error, openAIEmbedding.Embedding) {
	//TODO implement me
	panic("implement me")
}
