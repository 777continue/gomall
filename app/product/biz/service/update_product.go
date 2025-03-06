package service

import (
	"context"

	"github.com/777continue/gomall/app/product/biz/dal/mysql"
	"github.com/777continue/gomall/app/product/biz/model"
	product "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)

	updates := make(map[string]interface{})
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if len(req.Picture) > 0 {
		updates["picture"] = string(req.Picture)
	}
	if req.Price > 0 {
		updates["price"] = req.Price
	}
	if req.Stock > 0 {
		updates["stock"] = req.Stock
	}
	err = productQuery.UpdateProduct(int(req.Id), updates)
	if err != nil {
		return nil, err
	}

	return &product.UpdateProductResp{}, nil
}
