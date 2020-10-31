package store

import "gorm.io/gorm"

// Product ...
type Product struct {
	gorm.Model
	Code  string
	Name  string
	Price uint
}
