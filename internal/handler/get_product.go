package handler

import (
	"context"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *handler) GetProduct(ctx context.Context, req *proto.GetProductRequest) (*proto.GetProductResponse, error) {
	product, err := h.store.Products().GetProductWithId(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	response := &proto.Product{
		Id:          product.ID,
		Brand:       product.Brand,
		Name:        product.Name,
		Price:       product.Price,
		CategoryId:  product.CategoryId,
		Description: product.Description.String,
		Sizes:       string(product.Sizes),
		IsActive:    product.IsActive,
		CreatedAt:   timestamppb.New(product.CreatedAt),
	}

	return &proto.GetProductResponse{
		Product: response,
	}, nil
}
