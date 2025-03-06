package service

import (
	"context"
	"errors"

	"github.com/777continue/gomall/app/user/biz/dal/mysql"
	casbin "github.com/777continue/gomall/app/user/biz/infra"
	"github.com/777continue/gomall/app/user/biz/model"
	user "github.com/777continue/gomall/rpc_gen/kitex_gen/user"
)

type AddUserService struct {
	ctx context.Context
} // NewAddUserService new AddUserService
func NewAddUserService(ctx context.Context) *AddUserService {
	return &AddUserService{ctx: ctx}
}

// Run create note info
func (s *AddUserService) Run(req *user.AddUserReq) (resp *user.AddUserResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	err = model.AddUser(s.ctx, mysql.DB, &model.User{
		Email:          req.Email,
		PasswordHashed: req.Password,
	})
	if err != nil {
		return nil, err
	}
	// 根据isAdmin添加角色
	if req.IsAdmin {
		_, err = casbin.EF.AddRoleForUser(req.Email, "admin")
	} else {
		_, err = casbin.EF.AddRoleForUser(req.Email, "user")
	}
	if err != nil {
		return nil, err
	}

	return &user.AddUserResp{}, err
}
