package errors

import (
	"github.com/pkg/errors"
)

// 定义别名
var (
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

// 定义错误码
var (
	//HTTP COMMON [0, 600)
	ErrCodeParamInvalid = 400

	//COMMON
	ErrCodeAppNotExist = 600

	//Login [1000, 1200)
	ErrCodeLoginFailed = 1000

	//Account [1200, 1400)
	ErrCodeUsernameInvalid = 1200
	ErrCodeAccountSystemNotExist = 1201
)

// 定义错误
var (
	//ErrInvalidUserName         = New400Response(ErrCodeParamInvalid, "无效的用户名")
	//ErrInvalidPassword         = New400Response(ErrCodeParamInvalid, "无效的密码")
	//ErrBadRequest              = New400Response("请求发生错误")
	//ErrInvalidParent           = New400Response("无效的父级节点")
	//ErrNotAllowDeleteWithChild = New400Response("含有子级，不能删除")
	//ErrNotAllowDelete          = New400Response("资源不允许删除")
	//ErrInvalidUser             = New400Response("无效的用户")
	//ErrUserDisable             = New400Response("用户被禁用，请联系管理员")

	ErrInvalidToken    = NewResponse(10000, 401, "令牌无效")
	ErrNoPerm          = NewResponse(401, 401, "无访问权限")
	ErrNotFound        = NewResponse(404, 404, "资源不存在")
	ErrMethodNotAllow  = NewResponse(405, 405, "方法不被允许")
	ErrTooManyRequests = NewResponse(429, 429, "请求过于频繁")
	ErrInternalServer  = NewResponse(500, 500, "服务器发生错误")
)