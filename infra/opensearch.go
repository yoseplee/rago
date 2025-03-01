package infra

import (
	"crypto/tls"
	"net/http"

	"github.com/opensearch-project/opensearch-go"
	"github.com/yoseplee/rago/config"
)

// implement openSearch client

var OpenSearchClient *opensearch.Client

func init() {
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{config.Config.KnowledgeBase.OpenSearch.Address},
		Username:  config.Config.KnowledgeBase.OpenSearch.Username,
		Password:  config.Config.KnowledgeBase.OpenSearch.Password,
	})

	if err != nil {
		panic(err)
	}
	OpenSearchClient = client
}
