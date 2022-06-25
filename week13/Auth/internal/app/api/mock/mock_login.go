package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// LoginSet 注入Login
var LoginSet = wire.NewSet(wire.Struct(new(Login), "*"))

// Login 登录管理
type Login struct {
}

// Login 用户登录
// @Tags 登录管理
// @Summary 用户登录
// @Param body body schema.LoginParam true "请求参数"
// @Success 200 {object} schema.LoginTokenInfo
// @Failure 403 {object} schema.ErrorResult "{error:{code:600, message:该应用尚未接入}}<br>{error:{code:1000, message:登录失败，账号或密码错误}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:500, message:服务器错误}}"
// @Router /api/pub/login [post]
func (a *Login) Login(c *gin.Context) {
}

// Authenticate 调用鉴权
// @Tags 登录管理
// @Summary 调用鉴权
// @Param query query schema.AuthenticateParam true "请求参数"
// @Success 200 {object} schema.AuthenticateInfo
// @Failure 401 {object} schema.ErrorResult "{error:{code:401, message:登录信息无效}}"
// @Failure 403 {object} schema.ErrorResult "{error:{code:600, message:该应用尚未接入}}
// @Failure 500 {object} schema.ErrorResult "{error:{code:500,message:服务器错误}}"
// @Router /api/external/login/authenticate [get]
func (a *Login) Authenticate(c *gin.Context) {
}