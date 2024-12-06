package service

import (
	"context"

	cart "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/common"
	frontendUtils "github.com/777continue/gomall/app/frontend/utils"
	rpccart "github.com/777continue/gomall/rpc_gen/kitex_gen/cart"
	cart_client "github.com/777continue/gomall/rpc_gen/rpc/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	_, err = cart_client.Client.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  uint32(req.ProductNum),
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
