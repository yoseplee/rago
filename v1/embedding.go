package v1

type Embeddings struct {
	ModelName
	Dimension
	Embeddings []Embedding
}

type ModelName string

type Embedding []float64

func (e Embedding) Dimension() Dimension {
	return Dimension(len(e))
}

func (e Embedding) AllZero() bool {
	for _, v := range e {
		if v != 0 {
			return false
		}
	}
	return true
}

type Dimension int
