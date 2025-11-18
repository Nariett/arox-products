package tests

import (
	"arox-products/internal/handler"
	"arox-products/internal/models"
	mockstores "arox-products/internal/stores/mock"
	mockproducts "arox-products/internal/stores/products/mock"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"

	"database/sql"
	"testing"
)

var (
	products = []*models.ProductWithImage{
		{
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
			CreatedAt: time.Date(2025, 10, 10, 10, 10, 0, 0, time.UTC),
			Images: []byte(
				`[
			  {
				"id": 1,
				"url": "http://example.com/image1.png",
				"is_main": true,
				"is_active": true
			  },
			  {
				"id": 2,
				"url": "http://example.com/image1.png",
				"is_main": false,
				"is_active": true
			  }
			]`),
		},
		{
			Id:          1,
			Brand:       "Adidas",
			Name:        "T-Shirt HYD-124AS",
			CategoryId:  1,
			Price:       200,
			Description: sql.NullString{String: "T-Shirt HYD-124AS", Valid: true},
			Sizes: []byte(`
			{
			  "sizes": [
				{
				  "size": "S",
				  "count": 2
				},
				{
				  "size": "M", 
				  "count": 2
				},
				{
				  "size": "L",
				  "count": 5
				},
				{
				  "size": "XL",
				  "count": 1
				}
			  ]
			}
		`),
			IsActive:  true,
			CreatedAt: time.Date(2025, 8, 14, 12, 00, 0, 0, time.UTC),
			Images: []byte(
				`[
			  {
				"id": 1,
				"url": "http://example.com/image1.png",
				"is_main": true,
				"is_active": true
			  },
			  {
				"id": 2,
				"url": "http://example.com/image1.png",
				"is_main": false,
				"is_active": true
			  }
			]`),
		},
		{
			Id:          1,
			Brand:       "Carhartt",
			Name:        "Jacket Aviator",
			CategoryId:  4,
			Price:       145,
			Description: sql.NullString{String: "Jacket Aviator", Valid: true},
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
			CreatedAt: time.Date(2025, 5, 25, 12, 10, 0, 0, time.UTC),
			Images: []byte(
				`[
			  {
				"id": 1,
				"url": "http://example.com/image1.png",
				"is_main": true,
				"is_active": true
			  },
			  {
				"id": 2,
				"url": "http://example.com/image1.png",
				"is_main": false,
				"is_active": true
			  }
			]`),
		},
	}
)

func TestListProducts(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mockstores.NewMockStores(ctrl)
	mockProducts := mockproducts.NewMockStore(ctrl)

	t.Run("success", func(t *testing.T) {
		mockProducts.EXPECT().ListProducts(gomock.Any()).Return(products, nil)
		mockStore.EXPECT().Products().Return(mockProducts)

		h := handler.NewHandler(mockStore)

		resp, err := h.ListProducts(ctx, &emptypb.Empty{})
		if err != nil {
			t.Fatalf("ListProducts(): %v", err)
		}
		for i := range resp.Products {
			assert.Equal(t, products[i].Id, resp.Products[i].Id)
			assert.Equal(t, products[i].Brand, resp.Products[i].Brand)
			assert.Equal(t, products[i].Name, resp.Products[i].Name)
			assert.Equal(t, products[i].Price, resp.Products[i].Price)
			assert.Equal(t, products[i].CategoryId, resp.Products[i].CategoryId)
			assert.Equal(t, products[i].Description.String, resp.Products[i].Description)
			assert.Equal(t, products[i].IsActive, resp.Products[i].IsActive)

			assert.NotZero(t, products[i].CreatedAt, resp.Products[i].CreatedAt)
			assert.NotZero(t, resp.Products[i].Images)
			assert.NotZero(t, resp.Products[i].Sizes)
		}
	})

	t.Run("error not found", func(t *testing.T) {
		mockProducts.EXPECT().ListProducts(gomock.Any()).Return(nil, sql.ErrNoRows)
		mockStore.EXPECT().Products().Return(mockProducts)

		h := handler.NewHandler(mockStore)

		resp, err := h.ListProducts(ctx, &emptypb.Empty{})
		assert.Equal(t, codes.NotFound, status.Code(err))
		assert.Nil(t, resp)

	})

	t.Run("error internal", func(t *testing.T) {
		mockProducts.EXPECT().ListProducts(gomock.Any()).Return(nil, errors.New(`internal error`))
		mockStore.EXPECT().Products().Return(mockProducts)

		h := handler.NewHandler(mockStore)

		resp, err := h.ListProducts(ctx, &emptypb.Empty{})
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Nil(t, resp)
	})

}
