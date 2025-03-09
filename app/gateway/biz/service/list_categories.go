package service

import (
	"context"

	MNproduct "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/MNproduct"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	rpc_product "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListCategoriesService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListCategoriesService(Context context.Context, RequestContext *app.RequestContext) *ListCategoriesService {
	return &ListCategoriesService{RequestContext: RequestContext, Context: Context}
}

func (h *ListCategoriesService) Run(req *MNproduct.ListCategoriesReq) (resp *MNproduct.ListCategoriesResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	rpcResp, err := rpc_product.Client.ListCategories(h.Context, &product.ListCategoriesReq{})
	if err != nil {
		return nil, err
	}
	return &MNproduct.ListCategoriesResp{
		Categories: rpcResp.Categories,
	}, nil
}
