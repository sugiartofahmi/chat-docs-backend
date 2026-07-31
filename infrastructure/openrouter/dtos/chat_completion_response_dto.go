package dtos

type ChatCompletionChoiceDto struct {
	Message ChatMessageDto `json:"message"`
}

type ChatCompletionResponseDto struct {
	Choices []ChatCompletionChoiceDto `json:"choices"`
}
