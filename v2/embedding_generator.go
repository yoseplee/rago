package v2

type EmbeddingGenerator interface {
	Generate(documents Documents) (Embeddings, error)
}
