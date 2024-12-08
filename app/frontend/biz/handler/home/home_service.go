package home

import (
	"context"

	"github.com/777continue/gomall/app/frontend/biz/service"
	"github.com/777continue/gomall/app/frontend/biz/utils"
	common "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/common"
	frontendUtils "github.com/777continue/gomall/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Home .
// @router / [GET]

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
	hlog.Debug(resp)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	//midware resp
	resp["user_id"] = ctx.Value(frontendUtils.SessionUserId)

	c.HTML(consts.StatusOK, "home", resp)

}
