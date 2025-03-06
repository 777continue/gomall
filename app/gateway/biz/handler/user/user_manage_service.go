package user

import (
	"context"

	"github.com/777continue/gomall/app/frontend/biz/service"
	"github.com/777continue/gomall/app/frontend/biz/utils"
	user "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListUser .
// @router /users [GET]
func ListUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ListUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.ListUserResp{}
	resp, err = service.NewListUserService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AddUser .
// @router /users [POST]
func AddUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.AddUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.AddUserResp{}
	resp, err = service.NewAddUserService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteUser .
// @router /users/{user_id} [DELETE]
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DeleteUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.DeleteUserResp{}
	resp, err = service.NewDeleteUserService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
