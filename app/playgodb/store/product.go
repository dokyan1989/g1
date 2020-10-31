package store

// SearchProductsParams ...
type SearchProductsParams struct{}

// SearchProducts ...
func (s *SQLStore) SearchProducts(params *SearchProductsParams) ([]*Product, error) {
	return []*Product{}, nil
}

// CreateProductParams ...
type CreateProductParams struct{}

// CreateProduct ...
func (s *SQLStore) CreateProduct(params *CreateProductParams) (uint64, error) {
	return 0, nil
}

// GetProduct ...
func (s *SQLStore) GetProduct(id uint64) (*Product, error) {
	return nil, nil
}

// UpdateProductParams ...
type UpdateProductParams struct{}

// UpdateProduct ...
func (s *SQLStore) UpdateProduct(params *UpdateProductParams) error {
	return nil
}

// DeleteProduct ...
func (s *SQLStore) DeleteProduct(id uint64) error {
	return nil
}
