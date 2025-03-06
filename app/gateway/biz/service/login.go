package service

import (
	"context"

	auth "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/user"
	rpc_user "github.com/777continue/gomall/rpc_gen/rpc/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (userResp *user.LoginResp, redirect string, err error) {
	// todo edit your code

	userResp, err = rpc_user.Client.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return &user.LoginResp{}, "", err
	}

	/*session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	err = session.Save()
	if err != nil {
		return "", err
	}*/
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return userResp, redirect, err
}
