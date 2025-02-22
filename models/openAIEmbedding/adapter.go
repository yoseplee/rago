package openAIEmbedding

type Adapter interface {
	GenerateEmbedding(document string) (error, Embedding)
}
