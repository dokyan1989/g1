package store

import "gorm.io/gorm"

// Store implement Querier interface
type Store struct {
	db *gorm.DB
}

// NewStore creates a Store instance
func NewStore(db *gorm.DB) *Store {
	return &Store{db}
}
