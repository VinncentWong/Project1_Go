package dto

type RegisterDto struct {
	Name     string `json:"name" binding:"required,min=5"`
	Email    string `json:"email" binding:"required,endswith=.com"`
	Password string `json:"password" binding:"required,min=6"`
}
