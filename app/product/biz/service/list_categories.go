package service

import (
	"context"

	"github.com/777continue/gomall/app/product/biz/dal/mysql"
	"github.com/777continue/gomall/app/product/biz/model"
	product "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
)

type ListCategoriesService struct {
	ctx context.Context
} // NewListCategoriesService new ListCategoriesService
func NewListCategoriesService(ctx context.Context) *ListCategoriesService {
	return &ListCategoriesService{ctx: ctx}
}

// Run create note info
func (s *ListCategoriesService) Run(req *product.ListCategoriesReq) (resp *product.ListCategoriesResp, err error) {
	// Finish your business logic.
	CategoryQuery := *model.NewCategoryQuery(s.ctx, mysql.DB)
	categories, err := CategoryQuery.ListCategories()
	return &product.ListCategoriesResp{Categories: categories}, nil
}
