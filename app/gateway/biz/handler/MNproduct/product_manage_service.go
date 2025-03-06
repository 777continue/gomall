package MNproduct

import (
	"context"

	"github.com/777continue/gomall/app/frontend/biz/service"
	"github.com/777continue/gomall/app/frontend/biz/utils"
	MNproduct "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/MNproduct"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// AddProduct .
// @router /products [POST]
func AddProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req MNproduct.AddProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &MNproduct.AddProductResp{}
	resp, err = service.NewAddProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateProduct .
// @router /products/{id} [PUT]
func UpdateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req MNproduct.UpdateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &MNproduct.UpdateProductResp{}
	resp, err = service.NewUpdateProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteProduct .
// @router /products/{id} [DELETE]
func DeleteProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req MNproduct.DeleteProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &MNproduct.DeleteProductResp{}
	resp, err = service.NewDeleteProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
