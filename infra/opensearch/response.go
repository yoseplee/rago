package opensearch

type Response struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Id     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Embedding []float64 `json:"embedding"`
				Dimension int       `json:"dimension"`
				Content   string    `json:"content"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
