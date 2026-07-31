package dtos

type ChatMessageDto struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequestDto struct {
	Model    string           `json:"model"`
	Messages []ChatMessageDto `json:"messages"`
}
