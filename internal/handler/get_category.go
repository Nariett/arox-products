package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetCategory(ctx context.Context, req *proto.GetCategoryRequest) (*proto.GetCategoryResponse, error) {
	category, err := h.store.Categories().GetCategoryWithId(ctx, req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("category with id %d not found", req.Id))
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetCategoryResponse{
		Category: &proto.Category{
			Id:   category.Id,
			Name: category.Name,
			Slug: category.Slug,
		},
	}, nil
}
