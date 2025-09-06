package handler

import (
	"context"
	"encoding/json"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *handler) ListProducts(ctx context.Context, _ *emptypb.Empty) (*proto.ListProductsResponse, error) {
	response, err := h.store.Products().ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	var products []*proto.Product

	for _, r := range response {

		var sizes *proto.Sizes

		err := json.Unmarshal(r.Sizes, &sizes)
		if err != nil {
			return nil, err
		}

		product := &proto.Product{
			Id:          r.ID,
			Brand:       r.Brand,
			Name:        r.Name,
			Price:       r.Price,
			CategoryId:  r.CategoryId,
			Description: r.Description.String,
			Sizes:       sizes,
			IsActive:    r.IsActive,
			CreatedAt:   timestamppb.New(r.CreatedAt),
		}

		products = append(products, product)
	}

	return &proto.ListProductsResponse{
		Products: products,
	}, nil
}
