package api

import (
	"github.com/LyricTian/gin-admin/v6/internal/app/bll"
	"github.com/LyricTian/gin-admin/v6/internal/app/ginplus"
	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/errors"
	"github.com/LyricTian/gin-admin/v6/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// LoginSet 注入Login
var LoginSet = wire.NewSet(wire.Struct(new(Login), "*"))

// Login 登录管理
type Login struct {
	LoginBll 	bll.ILogin
	AppBll 		bll.IApp
}

//----------------------------------------------------

// Login 用户登录
func (a *Login) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.LoginParam
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	//if !captcha.VerifyString(item.CaptchaID, item.CaptchaCode) {
	//	ginplus.ResError(c, errors.New400Response("无效的验证码"))
	//	return
	//}

	account, err := a.LoginBll.Verify(ctx, item.AppKey, item.Username, item.Password)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	accountKey := account.AccountKey
	// 将账号标识放入上下文
	ginplus.SetAccountKey(c, accountKey)

	appResults, err := a.AppBll.Query(c, schema.AppQueryParam{AccountType: account.AccountType})
	if err != nil {
		ginplus.ResError(c, errors.WithStack(err))
		return
	}

	var info struct {
		AppKeys []string
	}
	for _, v := range appResults.Data {
		info.AppKeys = append(info.AppKeys, v.AppKey)
	}

	ctx = logger.NewAccountKeyContext(ctx, accountKey)
	tokenInfo, err := a.LoginBll.GenerateToken(ctx, accountKey, info)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	logger.StartSpan(ctx, logger.SetSpanTitle("用户登录"), logger.SetSpanFuncName("Login")).Infof("登入系统")
	ginplus.ResSuccess(c, tokenInfo)
}

// Authenticate 鉴权
func (a *Login) Authenticate(c *gin.Context) {
	ctx := c.Request.Context()

	accountKey := ginplus.GetAccountKey(c)
	appKey := c.GetHeader("RJ-AppKey")

	//TODO 奇怪..........
	var item schema.AuthenticateParam
	item.Feature = c.Query("feature")
	item.ResourceType = c.DefaultQuery("resourceType", "api")
	item.Method = c.DefaultQuery("method", "")

	r, err := a.LoginBll.Authenticate(ctx, appKey, accountKey, item.ResourceType, item.Feature, item.Method)
	if err != nil {
		ginplus.ResError(c, errors.WithStack(err))
		return
	}

	var grant int64 = 0
	if r {
		grant = 1
	}

	ginplus.ResSuccess(c, schema.AuthenticateInfo{Grant: grant})
}

//----------------------------------------------------

//// GetCaptcha 获取验证码信息
//func (a *Login) GetCaptcha(c *gin.Context) {
//	ctx := c.Request.Context()
//	item, err := a.LoginBll.GetCaptcha(ctx, config.C.Captcha.Length)
//	if err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//	ginplus.ResSuccess(c, item)
//}
//
//// ResCaptcha 响应图形验证码
//func (a *Login) ResCaptcha(c *gin.Context) {
//	ctx := c.Request.Context()
//	captchaID := c.Query("id")
//	if captchaID == "" {
//		ginplus.ResError(c, errors.New400Response("请提供验证码ID"))
//		return
//	}
//
//	if c.Query("reload") != "" {
//		if !captcha.Reload(captchaID) {
//			ginplus.ResError(c, errors.New400Response("未找到验证码ID"))
//			return
//		}
//	}
//
//	cfg := config.C.Captcha
//	err := a.LoginBll.ResCaptcha(ctx, c.Writer, captchaID, cfg.Width, cfg.Height)
//	if err != nil {
//		ginplus.ResError(c, err)
//	}
//}
//
//// Logout 用户登出
//func (a *Login) Logout(c *gin.Context) {
//	ctx := c.Request.Context()
//	// 检查用户是否处于登录状态，如果是则执行销毁
//	userID := ginplus.GetUserID(c)
//	if userID != "" {
//		err := a.LoginBll.DestroyToken(ctx, ginplus.GetToken(c))
//		if err != nil {
//			logger.Errorf(ctx, err.Error())
//		}
//		logger.StartSpan(ctx, logger.SetSpanTitle("用户登出"), logger.SetSpanFuncName("Logout")).Infof("登出系统")
//	}
//	ginplus.ResOK(c)
//}
//
//// RefreshToken 刷新令牌
//func (a *Login) RefreshToken(c *gin.Context) {
//	ctx := c.Request.Context()
//	tokenInfo, err := a.LoginBll.GenerateToken(ctx, ginplus.GetUserID(c))
//	if err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//	ginplus.ResSuccess(c, tokenInfo)
//}
//
//// GetUserInfo 获取当前用户信息
//func (a *Login) GetUserInfo(c *gin.Context) {
//	ctx := c.Request.Context()
//	info, err := a.LoginBll.GetLoginInfo(ctx, ginplus.GetUserID(c))
//	if err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//	ginplus.ResSuccess(c, info)
//}
//
//// QueryUserMenuTree 查询当前用户菜单树
//func (a *Login) QueryUserMenuTree(c *gin.Context) {
//	ctx := c.Request.Context()
//	menus, err := a.LoginBll.QueryUserMenuTree(ctx, ginplus.GetUserID(c))
//	if err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//	ginplus.ResList(c, menus)
//}
//
//// UpdatePassword 更新个人密码
//func (a *Login) UpdatePassword(c *gin.Context) {
//	ctx := c.Request.Context()
//	var item schema.UpdatePasswordParam
//	if err := ginplus.ParseJSON(c, &item); err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//
//	err := a.LoginBll.UpdatePassword(ctx, ginplus.GetUserID(c), item)
//	if err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//	ginplus.ResOK(c)
//}
