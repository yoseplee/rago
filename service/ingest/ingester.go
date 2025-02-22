package ingest

import (
	"github.com/yoseplee/rago/models"
)

type Ingester interface {
	Ingest(embedding models.Embedding) error
}
