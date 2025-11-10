package tests

import (
	"arox-products/internal/handler"
	"arox-products/internal/models"
	mockcategories "arox-products/internal/stores/categories/mock"
	mockstores "arox-products/internal/stores/mock"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.uber.org/mock/gomock"
	"testing"
)

var (
	categories = []*models.Category{
		{
			Id:   1,
			Name: "Футболка",
			Slug: "T-shirt",
		},
		{
			Id:   2,
			Name: "Джинсы",
			Slug: "jeans",
		},
		{
			Id:   3,
			Name: "Платья",
			Slug: "dresses",
		},
		{
			Id:   4,
			Name: "Обувь",
			Slug: "shoes",
		},
	}
)

func TestListCategories(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mockstores.NewMockStores(ctrl)
	mockCategoriesStore := mockcategories.NewMockStore(ctrl)

	t.Run("success", func(t *testing.T) {
		mockCategoriesStore.EXPECT().ListCategories(gomock.Any()).Return(categories, nil)
		mockStore.EXPECT().Categories().Return(mockCategoriesStore)

		h := handler.NewHandler(mockStore)

		resp, err := h.ListCategories(ctx, &emptypb.Empty{})
		if err != nil {
			t.Fatalf("ListCategories(): %v", err)
		}

		for i := range resp.Categories {
			assert.Equal(t, categories[i].Id, resp.Categories[i].Id, "Category id should be equal")
			assert.Equal(t, categories[i].Name, resp.Categories[i].Name, "Category name should be equal")
			assert.Equal(t, categories[i].Slug, resp.Categories[i].Slug, "Category slug should be equal")
		}
	})

	t.Run("internal error", func(t *testing.T) {
		mockCategoriesStore.EXPECT().ListCategories(gomock.Any()).Return(nil, errors.New(`test error`))
		mockStore.EXPECT().Categories().Return(mockCategoriesStore)

		h := handler.NewHandler(mockStore)

		_, err := h.ListCategories(ctx, &emptypb.Empty{})

		assert.Equal(t, codes.Internal, status.Code(err))
	})
}
