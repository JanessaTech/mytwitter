package dto

type PostDTO struct {
	Name    string `json:"name" binding:"required"`
	Message string `json:"message" binding:"required"`
}
