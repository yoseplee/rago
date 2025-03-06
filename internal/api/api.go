package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yoseplee/rago/infra/logger"
	"github.com/yoseplee/rago/infra/opensearch"
)

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, CommonResponse{Message: "success"})
}

func GetHelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, CommonResponse{Message: "hello world"})
}

func CreateKnnIndex(c echo.Context) error {
	indexName := c.Param("indexName")
	if err := opensearch.GetClient().CreateKnnIndex(indexName); err != nil {
		return c.JSON(http.StatusInternalServerError, CommonResponse{Message: err.Error()})
	}

	logger.Info(
		"create knn index",
		[]logger.F[any]{
			{
				"indexName",
				indexName,
			},
		},
	)
	return c.JSON(http.StatusOK, CommonResponse{Message: "success"})
}
