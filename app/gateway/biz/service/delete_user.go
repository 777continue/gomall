package service

import (
	"context"

	user2 "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/user"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/user"
	rpc_user "github.com/777continue/gomall/rpc_gen/rpc/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteUserService(Context context.Context, RequestContext *app.RequestContext) *DeleteUserService {
	return &DeleteUserService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteUserService) Run(req *user2.DeleteUserReq) (resp *user2.DeleteUserResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc_user.Client.DeleteUser(h.Context, &user.DeleteUserReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &user2.DeleteUserResp{}, nil
}
