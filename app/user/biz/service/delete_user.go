package service

import (
	"context"

	"github.com/777continue/gomall/app/user/biz/dal/mysql"
	casbin "github.com/777continue/gomall/app/user/biz/infra"
	"github.com/777continue/gomall/app/user/biz/model"
	user "github.com/777continue/gomall/rpc_gen/kitex_gen/user"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	// Finish your business logic.
	err = model.DeleteUser(s.ctx, mysql.DB, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	userInfo, err := model.GetByID(s.ctx, mysql.DB, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	// 删除Casbin中的用户角色
	_, err = casbin.EF.DeleteUser(userInfo.Email)
	if err != nil {
		return nil, err
	}

	return &user.DeleteUserResp{}, nil
}
