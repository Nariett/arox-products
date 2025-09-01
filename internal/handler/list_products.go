package handler

import (
	"context"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handler) ListProducts(context.Context, *emptypb.Empty) (*proto.Product, error) {
	return &proto.Product{
		Brand:       "Nike",
		Name:        "Air Force",
		Price:       123,
		Description: "Default white shoes",
	}, nil
}
