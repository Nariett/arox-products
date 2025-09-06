package handler

import (
	"arox-products/internal/stores"
	"context"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler interface {
	proto.ProductsServiceServer

	GetProduct(ctx context.Context, req *proto.GetProductRequest) (*proto.GetProductResponse, error)
	ListProducts(ctx context.Context, _ *emptypb.Empty) (*proto.ListProductsResponse, error)

	ListCategories(ctx context.Context, _ *emptypb.Empty) (*proto.ListCategoriesResponse, error)
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
