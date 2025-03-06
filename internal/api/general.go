package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yoseplee/rago/config"
	. "github.com/yoseplee/rago/infra/openai"
	"github.com/yoseplee/rago/infra/opensearch"
	v1 "github.com/yoseplee/rago/v1"
)

func Ingest(c echo.Context) error {
	indexName := c.Param("indexName")
	req := IngestRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, CommonResponse{
			Message: err.Error(),
		})
	}

	if len(req.Documents) == 0 {
		return c.JSON(http.StatusBadRequest, CommonResponse{
			Message: "documents is empty",
		})
	}

	ingester := v1.DefaultIngester{
		DocumentLoader: v1.StringDocumentLoader{
			Strings: req.Documents,
		},
		DocumentModifiers: nil,
		EmbeddingGenerator: v1.OpenAIEmbeddingGenerator{
			ModelName:            v1.ModelName(config.Config.Ingesters["default"].EmbeddingGenerator.Model),
			Dimension:            v1.Dimension(config.Config.Ingesters["default"].EmbeddingGenerator.Dimension),
			EmbeddingGeneratable: OpenAIClient,
		},
		KnowledgeAddable: v1.OpenSearchKnowledgeBase{
			CollectionName:  indexName,
			Indexable:       opensearch.GetClient(),
			IndexSearchable: opensearch.GetClient(),
		},
	}

	if err := ingester.Ingest(); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, CommonResponse{Message: "success"})
}

func Retrieve(c echo.Context) error {
	return nil
}
