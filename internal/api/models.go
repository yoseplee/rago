package api

type CommonResponse struct {
	Message string `json:"message"`
}

type IngestRequest struct {
	Documents []string `json:"documents"`
}

type RetrieveRequest struct {
	Query   string `json:"query" query:"query"`
	Message string `json:"message" query:"message"`
}
