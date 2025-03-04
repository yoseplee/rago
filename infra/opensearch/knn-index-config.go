package opensearch

type Method struct {
	Name       string `json:"name"`
	SpaceType  string `json:"space_type"`
	Engine     string `json:"engine"`
	Parameters struct {
		EfConstruction int `json:"ef_construction"`
		M              int `json:"m"`
	} `json:"parameters"`
}

type Embedding struct {
	Type      string `json:"type"`
	Dimension int    `json:"dimension"`
	Method    Method `json:"method"`
}

type Properties struct {
	Embedding Embedding `json:"embedding"`
}

type Mappings struct {
	Properties Properties `json:"properties"`
}

type IndexSettings struct {
	Knn                  bool `json:"knn"`
	KnnAlgoParamEfSearch int  `json:"knn.algo_param.ef_search"`
}

type Settings struct {
	Index IndexSettings `json:"index"`
}

type KNNIndexConfig struct {
	Settings Settings `json:"settings"`
	Mappings Mappings `json:"mappings"`
}
