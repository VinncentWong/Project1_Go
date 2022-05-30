package dto

type LoginDto struct {
	Email    string `json:"email" binding:"required" validation:"endswith=.com"`
	Password string `json:"password" binding:"required" validation:"min=6"`
}
