package openAIEmbedding

type Adapter interface {
	GenerateEmbeddings(document []string) ([]Embedding, error)
}
