package store

import (
	"context"
	"fmt"

	"github.com/dokyan1989/g1/app/web1/internal/models"
	"gorm.io/gorm"
)

// GetProduct ...
func (s *Store) GetProduct(ctx context.Context, id uint) (*models.Product, error) {
	var product models.Product

	err := s.db.WithContext(ctx).First(&product, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("The product with id = %d does not exist", id)
		}
		return nil, err
	}

	return &product, nil
}

// ListProductsParams ...
type ListProductsParams struct {
	Code string
}

// ListProducts ...
func (s *Store) ListProducts(ctx context.Context, arg ListProductsParams) ([]models.Product, error) {
	var products []models.Product
	db := s.db.WithContext(ctx)

	if arg.Code != "" {
		db = db.Where("code = ?", arg.Code)
	}

	err := db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("No product was found")
	}

	return products, nil
}

// CreateProductParams ...
type CreateProductParams struct {
	Code  string
	Price uint
}

// CreateProduct ...
func (s *Store) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	product := models.Product{
		Code:  arg.Code,
		Price: arg.Price,
	}

	err := s.db.Create(&product).Error
	if err != nil {
		return err
	}

	return nil
}
