package dto

type RegisterDto struct {
	Name     string `json:"name" binding:"required" validation:"min=5"`
	Email    string `json:"email" binding:"required" validation:"endswith=.com"`
	Password string `json:"password" binding:"required" validation:"min=6"`
}
