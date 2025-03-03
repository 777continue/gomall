package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/777continue/gomall/app/user/biz/dal/mysql"
	"github.com/777continue/gomall/app/user/biz/model"
	user "github.com/777continue/gomall/rpc_gen/kitex_gen/user"
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

func generateToken(row *model.User, exp int64) (tokenString string, err error) {
	// 生成JWT令牌
	// 这里使用HS256算法，你可以根据需要修改
	token := jwt.New(jwt.SigningMethodHS256)
	// 设置令牌的声明
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = row.ID
	claims["email"] = row.Email
	// 你可以根据需要设置过期时间
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// 从环境变量中获取JWT签名密钥
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		return "", errors.New("JWT_SECRET_KEY environment variable not set")
	}

	// 生成签名字符串
	tokenString, err = token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
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
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))

	if err != nil {
		return nil, err
	}

	// 初始化Casbin
	e, err := casbin.NewEnforcer("../config/casbin/model.conf", "../config/casbin/policy.csv")
	if err != nil {
		return nil, err
	}

	e.AddRoleForUser(row.Email, "user")

	tokenString, err := generateToken(row, time.Now().Add(time.Hour*24).Unix())
	if err != nil {
		return nil, err
	}

	return &user.LoginResp{
		UserId: int32(row.ID),
		Token:  tokenString,
	}, nil
}
