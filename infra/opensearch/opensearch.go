package opensearch

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"github.com/yoseplee/rago/config"
	"github.com/yoseplee/rago/infra/logger"
)

var c *opensearch.Client

func init() {
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{config.Config.KnowledgeBases.Opensearch["verda-dev"].Address},
		Username:  config.Config.KnowledgeBases.Opensearch["verda-dev"].Username,
		Password:  config.Config.KnowledgeBases.Opensearch["verda-dev"].Password,
	})

	if err != nil {
		panic(err)
	}
	c = client
}

// Index stores document into opensearch index.
func Index(indexName string, document Document) error {
	logger.Debug(
		"index document",
		[]logger.LogField[any]{
			{
				"indexName",
				indexName,
			},
			{
				"document",
				document,
			},
		},
	)

	documentJson, jsonErr := document.Json()
	if jsonErr != nil {
		return jsonErr
	}

	_, indexErr := c.Index(indexName, bytes.NewReader(documentJson))
	if indexErr != nil {
		return indexErr
	}

	return nil
}

func Search(indexNames []string, query Query) (Response, error) {
	logger.Debug(
		"search index",
		[]logger.LogField[any]{
			{
				"index",
				indexNames,
			},
			{
				"query",
				query.String(),
			},
		},
	)

	searchResponse, searchErr := c.Search(
		func(req *opensearchapi.SearchRequest) {
			req.Index = indexNames
			req.Body = strings.NewReader(query.String())
		})
	if searchErr != nil {
		return Response{}, searchErr
	}
	defer searchResponse.Body.Close()

	var response Response
	if decodeErr := json.NewDecoder(searchResponse.Body).Decode(&response); decodeErr != nil {
		return Response{}, decodeErr
	}

	return response, nil
}

func CreateKnnIndex(indexName string) error {
	indexConfig := KNNIndexConfig{
		Settings: Settings{
			Index: IndexSettings{
				Knn:                  true,
				KnnAlgoParamEfSearch: 100,
			},
		},
		Mappings: Mappings{
			Properties: Properties{
				Embedding: Embedding{
					Type:      "knn_vector",
					Dimension: 3072,
					Method: Method{
						Name:      "hnsw",
						SpaceType: "l2",
						Engine:    "nmslib",
						Parameters: struct {
							EfConstruction int `json:"ef_construction"`
							M              int `json:"m"`
						}{
							EfConstruction: 100,
							M:              16,
						},
					},
				},
			},
		},
	}

	body, err := json.Marshal(indexConfig)
	if err != nil {
		return err
	}

	createIndexResponse, err := c.Indices.Create(
		indexName,
		func(req *opensearchapi.IndicesCreateRequest) {
			req.Body = bytes.NewReader(body)
		})
	if err != nil {
		return err
	}
	defer createIndexResponse.Body.Close()

	if createIndexResponse.IsError() {
		return fmt.Errorf("error creating index: %s", createIndexResponse.String())
	}

	return nil
}

func DeleteIndex(indexName string) error {
	panic("Not Implemented yet.")
}
