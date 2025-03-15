package service

import (
	"context"
	"errors"
	"time"

	"github.com/777continue/gomall/app/frontend/biz/token"
	"github.com/777continue/gomall/app/user/biz/dal/mysql"
	casbin "github.com/777continue/gomall/app/user/biz/infra"
	"github.com/777continue/gomall/app/user/biz/model"
	user "github.com/777continue/gomall/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	//klog.Errorf("req.Password: %s, req.PasswordConfirm : %s", req.Password, req.PasswordConfirm)
	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("email or password is empty")
	}

	if req.Password != req.PasswordConfirm {
		return nil, errors.New("the passwords entered twice do not match")
	}
	// password stored in hash by bcrypt
	//passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(req.Password),
	}
	err = model.Create(s.ctx, mysql.DB, newUser)
	if err != nil {
		return nil, err
	}

	casbin.EF.AddRoleForUser(newUser.Email, "user")

	tokenString, err := token.GenerateToken(uint32(newUser.ID), time.Now().Add(time.Hour*24).Unix())
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{
		UserId: int32(newUser.ID),
		Token:  tokenString,
	}, nil
}
