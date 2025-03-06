package service

import (
	"context"
	"errors"
	"time"

	"github.com/777continue/gomall/app/user/biz/dal/mysql"
	casbin "github.com/777continue/gomall/app/user/biz/infra"
	"github.com/777continue/gomall/app/user/biz/model"
	user "github.com/777continue/gomall/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	row, err := model.GetByEmail(s.ctx, mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}
	//err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))

	if row.PasswordHashed != req.Password {
		return nil, err
	}
	// 检查用户是否是admin
	isAdmin, err := casbin.EF.HasRoleForUser(row.Email, "admin")
	if err != nil {
		return nil, err
	}

	tokenString, err := GenerateToken(row, time.Now().Add(time.Hour*24).Unix())
	if err != nil {
		return nil, err
	}

	return &user.LoginResp{
		Token:   tokenString,
		IsAdmin: isAdmin, // 返回用户是否是admin
	}, nil
}
