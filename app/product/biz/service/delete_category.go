package service

import (
	"context"

	"github.com/777continue/gomall/app/product/biz/dal/mysql"
	"github.com/777continue/gomall/app/product/biz/model"
	product "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
)

type DeleteCategoryService struct {
	ctx context.Context
} // NewDeleteCategoryService new DeleteCategoryService
func NewDeleteCategoryService(ctx context.Context) *DeleteCategoryService {
	return &DeleteCategoryService{ctx: ctx}
}

// Run create note info
func (s *DeleteCategoryService) Run(req *product.DeleteCategoryReq) (resp *product.DeleteCategoryResp, err error) {
	// Finish your business logic.
	CategoryQuery := *model.NewCategoryQuery(s.ctx, mysql.DB)
	CategoryQuery.DeleteCategory(req.Name)
	return &product.DeleteCategoryResp{}, nil
}
