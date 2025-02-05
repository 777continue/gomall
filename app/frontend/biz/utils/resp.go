package utils

import (
	"context"

	frontendUtils "github.com/777continue/gomall/app/frontend/utils"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/cart"
	cart_client "github.com/777continue/gomall/rpc_gen/rpc/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, resp map[string]any) map[string]any {
	userId := frontendUtils.GetUserIdFromCtx(ctx)
	resp["user_id"] = userId

	if userId > 0 {
		cartResp, err := cart_client.Client.GetCart(ctx, &cart.GetCartReq{
			UserId: uint32(userId),
		})
		if err == nil && cartResp != nil {
			resp["cart_num"] = len(cartResp.Items)
		}
	}

	return resp
}
