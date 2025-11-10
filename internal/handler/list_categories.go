package handler

import (
	"context"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handler) ListCategories(ctx context.Context, _ *emptypb.Empty) (*proto.ListCategoriesResponse, error) {
	response, err := h.store.Categories().ListCategories(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var categories []*proto.Category

	for _, r := range response {
		category := &proto.Category{
			Id:   r.Id,
			Name: r.Name,
			Slug: r.Slug,
		}

		categories = append(categories, category)
	}

	return &proto.ListCategoriesResponse{
		Categories: categories,
	}, nil
}
