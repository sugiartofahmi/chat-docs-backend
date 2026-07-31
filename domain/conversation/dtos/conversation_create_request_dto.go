package dtos

type ConversationCreateRequestDto struct {
	Content string `json:"content" binding:"required"`
}
