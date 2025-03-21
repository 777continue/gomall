package service

import (
	"context"

	product "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/product"
	rpcproduct "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	product_client "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

// var p *rpcproduct.Product
func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	p, err := product_client.Client.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"item": p.Product,
	}, nil
}
