package service

import (
	"context"

	"github.com/777continue/gomall/app/product/biz/dal/mysql"
	"github.com/777continue/gomall/app/product/biz/model"
	product "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
)

type AddProductService struct {
	ctx context.Context
} // NewAddProductService new AddProductService
func NewAddProductService(ctx context.Context) *AddProductService {
	return &AddProductService{ctx: ctx}
}

// Run create note info
func (s *AddProductService) Run(req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	// Finish your business logic.
	ProductQuery := *model.NewProductQuery(s.ctx, mysql.DB)
	ProductQuery.AddProduct(&model.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     string(req.Picture),
		Price:       req.Price,
		Stock:       req.Stock,
	})
	return &product.AddProductResp{}, nil
}
