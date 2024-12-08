package service

import (
	"context"

	common "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	product_client "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	resp, err := product_client.Client.ListProducts(h.Context, &product.ListProductsReq{})

	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Hot sale",
		"items": resp.Products,
	}, nil
}
