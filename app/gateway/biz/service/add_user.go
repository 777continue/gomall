package service

import (
	"context"

	user2 "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/user"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/user"
	rpc_user "github.com/777continue/gomall/rpc_gen/rpc/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddUserService(Context context.Context, RequestContext *app.RequestContext) *AddUserService {
	return &AddUserService{RequestContext: RequestContext, Context: Context}
}

func (h *AddUserService) Run(req *user2.AddUserReq) (resp *user2.AddUserResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc_user.Client.AddUser(h.Context, &user.AddUserReq{
		Email:    req.Email,
		Password: req.Password,
		IsAdmin:  req.IsAdmin,
	})
	if err != nil {
		return nil, err
	}
	return &user2.AddUserResp{}, nil
}
