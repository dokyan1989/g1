package store

// Store ...
type Store interface {
	// Common
	Ping() error
	Close() error

	// Product
	SearchProducts(params *SearchProductsParams) ([]*Product, error)
	CreateProduct(params *CreateProductParams) (uint64, error)
	GetProduct(id uint64) (*Product, error)
	UpdateProduct(params *UpdateProductParams) error
	DeleteProduct(id uint64) error
}
