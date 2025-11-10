package handler

import (
	"arox-products/internal/models"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *handler) GetProduct(ctx context.Context, req *proto.GetProductRequest) (*proto.GetProductResponse, error) {
	product, err := h.store.Products().GetProductWithId(ctx, req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("product with id %d not found", req.Id))
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	images, err := h.store.Images().GetImagesWithId(ctx, req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("images with product_id %d not found", req.Id))
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	var protoImages []*proto.Image

	for _, image := range images {
		protoImages = append(protoImages, &proto.Image{
			Id:        image.Id,
			IdProduct: image.IdProduct,
			Url:       image.Url,
			IsMain:    image.IsMain,
			IsActive:  image.IsActive,
		})
	}

	var sizes *models.Sizes

	err = json.Unmarshal(product.Sizes, &sizes)
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

	response := &proto.Product{
		Id:          product.Id,
		Brand:       product.Brand,
		Name:        product.Name,
		Price:       product.Price,
		CategoryId:  product.CategoryId,
		Description: product.Description.String,
		Sizes:       size,
		Images:      protoImages,
		IsActive:    product.IsActive,
		CreatedAt:   timestamppb.New(product.CreatedAt),
	}

	return &proto.GetProductResponse{
		Product: response,
	}, nil
}
