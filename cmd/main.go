package main

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/yoseplee/rago/config"
	"github.com/yoseplee/rago/infra"
	"github.com/yoseplee/rago/infra/logger"
	v2 "github.com/yoseplee/rago/v2"
)

func main() {
	defer logger.SyncLogger()

	retriever := v2.DefaultRetriever{
		CollectionName: config.Config.Retrievers["default"].KnowledgeBaseSearch.Collection,
		TopK:           config.Config.Retrievers["default"].KnowledgeBaseSearch.TopK,
		EmbeddingGenerator: v2.OpenAIEmbeddingGenerator{
			ModelName: v2.ModelName(config.Config.Retrievers["default"].Embedding.Model),
			Dimension: v2.Dimension(config.Config.Retrievers["default"].Embedding.Dimension),
		},
		KnowledgeSearchable: v2.OpenSearchKnowledgeBase{},
	}

	retrieve, err := retriever.Retrieve("大塚製薬　ポカリスエット　500ml（45019517）")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Retrieved %d documents\n", len(retrieve))
	for _, results := range retrieve {
		var documents v2.Documents
		var scores []float64
		for _, r := range results {
			documents = append(documents, r.Document)
			scores = append(scores, r.Score)
		}

		chatCompletion, err := infra.LinecorpOpenAIClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(fmt.Sprintf(
					"Here are some context documents retrieved from our Vector Database: [%+v]. These documents are potential candidates. Each candidate has a relevance score between 0 and 1: [%+v].",
					documents,
					scores,
				)),
				openai.UserMessage("Please suggest the best alternative for the item [大塚製薬　ポカリスエット　500ml（45019517）]."),
				openai.UserMessage("Follow these instructions carefully: Provide the name of the recommended item, explain the reason for each suggestion (do not mention the score), and list up to 5 items as alternatives."),
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
