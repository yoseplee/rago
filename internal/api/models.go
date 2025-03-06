package api

type CommonResponse struct {
	Message string `json:"message"`
}

type IngestRequest struct {
	Document string `json:"document"`
}

type RetrieveRequest struct {
	Query string `json:"query" query:"query"`
}
