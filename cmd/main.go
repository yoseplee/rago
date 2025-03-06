package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/openai/openai-go"
	"github.com/yoseplee/rago/config"
	"github.com/yoseplee/rago/infra/logger"
	. "github.com/yoseplee/rago/infra/openai"
	"github.com/yoseplee/rago/infra/opensearch"
	"github.com/yoseplee/rago/v1"
)

func main() {
	defer logger.SyncLogger()

	e := echo.New()

	// index name will be an identifier for this app.
	e.GET("/", getHelloWorld)

	e.GET("/healthCheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/retrieve/:indexName", retrieve)

	e.POST("/ingest/:indexName", ingest)

	// start the echo server.
	e.Logger.Fatal(e.Start(":1323"))
}

var getHelloWorld = func(c echo.Context) error {
	return c.String(200, "Hello, World!")
}

func ingest(c echo.Context) error {
	indexName := c.Param("indexName")

	ingester := v1.DefaultIngester{
		DocumentLoader:    v1.JSONDocumentLoader{FilePath: "data/sample_shop_items.json"},
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

	return c.JSON(http.StatusOK, "OK")
}

func retrieve(c echo.Context) error {
	indexName := c.Param("indexName")

	retriever := v1.DefaultRetriever{
		TopK: config.Config.Retrievers["default"].KnowledgeBaseSearch.TopK,
		EmbeddingGenerator: v1.OpenAIEmbeddingGenerator{
			ModelName:            v1.ModelName(config.Config.Retrievers["default"].EmbeddingGenerator.Model),
			Dimension:            v1.Dimension(config.Config.Retrievers["default"].EmbeddingGenerator.Dimension),
			EmbeddingGeneratable: OpenAIClient,
		},
		KnowledgeSearchable: v1.OpenSearchKnowledgeBase{
			CollectionName:  indexName,
			Indexable:       opensearch.GetClient(),
			IndexSearchable: opensearch.GetClient(),
		},
	}

	items := []v1.Document{
		"大塚製薬　ポカリスエット　500ml（45019517）",
		"アンシャンテ メイクアップスポンジ 三角タイプ 38個（4540474777979）",
	}

	retrieved, err := retriever.Retrieve(items)
	if err != nil {
		logger.Error(
			"failed to retrive documents",
			[]logger.F[any]{
				{
					"err",
					err.Error(),
				},
			},
		)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Printf("Retrieved %d documents\n", len(retrieved))
	var chatCompletions []string
	for i, result := range retrieved {
		documents := result.Documents()
		scores := result.Scores()

		chatCompletion, err := LinecorpOpenAIClient.Client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(fmt.Sprintf(
					"Here are some context documents retrieved from our Vector Database: [%+v]. These documents are potential candidates. Each candidate has a relevance score between 0 and 1: [%+v].",
					documents,
					scores,
				)),
				openai.UserMessage(fmt.Sprintf("Please suggest the best alternative for the item [%s].", items[i])),
				openai.UserMessage("Follow these instructions carefully: Provide the name of the recommended item, explain the reason for each suggestion (do not mention the score), and list up to 5 items as alternatives."),
				openai.UserMessage("You must exclude the given item from the alternatives."),
				openai.UserMessage("Note that the user may not be satisfied with your suggestions. Please be cautious with your recommendations."),
			}),
			Model: openai.F(openai.ChatModelGPT4o),
		})
		if err != nil {
			panic(err.Error())
		}

		chatCompletions = append(chatCompletions, chatCompletion.Choices[0].Message.Content)
	}

	return c.JSON(http.StatusOK, chatCompletions)
}
