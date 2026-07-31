package dtos

type ProjectCreateRequestDto struct {
	Name string `json:"name" binding:"required"`
}
