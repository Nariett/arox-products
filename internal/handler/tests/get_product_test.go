package tests

import (
	"arox-products/internal/handler"
	"arox-products/internal/models"
	mockimages "arox-products/internal/stores/images/mock"
	mockstores "arox-products/internal/stores/mock"
	mockproducts "arox-products/internal/stores/products/mock"
	"context"
	"database/sql"
	"github.com/Nariett/arox-pkg/grpc/pb/products"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

var (
	timeNow       = time.Now()
	productFromDB = &models.ProductWithImage{
		Id:          1,
		Brand:       "Nike",
		Name:        "T-Shirt HV9803-110",
		CategoryId:  1,
		Price:       145,
		Description: sql.NullString{String: "T-Shirt HV9803-110", Valid: true},
		Sizes: []byte(`
			{
			  "sizes": [
				{
				  "size": "S",
				  "count": 10
				},
				{
				  "size": "M", 
				  "count": 15
				},
				{
				  "size": "L",
				  "count": 8
				}
			  ]
			}
		`),
		IsActive:  true,
		CreatedAt: timeNow,
	}

	errProductFromDB = &models.ProductWithImage{
		Id:          1,
		Brand:       "Nike",
		Name:        "T-Shirt HV9803-110",
		CategoryId:  1,
		Price:       145,
		Description: sql.NullString{String: "T-Shirt HV9803-110", Valid: true},
		Sizes: []byte(`
			{
			  "sizes": [
				{
				  size": "S",
				  "count": 10
				},
				{
				  "size": "M", 
				  "count": 15
				},
				{
				  "size": "L",
				  "count": 8
				}
			  ]
			}
		`),
		IsActive:  true,
		CreatedAt: timeNow,
	}

	imageFromDB = []*models.Image{
		{
			Id:        1,
			IdProduct: 1,
			Url:       "http://example.com/image1.png",
			IsMain:    true,
			IsActive:  true,
		},
		{
			Id:        2,
			IdProduct: 1,
			Url:       "http://example.com/image2.png",
			IsMain:    true,
			IsActive:  true,
		},
	}

	product = &products.GetProductResponse{
		Product: &products.Product{
			Id:          1,
			Brand:       "Nike",
			Name:        "T-Shirt HV9803-110",
			CategoryId:  1,
			Price:       145,
			Description: "T-Shirt HV9803-110",
			Sizes: []*products.Size{
				{
					Size:  "S",
					Count: 11,
				},
				{
					Size:  "M",
					Count: 6,
				},
				{
					Size:  "L",
					Count: 10,
				},
			},
			IsActive:  true,
			CreatedAt: timestamppb.New(timeNow),
			Images: []*products.Image{
				{
					Id:        1,
					IdProduct: 1,
					Url:       "http://example.com/image1.png",
					IsMain:    true,
				},
				{
					Id:        2,
					IdProduct: 1,
					Url:       "http://example.com/image2.png",
					IsMain:    false,
				},
			},
		},
	}
)

func TestGetProduct(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mockstores.NewMockStores(ctrl)
	mockProductsStore := mockproducts.NewMockStore(ctrl)
	mockImagesStore := mockimages.NewMockStore(ctrl)

	t.Run("success", func(t *testing.T) {
		mockProductsStore.EXPECT().GetProductWithId(gomock.Any(), int64(1)).Return(productFromDB, nil)
		mockStore.EXPECT().Products().Return(mockProductsStore)

		mockImagesStore.EXPECT().GetImagesWithId(gomock.Any(), int64(1)).Return(imageFromDB, nil)
		mockStore.EXPECT().Images().Return(mockImagesStore)

		h := handler.NewHandler(mockStore)

		resp, err := h.GetProduct(ctx, &products.GetProductRequest{Id: 1})
		if err != nil {
			t.Fatalf("GetProduct(): %v", err)
		}

		assert.Equal(t, product.Product.Id, resp.Product.Id)
		assert.Equal(t, product.Product.Brand, resp.Product.Brand)
		assert.Equal(t, product.Product.Name, resp.Product.Name)
		assert.Equal(t, product.Product.CategoryId, resp.Product.CategoryId)
		assert.Equal(t, product.Product.Price, resp.Product.Price)
		assert.Equal(t, product.Product.Description, resp.Product.Description)
		assert.Equal(t, len(product.Product.Sizes), len(resp.Product.Sizes))
		assert.Equal(t, product.Product.IsActive, resp.Product.IsActive)
		assert.Equal(t, product.Product.CreatedAt, resp.Product.CreatedAt)
	})

	t.Run("error product not found", func(t *testing.T) {
		mockProductsStore.EXPECT().GetProductWithId(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)
		mockStore.EXPECT().Products().Return(mockProductsStore)

		h := handler.NewHandler(mockStore)

		resp, err := h.GetProduct(ctx, &products.GetProductRequest{Id: 100})
		assert.Equal(t, codes.NotFound, status.Code(err))
		assert.Nil(t, resp)
	})

	t.Run("error image not found", func(t *testing.T) {
		mockProductsStore.EXPECT().GetProductWithId(gomock.Any(), gomock.Any()).Return(productFromDB, nil)
		mockStore.EXPECT().Products().Return(mockProductsStore)

		mockImagesStore.EXPECT().GetImagesWithId(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)
		mockStore.EXPECT().Images().Return(mockImagesStore)

		h := handler.NewHandler(mockStore)

		resp, err := h.GetProduct(ctx, &products.GetProductRequest{Id: 100})
		assert.Equal(t, codes.NotFound, status.Code(err))
		assert.Nil(t, resp)
	})

	t.Run("error with unmarshal", func(t *testing.T) {
		mockProductsStore.EXPECT().GetProductWithId(gomock.Any(), int64(1)).Return(errProductFromDB, nil)
		mockStore.EXPECT().Products().Return(mockProductsStore)

		mockImagesStore.EXPECT().GetImagesWithId(gomock.Any(), int64(1)).Return(imageFromDB, nil)
		mockStore.EXPECT().Images().Return(mockImagesStore)

		h := handler.NewHandler(mockStore)

		resp, err := h.GetProduct(ctx, &products.GetProductRequest{Id: 1})
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Nil(t, resp)
	})
}
