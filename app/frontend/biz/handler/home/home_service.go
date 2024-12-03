package home

import (
	"context"

	"github.com/777continue/gomall/app/frontend/biz/service"
	"github.com/777continue/gomall/app/frontend/biz/utils"
	common "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Home .
// @router / [GET]
type SessionUserIdKey string

const (
	SessionUserId SessionUserIdKey = "user_id"
)

func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//service resp
	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	//midware resp
	resp["user_id"] = ctx.Value(SessionUserId)
	
	c.HTML(consts.StatusOK, "home", resp)

	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
func Test(ctx context.Context, c *app.RequestContext) {
	resp := map[string]any{"Title": "Test"}
	c.HTML(consts.StatusOK, "test", resp)
}
