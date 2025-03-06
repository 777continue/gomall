package service

import (
	"errors"
	"os"
	"time"

	"github.com/777continue/gomall/app/user/biz/model"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(row *model.User, exp int64) (tokenString string, err error) {
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
