package service

import (
	"context"

	MNproduct "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/MNproduct"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	rpc_product "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCategoryService(Context context.Context, RequestContext *app.RequestContext) *AddCategoryService {
	return &AddCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCategoryService) Run(req *MNproduct.AddCategoryReq) (resp *MNproduct.AddCategoryResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	_, err = rpc_product.Client.AddCategory(h.Context, &product.AddCategoryReq{
		Name: req.Category,
	})
	if err != nil {
		return nil, err
	}
	return &MNproduct.AddCategoryResp{}, nil
}
