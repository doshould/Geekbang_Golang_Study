package api

import (
	"github.com/LyricTian/gin-admin/v6/internal/app/bll"
	"github.com/LyricTian/gin-admin/v6/internal/app/ginplus"
	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// UserSet 注入User
var AccountSet = wire.NewSet(wire.Struct(new(Account), "*"))

// Account 账号管理
type Account struct {
	AccountBll bll.IAccount
}

// Create 创建数据
func (a *Account) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.AccountCreateParam
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	accountKey, err := a.AccountBll.Create(ctx, item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, accountKey)
}

//// Query 查询数据
//func (a *Account) Query(c *gin.Context) {
//	ctx := c.Request.Context()
//	var params schema.AccountQueryParam
//	if err := ginplus.ParseQuery(c, &params); err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//	if v := c.Query("roleIDs"); v != "" {
//		params.RoleIDs = strings.Split(v, ",")
//	}
//
//	params.Pagination = true
//	result, err := a.UserBll.QueryShow(ctx, params)
//	if err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//	ginplus.ResPage(c, result.Data, result.PageResult)
//}
//
//// Get 查询指定数据
//func (a *Account) Get(c *gin.Context) {
//	ctx := c.Request.Context()
//	item, err := a.UserBll.Get(ctx, c.Param("id"))
//	if err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//	ginplus.ResSuccess(c, item.CleanSecure())
//}

//// Update 更新数据
//func (a *Account) Update(c *gin.Context) {
//	ctx := c.Request.Context()
//	var item schema.Account
//	if err := ginplus.ParseJSON(c, &item); err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//
//	err := a.UserBll.Update(ctx, c.Param("id"), item)
//	if err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//	ginplus.ResOK(c)
//}
//
//// Delete 删除数据
//func (a *Account) Delete(c *gin.Context) {
//	ctx := c.Request.Context()
//	err := a.UserBll.Delete(ctx, c.Param("id"))
//	if err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//	ginplus.ResOK(c)
//}
//
//// Enable 启用数据
//func (a *Account) Enable(c *gin.Context) {
//	ctx := c.Request.Context()
//	err := a.UserBll.UpdateStatus(ctx, c.Param("id"), 1)
//	if err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//	ginplus.ResOK(c)
//}
//
//// Disable 禁用数据
//func (a *Account) Disable(c *gin.Context) {
//	ctx := c.Request.Context()
//	err := a.UserBll.UpdateStatus(ctx, c.Param("id"), 2)
//	if err != nil {
//		ginplus.ResError(c, err)
//		return
//	}
//	ginplus.ResOK(c)
//}
