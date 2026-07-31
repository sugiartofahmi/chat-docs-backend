package dtos

type EmbeddingResponseDto struct {
	Object string             `json:"object"`
	Data   []EmbeddingDataDto `json:"data"`
	Model  string             `json:"model"`
}

type EmbeddingDataDto struct {
	Object    string    `json:"object"`
	Index     int       `json:"index"`
	Embedding []float64 `json:"embedding"`
}
