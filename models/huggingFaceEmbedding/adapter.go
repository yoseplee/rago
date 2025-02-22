package huggingFaceEmbedding

type Adapter interface {
	GenerateEmbedding(document string) (error, Embedding)
}
