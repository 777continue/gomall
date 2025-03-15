package service

import (
	"context"

	checkout "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/checkout"
	frontendutils "github.com/777continue/gomall/app/frontend/utils"
	rpccheckout "github.com/777continue/gomall/rpc_gen/kitex_gen/checkout"
	checkout_client "github.com/777continue/gomall/rpc_gen/rpc/checkout"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	_, err = checkout_client.Client.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:  uint32(userId),
		Address: req.Address,
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"message": "success",
	}, nil
}
