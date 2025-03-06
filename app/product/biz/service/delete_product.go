package service

import (
	"context"

	"github.com/777continue/gomall/app/product/biz/dal/mysql"
	"github.com/777continue/gomall/app/product/biz/model"
	product "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	productQuery.DeleteProduct(int(req.Id))
	resp = &product.DeleteProductResp{}
	return resp, nil
}
