package models

import "gorm.io/gorm"

// Product ...
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// TableName ...
func (Product) TableName() string {
	return "product"
}
