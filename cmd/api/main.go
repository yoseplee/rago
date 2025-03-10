package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yoseplee/rago/infra/logger"
	"github.com/yoseplee/rago/internal/api"
)

func main() {
	defer logger.SyncLogger()

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", api.GetHelloWorld)
	e.GET("/healthCheck", api.HealthCheck)
	e.GET("/indices", api.ListIndices)
	e.POST("/index/knn/:indexName", api.CreateKnnIndex)
	e.GET("/retrieve/:indexName", api.Retrieve)
	e.POST("/ingest/:indexName", api.Ingest)

	e.GET("/retrieve/similar/item/:indexName", api.RetrieveSimilarItems)
	e.POST("/ingest/similar/item/:indexName", api.IngestSimilarItem)

	// start the echo server.
	e.Logger.Fatal(e.Start(":1323"))
}
