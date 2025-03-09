package service

import (
	"context"

	"github.com/777continue/gomall/app/product/biz/dal/mysql"
	"github.com/777continue/gomall/app/product/biz/model"
	product "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
)

type AddCategoryService struct {
	ctx context.Context
} // NewAddCategoryService new AddCategoryService
func NewAddCategoryService(ctx context.Context) *AddCategoryService {
	return &AddCategoryService{ctx: ctx}
}

// Run create note info
func (s *AddCategoryService) Run(req *product.AddCategoryReq) (resp *product.AddCategoryResp, err error) {
	// Finish your business logic.
	CategoryQuery := *model.NewCategoryQuery(s.ctx, mysql.DB)
	CategoryQuery.AddCategory(req.Name)
	return &product.AddCategoryResp{}, nil
}
