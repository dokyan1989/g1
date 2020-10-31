package store

import (
	"context"
	"database/sql"
)

type Store interface {
	ListProducts(ctx context.Context, params ListProductsParams) ([]Product, error)
	CreateProduct(ctx context.Context, params CreateProductParams) (uint64, error)
}

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{db}
}
