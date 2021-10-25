package models

import "gorm.io/gorm"

type Price struct {
	gorm.Model
	Value        float32 `json:"value"`
	ProductRefer uint    `json:"productId"`
}
