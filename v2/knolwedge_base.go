package v2

import "github.com/yoseplee/rago/infra/opensearch"

type KnowledgeSearchable interface {
	Search(embeddings Embeddings, topK int) ([]SimilarKnowledgeSearchResults, error)
}

type SimilarKnowledgeSearchResult struct {
	Document
	Score float64
}

type SimilarKnowledgeSearchResults []SimilarKnowledgeSearchResult

type KnowledgeAddable interface {
	Add(embeddings Embeddings, contents Documents) error
}

type KnowledgeBase interface {
	KnowledgeAddable
	KnowledgeSearchable
}

type OpenSearchKnowledgeBase struct {
	CollectionName string
}

func (o OpenSearchKnowledgeBase) Add(embeddings Embeddings, contents Documents) error {
	for i, e := range embeddings.Embeddings {
		document := opensearch.Document{
			Embedding: e,
			Dimension: int(embeddings.Dimension),
			Content:   contents[i],
		}

		if err := opensearch.Index(o.CollectionName, document); err != nil {
			return err
		}

	}
	return nil
}

func (o OpenSearchKnowledgeBase) Search(embeddings Embeddings, topK int) ([]SimilarKnowledgeSearchResults, error) {
	var results []SimilarKnowledgeSearchResults
	for _, e := range embeddings.Embeddings {
		queryResult, err := opensearch.Search([]string{o.CollectionName}, opensearch.NewKNNQuery(e, topK))
		if err != nil {
			return nil, err
		}
		var searchResult SimilarKnowledgeSearchResults
		for _, hit := range queryResult.Hits.Hits {
			searchResult = append(searchResult, SimilarKnowledgeSearchResult{
				Document: Document(hit.Source.Content),
				Score:    hit.Score,
			})
		}

		results = append(results, searchResult)
	}
	return results, nil
}
