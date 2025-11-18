package tests

import (
	"arox-products/internal/handler"
	"arox-products/internal/models"
	mockcategories "arox-products/internal/stores/categories/mock"
	mockstores "arox-products/internal/stores/mock"
	"context"
	"database/sql"
	"errors"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

var category = &models.Category{
	Id:   1,
	Name: "Футболка",
	Slug: "T-shirt",
}

func TestGetCategory(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mockstores.NewMockStores(ctrl)
	mockCategoriesStore := mockcategories.NewMockStore(ctrl)

	t.Run("success", func(t *testing.T) {
		mockCategoriesStore.EXPECT().GetCategoryWithId(gomock.Any(), int64(1)).Return(category, nil)
		mockStore.EXPECT().Categories().Return(mockCategoriesStore)

		h := handler.NewHandler(mockStore)

		resp, err := h.GetCategory(ctx, &proto.GetCategoryRequest{Id: 1})
		if err != nil {
			t.Fatalf("GetCategory(): %v", err)
		}

		assert.Equal(t, category.Id, resp.Category.Id)
		assert.Equal(t, category.Name, resp.Category.Name)
		assert.Equal(t, category.Slug, resp.Category.Slug)
	})

	t.Run("error not found", func(t *testing.T) {
		mockCategoriesStore.EXPECT().GetCategoryWithId(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)
		mockStore.EXPECT().Categories().Return(mockCategoriesStore)

		h := handler.NewHandler(mockStore)

		resp, err := h.GetCategory(ctx, &proto.GetCategoryRequest{Id: 100})
		assert.Equal(t, codes.NotFound, status.Code(err))
		assert.Nil(t, resp)

	})

	t.Run("error internal", func(t *testing.T) {
		mockCategoriesStore.EXPECT().GetCategoryWithId(gomock.Any(), gomock.Any()).Return(nil, errors.New(`internal error`))
		mockStore.EXPECT().Categories().Return(mockCategoriesStore)

		h := handler.NewHandler(mockStore)

		resp, err := h.GetCategory(ctx, &proto.GetCategoryRequest{Id: 100})
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Nil(t, resp)
	})
}
