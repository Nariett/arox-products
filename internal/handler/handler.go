package handler

import (
	"context"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler interface {
	proto.ProductsServiceServer

	ListProducts(ctx context.Context, empty *emptypb.Empty) (*proto.Product, error)
}

type handler struct {
	proto.UnimplementedProductsServiceServer
}

func NewHandler() Handler {
	return &handler{}
}
