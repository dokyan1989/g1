package store

import (
	"context"

	"github.com/dokyan1989/g1/app/web1/internal/models"
)

// Querier defines operations on database
type Querier interface {
	GetProduct(ctx context.Context, id uint) (*models.Product, error)
	ListProducts(ctx context.Context, arg ListProductsParams) ([]models.Product, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) error
}
