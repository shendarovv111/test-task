package api

type EncryptRequest struct {
	Input     string `json:"input" binding:"required"`
	Algorithm string `json:"algorithm" binding:"required"`
}

type EncryptResponse struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

type AlgorithmsResponse struct {
	Algorithms []string `json:"algorithms"`
}
