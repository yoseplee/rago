package api

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
	indexName := c.Param("indexName")
	req := RetrieveRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, CommonResponse{Message: err.Error()})
	}

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

	items := []v1.Document{v1.Document(req.Query)}

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
	for _, result := range retrieved {
		documents := result.Documents()
		scores := result.Scores()

		chatCompletion, err := LinecorpOpenAIClient.Client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("CONTEXT"),
				openai.UserMessage("Here are contexts that are relevant to your query."),
				openai.UserMessage("Those context are results of vector search based on the query."),
				openai.UserMessage("Note that scores are numbers from 0 to 1 and represents similarity to the given query."),
				openai.UserMessage(fmt.Sprintf("%v", documents)),
				openai.UserMessage(fmt.Sprintf("%v", scores)),
				openai.UserMessage("==========================="),
				openai.UserMessage("INSTRUCTION"),
				openai.UserMessage("Generate response based on the given context above."),
				openai.UserMessage("You must generate a response that is relevant to the given context."),
				openai.UserMessage("DO NOT MODIFY THE CONTEXT."),
				openai.UserMessage("YOU MUST INCLUDE DETAILED REASONING FOR YOUR RESPONSE."),
				openai.UserMessage("DO YOUR BEST AT REASONING. THE GIVEN SCORES IN CONTEXT SHOULD BE USED ONLY AS REFERENCES FOR YOUR REASONING."),
				openai.UserMessage("==========================="),
				openai.UserMessage("QUERY"),
				openai.UserMessage(req.Message),
				openai.UserMessage("==========================="),
			}),
			Model: openai.F(openai.ChatModelGPT4o),
		})
		if err != nil {
			panic(err.Error())
		}

		chatCompletions = append(chatCompletions, chatCompletion.Choices[0].Message.Content)
	}

	return c.JSON(http.StatusOK, CommonResponse{Message: fmt.Sprintf("%v", chatCompletions)})
}
