package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	openai2 "github.com/openai/openai-go"
	"github.com/yoseplee/rago/config"
	"github.com/yoseplee/rago/infra/logger"
	"github.com/yoseplee/rago/infra/openai"
	"github.com/yoseplee/rago/infra/opensearch"
	"github.com/yoseplee/rago/v1"
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

func Ingest(c echo.Context) error {
	indexName := c.Param("indexName")
	req := IngestRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, CommonResponse{Message: err.Error()})
	}

	ingester := v1.DefaultIngester{
		DocumentLoader: v1.StringDocumentLoader{
			Strings: []string{
				req.Document,
			},
		},
		DocumentModifiers: nil,
		EmbeddingGenerator: v1.OpenAIEmbeddingGenerator{
			ModelName:            v1.ModelName(config.Config.Ingesters["default"].EmbeddingGenerator.Model),
			Dimension:            v1.Dimension(config.Config.Ingesters["default"].EmbeddingGenerator.Dimension),
			EmbeddingGeneratable: openai.OpenAIClient,
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
	indexName := c.Param("indexName")

	req := RetrieveRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, CommonResponse{Message: err.Error()})
	}

	if false {
		retriever := v1.DefaultRetriever{
			TopK: config.Config.Retrievers["default"].KnowledgeBaseSearch.TopK,
			EmbeddingGenerator: v1.OpenAIEmbeddingGenerator{
				ModelName:            v1.ModelName(config.Config.Retrievers["default"].EmbeddingGenerator.Model),
				Dimension:            v1.Dimension(config.Config.Retrievers["default"].EmbeddingGenerator.Dimension),
				EmbeddingGeneratable: openai.OpenAIClient,
			},
			KnowledgeSearchable: v1.OpenSearchKnowledgeBase{
				CollectionName:  indexName,
				Indexable:       opensearch.GetClient(),
				IndexSearchable: opensearch.GetClient(),
			},
		}

		items := []v1.Document{v1.Document("")}

		retrieved, err := retriever.Retrieve(items)
		if err != nil {
			logger.Error(
				"failed to retrieve documents",
				[]logger.F[any]{
					{
						"err",
						err.Error(),
					},
				},
			)
			return c.JSON(http.StatusInternalServerError, CommonResponse{Message: err.Error()})
		}

		var chatCompletions []string
		for i, result := range retrieved {
			documents := result.Documents()
			scores := result.Scores()

			chatCompletion, err := openai.LinecorpOpenAIClient.Client.Chat.Completions.New(context.TODO(), openai2.ChatCompletionNewParams{
				Messages: openai2.F([]openai2.ChatCompletionMessageParamUnion{
					openai2.UserMessage(fmt.Sprintf(
						"Here are some context documents retrieved from our Vector Database: [%+v]. These documents are potential candidates. Each candidate has a relevance score between 0 and 1: [%+v].",
						documents,
						scores,
					)),
					openai2.UserMessage(fmt.Sprintf("Please suggest the best alternative for the item [%s].", items[i])),
					openai2.UserMessage("Follow these instructions carefully: Provide the name of the recommended item, explain the reason for each suggestion (do not mention the score), and list up to 5 items as alternatives."),
					openai2.UserMessage("You must exclude the given item from the alternatives."),
					openai2.UserMessage("Note that the user may not be satisfied with your suggestions. Please be cautious with your recommendations."),
				}),
				Model: openai2.F(openai2.ChatModelGPT4o),
			})
			if err != nil {
				panic(err.Error())
			}

			chatCompletions = append(chatCompletions, chatCompletion.Choices[0].Message.Content)
		}

		return c.JSON(http.StatusOK, CommonResponse{Message: fmt.Sprintf("%v", chatCompletions)})
	}

	return c.JSON(http.StatusOK, CommonResponse{Message: "success"})
}
