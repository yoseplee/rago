package main

import (
	"fmt"

	"github.com/yoseplee/rago/config"
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
	for _, documents := range retrieve {
		for _, document := range documents {
			fmt.Printf("%s\n", document)
		}
	}
}
