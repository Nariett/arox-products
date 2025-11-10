package tests

//func TestListProducts(t *testing.T) {
//	ctx := context.Background()
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockStore := mockstores.NewMockStores(ctrl)
//	mockCategoriesStore := mockcategories.NewMockStore(ctrl)
//
//	t.Run("success", func(t *testing.T) {
//		mockCategoriesStore.EXPECT().ListCategories(gomock.Any()).Return(categories, nil)
//		mockStore.EXPECT().Categories().Return(mockCategoriesStore)
//
//		h := handler.NewHandler(mockStore)
//
//		resp, err := h.ListCategories(ctx, &emptypb.Empty{})
//		if err != nil {
//			t.Fatalf("ListCategories(): %v", err)
//		}
//		for i := range resp.Categories {
//			assert.Equal(t, categories[i].Id, resp.Categories[i].Id, "Category id should be equal")
//			assert.Equal(t, categories[i].Name, resp.Categories[i].Name, "Category name should be equal")
//			assert.Equal(t, categories[i].Slug, resp.Categories[i].Slug, "Category slug should be equal")
//		}
//	})
//}
