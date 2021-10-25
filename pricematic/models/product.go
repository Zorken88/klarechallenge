package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model  `swaggerignore:"true"`
	Name        string  `json:"name" `
	ActualPrice float32 `json:"price"`
	Prices      []Price `gorm:"foreignKey:ProductRefer" swaggerignore:"true"`
}
