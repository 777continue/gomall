package service

import (
	"context"

	MNproduct "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/MNproduct"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	rpc_product "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductService {
	return &UpdateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductService) Run(req *MNproduct.UpdateProductReq) (resp *MNproduct.UpdateProductResp, err error) {
	_, err = rpc_product.Client.UpdateProduct(h.Context, &product.UpdateProductReq{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Stock:       req.Stock,
		Categories:  req.Categories,
	})
	if err != nil {
		return nil, err
	}
	return &MNproduct.UpdateProductResp{}, nil
}
