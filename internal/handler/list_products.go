package handler

import (
	"arox-products/internal/models"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *handler) ListProducts(ctx context.Context, _ *emptypb.Empty) (*proto.ListProductsResponse, error) {
	response, err := h.store.Products().ListProducts(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "products not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	var products []*proto.Product

	for _, r := range response {

		var sizes *models.Sizes

		err = json.Unmarshal(r.Sizes, &sizes)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		var size []*proto.Size

		for _, s := range sizes.Sizes {
			size = append(size, &proto.Size{
				Size:  s.Size,
				Count: s.Count,
			})
		}

		var images []*proto.Image
		err = json.Unmarshal(r.Images, &images)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		product := &proto.Product{
			Id:          r.Id,
			Brand:       r.Brand,
			Name:        r.Name,
			Price:       r.Price,
			CategoryId:  r.CategoryId,
			Description: r.Description.String,
			Sizes:       size,
			IsActive:    r.IsActive,
			CreatedAt:   timestamppb.New(r.CreatedAt),
			Images:      images,
		}

		products = append(products, product)
	}

	return &proto.ListProductsResponse{
		Products: products,
	}, nil
}
