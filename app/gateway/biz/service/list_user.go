package service

import (
	"context"

	user2 "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/user"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/user"
	rpc_user "github.com/777continue/gomall/rpc_gen/rpc/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListUserService(Context context.Context, RequestContext *app.RequestContext) *ListUserService {
	return &ListUserService{RequestContext: RequestContext, Context: Context}
}

func (h *ListUserService) Run(req *user2.ListUserReq) (resp *user2.ListUserResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	rpcResp, err := rpc_user.Client.ListUser(h.Context, &user.ListUserReq{})
	if err != nil {
		return nil, err
	}

	users := make([]*user2.User, 0, len(rpcResp.Users))
	for _, u := range rpcResp.Users {
		users = append(users, &user2.User{
			UserId:  u.UserId,
			Email:   u.Email,
			IsAdmin: u.IsAdmin,
		})
	}

	return &user2.ListUserResp{
		Users: users,
	}, nil
}
