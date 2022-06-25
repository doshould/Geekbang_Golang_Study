package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// UserSet 注入User
var AccountSet = wire.NewSet(wire.Struct(new(Account), "*"))

// Account 账号管理
type Account struct {
}

// Create 创建账号
// @Tags 账号管理
// @Summary 创建账号
// @Security ApiKeyAuth
// @Param body body schema.AccountCreateParam true "创建账号"
// @Success 200 {object} schema.NilResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:400,message:无效的请求参数。（账号标识不符合格式/账号名不符合格式）}}"
// @Failure 403 {object} schema.ErrorResult "{error:{code:1200,message:账号标识已存在/账号名已存在}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:500,message:服务器错误}}"
// @Router /api/pub/users [post]
func (a *Account) Create(c *gin.Context) {
}