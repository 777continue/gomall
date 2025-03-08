package service

import (
	"context"

	MNproduct "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/MNproduct"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	rpc_product "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddProductService(Context context.Context, RequestContext *app.RequestContext) *AddProductService {
	return &AddProductService{RequestContext: RequestContext, Context: Context}
}

func (h *AddProductService) Run(req *MNproduct.AddProductReq) (resp *MNproduct.AddProductResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc_product.Client.AddProduct(h.Context, &product.AddProductReq{
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
	return &MNproduct.AddProductResp{}, nil
}
