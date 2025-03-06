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

	e.GET("/", getHelloWorld)

	e.GET("/retrieve", getSample)
	e.Logger.Fatal(e.Start(":1323"))
}

var getHelloWorld = func(c echo.Context) error {
	return c.String(200, "Hello, World!")
}

var getSample = func(c echo.Context) error {
	completions, err := retrieve()
	var jsons []string
	for _, c := range completions {
		jsons = append(jsons, c.JSON.RawJSON())
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, jsons)
}

func ingest() {
	ingester := v1.DefaultIngester{
		DocumentLoader:    v1.JSONDocumentLoader{FilePath: "data/sample_shop_items.json"},
		DocumentModifiers: nil,
		EmbeddingGenerator: v1.OpenAIEmbeddingGenerator{
			ModelName:            v1.ModelName(config.Config.Ingesters["default"].EmbeddingGenerator.Model),
			Dimension:            v1.Dimension(config.Config.Ingesters["default"].EmbeddingGenerator.Dimension),
			EmbeddingGeneratable: OpenAIClient,
		},
		KnowledgeAddable: v1.OpenSearchKnowledgeBase{
			CollectionName:  config.Config.Ingesters["default"].KnowledgeBaseAdd.Collection,
			Indexable:       opensearch.GetClient(),
			IndexSearchable: opensearch.GetClient(),
		},
	}

	if err := ingester.Ingest(); err != nil {
		panic(err)
	}
}

func retrieve() ([]*openai.ChatCompletion, error) {
	retriever := v1.DefaultRetriever{
		TopK: config.Config.Retrievers["default"].KnowledgeBaseSearch.TopK,
		EmbeddingGenerator: v1.OpenAIEmbeddingGenerator{
			ModelName:            v1.ModelName(config.Config.Retrievers["default"].EmbeddingGenerator.Model),
			Dimension:            v1.Dimension(config.Config.Retrievers["default"].EmbeddingGenerator.Dimension),
			EmbeddingGeneratable: OpenAIClient,
		},
		KnowledgeSearchable: v1.OpenSearchKnowledgeBase{
			CollectionName:  config.Config.Retrievers["default"].KnowledgeBaseSearch.Collection,
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
		return nil, err
	}

	fmt.Printf("Retrieved %d documents\n", len(retrieved))
	var chatCompletions []*openai.ChatCompletion
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

		fmt.Println(chatCompletion.Choices[0].Message.Content)
		chatCompletions = append(chatCompletions, chatCompletion)
	}
	return chatCompletions, nil
}
