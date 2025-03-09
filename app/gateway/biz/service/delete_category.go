package service

import (
	"context"

	MNproduct "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/MNproduct"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	rpc_product "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteCategoryService(Context context.Context, RequestContext *app.RequestContext) *DeleteCategoryService {
	return &DeleteCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteCategoryService) Run(req *MNproduct.DeleteCategoryReq) (resp *MNproduct.DeleteCategoryResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc_product.Client.DeleteCategory(h.Context, &product.DeleteCategoryReq{
		Name: req.Category,
	})
	if err != nil {
		return nil, err
	}
	return &MNproduct.DeleteCategoryResp{}, nil
}
