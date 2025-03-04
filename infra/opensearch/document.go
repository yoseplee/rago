package opensearch

import "encoding/json"

type Document struct {
	Embedding []float64   `json:"embedding"`
	Dimension int         `json:"dimension"`
	Content   interface{} `json:"content"`
}

func (o Document) Json() ([]byte, error) {
	return json.Marshal(o)
}
