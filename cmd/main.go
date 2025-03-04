package main

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/yoseplee/rago/config"
	"github.com/yoseplee/rago/infra/logger"
	openai2 "github.com/yoseplee/rago/infra/openai"
	"github.com/yoseplee/rago/infra/opensearch"
	"github.com/yoseplee/rago/v1"
)

func main() {
	defer logger.SyncLogger()
	retrieve()
}

func ingest() {
	ingester := v1.DefaultIngester{
		DocumentLoader:    v1.JSONDocumentLoader{FilePath: "data/sample_shop_items.json"},
		DocumentModifiers: nil,
		EmbeddingGenerator: v1.OpenAIEmbeddingGenerator{
			ModelName:            v1.ModelName(config.Config.Ingesters["default"].EmbeddingGenerator.Model),
			Dimension:            v1.Dimension(config.Config.Ingesters["default"].EmbeddingGenerator.Dimension),
			EmbeddingGeneratable: openai2.OpenAIClient,
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

func retrieve() {
	retriever := v1.DefaultRetriever{
		TopK: config.Config.Retrievers["default"].KnowledgeBaseSearch.TopK,
		EmbeddingGenerator: v1.OpenAIEmbeddingGenerator{
			ModelName:            v1.ModelName(config.Config.Retrievers["default"].EmbeddingGenerator.Model),
			Dimension:            v1.Dimension(config.Config.Retrievers["default"].EmbeddingGenerator.Dimension),
			EmbeddingGeneratable: openai2.OpenAIClient,
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
		return
	}

	fmt.Printf("Retrieved %d documents\n", len(retrieved))
	for i, result := range retrieved {
		documents := result.Documents()
		scores := result.Scores()

		chatCompletion, err := openai2.LinecorpOpenAIClient.Client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
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
	}
}
