package service

import (
	"context"

	auth "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/user"
	rpc_user "github.com/777continue/gomall/rpc_gen/rpc/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	userResp, err := rpc_user.Client.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userResp.UserId) //resp.Id
	err = session.Save()
	if err != nil {
		return nil, err
	}
	// empty struct -> complier don't give error
	return
}
