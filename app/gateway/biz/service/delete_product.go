package service

import (
	"context"

	MNproduct "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/MNproduct"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	rpc_product "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteProductService(Context context.Context, RequestContext *app.RequestContext) *DeleteProductService {
	return &DeleteProductService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteProductService) Run(req *MNproduct.DeleteProductReq) (resp *MNproduct.DeleteProductResp, err error) {
	_, err = rpc_product.Client.DeleteProduct(h.Context, &product.DeleteProductReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &MNproduct.DeleteProductResp{}, nil
}
