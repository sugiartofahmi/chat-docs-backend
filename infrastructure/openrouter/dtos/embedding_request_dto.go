package dtos

type EmbeddingRequestDto struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}
