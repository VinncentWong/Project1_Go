package entities

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	Name     string `json:"name" binding:"required" validation:"min=5"`
	Synopsis string `json:"synopsis"`
	Price    uint64 `json:"price" binding:"required" validation:"gt=0"`
	Stock    uint64 `json:"stock" binding:"required" validation:"gt=0"`
}
