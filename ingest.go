package rago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yoseplee/rago/infra"
	"github.com/yoseplee/rago/models/openAIEmbedding"
	"os"
)

type Ingester interface {
	Ingest() error
}

type DefaultIngester struct {
	OpenAIEmbeddingAdapter    openAIEmbedding.Adapter
	VectorSearchEngineAdapter interface{}
}

func NewDefaultIngester() *DefaultIngester {
	return &DefaultIngester{
		OpenAIEmbeddingAdapter:    openAIEmbedding.NewAdapter(),
		VectorSearchEngineAdapter: nil,
	}
}

type ShopItem struct {
	ID                 string `json:"_id"`
	ChainItemID        string `json:"chainItemId"`
	ChainID            string `json:"chainId"`
	ShopID             string `json:"shopId"`
	ItemName           string `json:"itemName"`
	ItemClassification string `json:"itemClassification"`
	DisplayRule        struct {
		ItemImage struct {
			ObsOid   string `json:"obsOid"`
			ImageURL string `json:"imageUrl"`
		} `json:"itemImage"`
		Description       string `json:"description"`
		DescriptionDetail string `json:"descriptionDetail"`
		SortOrder         *int   `json:"sortOrder"`
	} `json:"displayRule"`
}

func (d *DefaultIngester) Ingest() error {
	data, fileReadErr := os.ReadFile("sample_shop_items_all.json")
	if fileReadErr != nil {
		return fileReadErr
	}

	var shopItems []ShopItem
	if unmarshalErr := json.Unmarshal(data, &shopItems); unmarshalErr != nil {
		return unmarshalErr
	}

	var documents []string
	for _, shopItem := range shopItems {
		documents = append(documents, fmt.Sprintf("%v", shopItem))
	}

	embeddings, fileReadErr := d.OpenAIEmbeddingAdapter.GenerateEmbeddings(documents)
	if fileReadErr != nil {
		return fileReadErr
	}

	for i, embedding := range embeddings {
		fmt.Printf("[%d] embedding: %+v\n", i, embedding.Vector())
	}

	for i, embedding := range embeddings {
		fmt.Printf("[%d] dimension of embedding: %+v\n", i, embedding.Dimension())
	}

	for i, embedding := range embeddings {
		document := map[string]interface{}{
			"id":        shopItems[i].ID,
			"embedding": embedding.Vector(),
			"dimension": embedding.Dimension(),
			"shopItem":  shopItems[i],
		}

		documentJSON, err := json.Marshal(document)
		if err != nil {
			return err
		}

		_, err = infra.OpenSearchClient.Index("sim-search-test", bytes.NewReader(documentJSON))
		if err != nil {
			return err
		}
	}

	return nil
}
