package main

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/yoseplee/rago/config"
	"github.com/yoseplee/rago/infra"
	"github.com/yoseplee/rago/infra/logger"
	"github.com/yoseplee/rago/infra/opensearch"
	v2 "github.com/yoseplee/rago/v2"
)

func main() {
	config.InitConfig()
	defer logger.SyncLogger()

	retriever := v2.DefaultRetriever{
		TopK: config.Config.Retrievers["default"].KnowledgeBaseSearch.TopK,
		EmbeddingGenerator: v2.OpenAIEmbeddingGenerator{
			ModelName: v2.ModelName(config.Config.Retrievers["default"].EmbeddingGenerator.Model),
			Dimension: v2.Dimension(config.Config.Retrievers["default"].EmbeddingGenerator.Dimension),
			Client:    infra.OpenAIClient,
		},
		KnowledgeSearchable: v2.OpenSearchKnowledgeBase{
			CollectionName:  config.Config.Retrievers["default"].KnowledgeBaseSearch.Collection,
			Indexable:       opensearch.GetClient(),
			IndexSearchable: opensearch.GetClient(),
		},
	}

	item := []v2.Document{
		"大塚製薬　ポカリスエット　500ml（45019517）",
		"アンシャンテ メイクアップスポンジ 三角タイプ 38個（4540474777979）",
	}

	retrieved, err := retriever.Retrieve(item)
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
	for _, result := range retrieved {
		documents := result.Documents()
		scores := result.Scores()

		chatCompletion, err := infra.LinecorpOpenAIClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(fmt.Sprintf(
					"Here are some context documents retrieved from our Vector Database: [%+v]. These documents are potential candidates. Each candidate has a relevance score between 0 and 1: [%+v].",
					documents,
					scores,
				)),
				openai.UserMessage(fmt.Sprintf("Please suggest the best alternative for the item [%s].", item)),
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
