package v2

type Embeddings interface {
	Model() ModelName
	Dimension() Dimension
	Embeddings() []Embedding
}

type ModelName string

type Embedding []float64

type Dimension int
