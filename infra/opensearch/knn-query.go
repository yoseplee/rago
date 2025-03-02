package opensearch

import (
	"encoding/json"

	v2 "github.com/yoseplee/rago/v2"
)

type Query interface {
	String() string
}

type KNNQuery struct {
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
	embedding v2.Embedding,
	topK int,
) KNNQuery {
	return KNNQuery{
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
