package models

type OllamaRequest struct {
	Message string `json:"message"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}
