package mw

import (
	"context"
	"strings"
	"time"

	TokenUtils "github.com/777continue/gomall/app/frontend/biz/token"
	frontendUtils "github.com/777continue/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgrijalva/jwt-go"
)

func Cors() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 允许所有来源的跨域请求
		c.Header("Access-Control-Allow-Origin", "*")
		// 允许的请求方法
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许的请求头，包括 authorization
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 允许携带凭证（如 cookies）
		c.Header("Access-Control-Allow-Credentials", "true")

		// 如果是预检请求（OPTIONS 请求），直接返回 200 状态码
		if string(c.Request.Method()) == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		// 继续处理其他请求
		c.Next(ctx)
	}
}
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		hlog.Info("GlobalAuth")
		tokenString := string(c.GetHeader("authorization"))
		hlog.Infof("tokenString: %v", tokenString)
		if len(tokenString) > 7 && strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
			tokenString = tokenString[7:]
		}
		hlog.Infof("tokenString: %v", tokenString)
		if tokenString != "" {
			userId, err := TokenUtils.ValidateToken(tokenString)
			if err == nil {
				ctx = context.WithValue(ctx, frontendUtils.SessionUserId, userId)
				hlog.Infof("token validate success, userId: %v", userId)
				// 解析token获取过期时间
				token, _ := jwt.Parse(tokenString, nil)
				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					exp := int64(claims["exp"].(float64))
					now := time.Now().Unix()

					// 如果token即将过期（剩余时间小于30分钟），则续期
					if exp-now < 1800 {
						newToken, err := TokenUtils.GenerateToken(userId, time.Now().Add(time.Hour*24).Unix())
						if err == nil {
							c.Header("X-Renewed-Token", newToken)
						}
					}
				}
			} else {
				hlog.Infof("token error: %v", err)
			}
		}
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		userId := ctx.Value(frontendUtils.SessionUserId)
		hlog.Infof("userId: %v", userId)
		if userId == nil {
			c.JSON(401, map[string]interface{}{
				"code":    401,
				"message": "Unauthorized: Please login first",
			})
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
