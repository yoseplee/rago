package opensearch

import (
	"encoding/json"
)

type Query interface {
	String() string
}

type KNNQuery struct {
	Size  int `json:"size"`
	Query struct {
		Knn struct {
			Embedding struct {
				Vector []float64 `json:"vector"`
				K      int       `json:"k"`
			} `json:"embedding"`
		} `json:"knn"`
	} `json:"query"`
}

func (k KNNQuery) String() string {
	marshal, _ := json.Marshal(k)
	return string(marshal)
}

func NewKNNQuery(
	embedding []float64,
	topK int,
) KNNQuery {
	return KNNQuery{
		Size: topK,
		Query: struct {
			Knn struct {
				Embedding struct {
					Vector []float64 `json:"vector"`
					K      int       `json:"k"`
				} `json:"embedding"`
			} `json:"knn"`
		}{
			Knn: struct {
				Embedding struct {
					Vector []float64 `json:"vector"`
					K      int       `json:"k"`
				} `json:"embedding"`
			}{
				Embedding: struct {
					Vector []float64 `json:"vector"`
					K      int       `json:"k"`
				}{
					Vector: embedding,
					K:      topK,
				},
			},
		},
	}
}
