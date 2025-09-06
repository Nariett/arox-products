package handler

import (
	"arox-products/internal/stores"
	"context"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler interface {
	proto.ProductsServiceServer

	ListProducts(ctx context.Context, empty *emptypb.Empty) (*proto.Products, error)
}

type handler struct {
	proto.UnimplementedProductsServiceServer
	store stores.Stores
}

func NewHandler(store stores.Stores) Handler {
	return &handler{
		store: store,
	}
}
