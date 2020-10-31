package server

import (
	"context"

	"github.com/dokyan1989/g1/app/service1/internal/store"
	"github.com/dokyan1989/g1/app/service1/pb"
)

// ListProducts ...
func (s *Server) ListProducts(ctx context.Context, in *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := s.store.ListProducts(ctx, store.ListProductsParams{
		IDs:   in.Ids,
		Names: in.Names,
		Limit: in.Limit,
		Offset: in.Offset,
	})
	if err != nil {
		return nil, err
	}
	data := make([]*pb.Product, len(products))
	for i, p := range products {
		data[i] = &pb.Product{
			Id:    p.ID,
			Name:  p.Name,
			Price: p.Price,
		}
	}
	return &pb.ListProductsResponse{
		Data: data,
	}, nil
}

// CreateProduct ...
func (s *Server) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	lastId, err := s.store.CreateProduct(ctx, store.CreateProductParams{
		Name:  in.Name,
		Price: in.Price,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Id: lastId,
	}, nil
}
