package repository

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
