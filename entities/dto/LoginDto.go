package dto

type LoginDto struct {
	Email    string `json:"email" binding:"required, endswith=.com"`
	Password string `json:"password" binding:"required, min=6"`
}
