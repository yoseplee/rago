package openAIEmbedding

type Adapter interface {
	GenerateEmbedding(document string) (Embedding, error)
}
