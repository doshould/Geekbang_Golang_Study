package bll

import (
	"context"
	"fmt"
	"regexp"

	"github.com/LyricTian/gin-admin/v6/internal/app/bll"
	"github.com/LyricTian/gin-admin/v6/internal/app/iutil"
	"github.com/LyricTian/gin-admin/v6/internal/app/model"
	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/errors"
	"github.com/LyricTian/gin-admin/v6/pkg/util"
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
)

// Account 用户管理
type Account struct {
	Enforcer      		*casbin.SyncedEnforcer
	TransModel    		model.ITrans
	AccountModel     	model.IAccount
	AccountSystemModel  model.IAccountSystem
	AccountRoleModel 	model.IAccountRole
	RoleModel     		model.IRole
}

var _ bll.IAccount = (*Account)(nil)

// AccountSet 注入Account
var AccountSet = wire.NewSet(wire.Struct(new(Account), "*"), wire.Bind(new(bll.IAccount), new(*Account)))

// Query 查询数据
func (a *Account) Query(ctx context.Context, params schema.AccountQueryParam, opts ...schema.AccountQueryOptions) (*schema.AccountQueryResult, error) {
	return a.AccountModel.Query(ctx, params, opts...)
}

// Create 创建数据
func (a *Account) Create(ctx context.Context, item schema.AccountCreateParam) (string, error) {
	var err error

	isExist, err := a.isAccountTypeExist(ctx, item.AccountType)
	if err != nil {
		return "", err
	} else if isExist == false {
		return "", errors.New403Response(errors.ErrCodeAccountSystemNotExist, "账号类型不存在")
	}

	err = a.checkUsername(ctx, item.AccountType, item.Username)
	if err != nil {
		return "", err
	}

	err = a.checkPassword(ctx, item.Password)
	if err != nil {
		return "", err
	}

	var account schema.Account
	account.AccountKey = iutil.NewID()
	account.AccountType = item.AccountType
	account.Username = item.Username
	account.Password = util.MD5HashString(item.Password)

	//err = ExecTrans(ctx, a.TransModel, func(ctx context.Context) error {
	//	return a.AccountModel.Create(ctx, item)
	//})
	err = a.AccountModel.Create(ctx, account)
	if err != nil {
		return "", err
	}

	return account.AccountKey, nil
}

func (a *Account) isAccountTypeExist(ctx context.Context, accountType string) (exist bool, err error) {
	result, err := a.AccountSystemModel.Query(ctx, schema.AccountSystemQueryParam{
		PaginationParam: 	schema.PaginationParam{OnlyCount: true},
		AccountType:		accountType,
	})
	if err != nil {
		return false, err
	} else if result.PageResult.Total > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (a *Account) checkUsername(ctx context.Context, accountType, username string) error {
	match, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username)
	if match == false {
		return errors.New400Response(errors.ErrCodeParamInvalid, "无效的请求参数（账号名不符合格式，应由字母或数字组成，长度为4~16位）")
	}

	if username == schema.GetRootAccount().Username {
		return errors.New403Response(errors.ErrCodeUsernameInvalid, fmt.Sprintf("%s为内置系统账号，不能使用该账号名", username))
	}

	result, err := a.AccountModel.Query(ctx, schema.AccountQueryParam{
		PaginationParam: 	schema.PaginationParam{OnlyCount: true},
		Username:			username,
		AccountType:		accountType,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New403Response(errors.ErrCodeUsernameInvalid, "账号名已经存在")
	}

	return nil
}

func (a *Account) checkPassword(ctx context.Context, password string) error {
	match , _ := regexp.MatchString("^[a-zA-Z0-9()`~!@#$%^&*\\-_+=|{}\\[\\]:;'<>,.?\\/]{8,16}$", password)
	if match == false {
		return errors.New400Response(errors.ErrCodeParamInvalid, "无效的请求参数（密码不符合格式，应由字母、数字或特殊字符组成，长度为8~16位）")
	}

	return nil
}