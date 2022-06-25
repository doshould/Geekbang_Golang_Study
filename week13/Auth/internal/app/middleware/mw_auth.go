package middleware

import (
	"github.com/LyricTian/gin-admin/v6/internal/app/config"
	"github.com/LyricTian/gin-admin/v6/internal/app/ginplus"
	"github.com/LyricTian/gin-admin/v6/internal/app/icontext"
	"github.com/LyricTian/gin-admin/v6/pkg/auth"
	"github.com/LyricTian/gin-admin/v6/pkg/errors"
	"github.com/LyricTian/gin-admin/v6/pkg/logger"
	"github.com/gin-gonic/gin"
)

func wrapAccountAuthContext(c *gin.Context, accountKey string) {
	ginplus.SetAccountKey(c, accountKey)
	ctx := icontext.NewAccountKey(c.Request.Context(), accountKey)
	ctx = logger.NewAccountKeyContext(ctx, accountKey)
	c.Request = c.Request.WithContext(ctx)
}

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(a auth.Auther, skippers ...SkipperFunc) gin.HandlerFunc {
	if !config.C.JWTAuth.Enable {
		return func(c *gin.Context) {
			wrapAccountAuthContext(c, config.C.Root.Username)
			c.Next()
		}
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID, info, err := a.ParseUserID(c.Request.Context(), ginplus.GetToken(c))
		if err != nil {
			if err == auth.ErrInvalidToken {
				if config.C.IsDebugMode() {
					wrapAccountAuthContext(c, config.C.Root.Username)
					c.Next()
					return
				}
				ginplus.ResError(c, errors.ErrInvalidToken)
				return
			}

			//ginplus.ResError(c, errors.WithStack(err))
			ginplus.ResError(c, errors.ErrInvalidToken)
			return
		}

		//或者可利用json序列化再反序列化取出数据
		_, ok := info["AppKeys"]
		if !ok {
			ginplus.ResError(c, errors.ErrInvalidToken)
			return
		}

		appKeys, ok := info["AppKeys"].([]interface{})
		if !ok {
			ginplus.ResError(c, errors.ErrInvalidToken)
			return
		}

		appKey := c.GetHeader("RJ-AppKey")
		support := false
		for _, v := range appKeys {
			if rv, ok := v.(string); ok && rv == appKey {
				support = true
				break
			}
		}
		if !support {
			ginplus.ResError(c, errors.ErrInvalidToken)
			return
		}

		wrapAccountAuthContext(c, userID)
		c.Next()
	}
}
