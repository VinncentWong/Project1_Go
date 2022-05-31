package entities

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	Name     string `json:"name" binding:"required,min=5"`
	Synopsis string `json:"synopsis"`
	Price    uint64 `json:"price" binding:"required,gt=0"`
	Stock    uint64 `json:"stock" binding:"required,gt=0"`
}
