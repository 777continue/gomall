package mw

import (
	"context"
	"errors"
	"fmt"
	"os"

	frontendUtils "github.com/777continue/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgrijalva/jwt-go"
)

// 校验 JWT 的函数
func validateToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 从环境变量获取JWT密钥
		jwtSecret := os.Getenv("JWT_SECRET_KEY")
		if jwtSecret == "" {
			return nil, errors.New("JWT_SECRET_KEY environment variable not set")
		}
		return []byte(jwtSecret), nil
	})

	// 处理解析错误
	if err != nil {
		return 0, err
	}

	// 验证token是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 获取用户ID
		if userId, ok := claims["userId"].(float64); ok {
			return int(userId), nil
		}
		return 0, errors.New("invalid user ID in token")
	}

	return 0, errors.New("invalid token")
}
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		tokenString := string(c.GetHeader("Authorization"))
		if tokenString != "" {
			if userId, err := validateToken(tokenString); err == nil {
				ctx = context.WithValue(ctx, frontendUtils.SessionUserId, userId)
			}
		}
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		userId := ctx.Value(frontendUtils.SessionUserId)
		if userId == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
