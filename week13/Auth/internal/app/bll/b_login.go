package bll

import (
	"context"
	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
)

// ILogin 登录业务逻辑接口
type ILogin interface {
	// 登录验证
	Verify(ctx context.Context, appKey, username, password string) (*schema.Account, error)
	// 生成令牌
	GenerateToken(ctx context.Context, accountKey string, info interface{}) (*schema.LoginTokenInfo, error)
	// 鉴权
	Authenticate(ctx context.Context, appKey, accountKey, resourceType, feature, method string) (bool, error)

	//// 获取图形验证码信息
	//GetCaptcha(ctx context.Context, length int) (*schema.LoginCaptcha, error)
	//// 生成并响应图形验证码
	//ResCaptcha(ctx context.Context, w http.ResponseWriter, captchaID string, width, height int) error
	// 销毁令牌
	//DestroyToken(ctx context.Context, tokenString string) error
	//// 获取用户登录信息
	//GetLoginInfo(ctx context.Context, userID string) (*schema.UserLoginInfo, error)
	//// 查询用户的权限菜单树
	//QueryUserMenuTree(ctx context.Context, userID string) (schema.MenuTrees, error)
}
