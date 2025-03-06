package service

import (
	"context"

	"github.com/777continue/gomall/app/user/biz/dal/mysql"
	"github.com/777continue/gomall/app/user/biz/model"
	user "github.com/777continue/gomall/rpc_gen/kitex_gen/user"
)

type ListUserService struct {
	ctx context.Context
} // NewListUserService new ListUserService
func NewListUserService(ctx context.Context) *ListUserService {
	return &ListUserService{ctx: ctx}
}

// Run create note info
func (s *ListUserService) Run(req *user.ListUserReq) (resp *user.ListUserResp, err error) {
	// Finish your business logic.
	users, err := model.GetAll(s.ctx, mysql.DB)
	if err != nil {
		return nil, err
	}
	ListUsers := make([]*user.User, 0, len(users))
	for _, u := range users {
		ListUsers = append(ListUsers, &user.User{
			UserId: int32(u.ID),
			Email:  u.Email,
		})
	}
	return &user.ListUserResp{
		Users: ListUsers,
	}, nil
}
