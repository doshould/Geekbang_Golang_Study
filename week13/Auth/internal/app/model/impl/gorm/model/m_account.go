package model

import (
	"context"
	"github.com/LyricTian/gin-admin/v6/internal/app/model"
	"github.com/LyricTian/gin-admin/v6/internal/app/model/impl/gorm/entity"
	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/errors"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var _ model.IAccount = (*Account)(nil)

// AccountSet 注入Account
var AccountSet = wire.NewSet(wire.Struct(new(Account), "*"), wire.Bind(new(model.IAccount), new(*Account)))

// User 用户存储
type Account struct {
	DB *gorm.DB
}

func (a *Account) getQueryOption(opts ...schema.AccountQueryOptions) schema.AccountQueryOptions {
	var opt schema.AccountQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Account) Query(ctx context.Context, params schema.AccountQueryParam, opts ...schema.AccountQueryOptions) (*schema.AccountQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetAccountDB(ctx, a.DB)
	if v := params.AppKey; v != "" {
		db = db.Joins("left join App on App.AccountType = Account.AccountType")
		db = db.Where("App.appKey = ?", v)
	}

	if v := params.Username; v != "" {
		db = db.Where("username = ?", v)
	}
	if v := params.Password; v != "" {
		db = db.Where("password = ?", v)
	}
	if v := params.AccountType; v != "" {
		db = db.Where("Account.accountType = ?", v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.Accounts
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.AccountQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaAccounts(),
	}
	return qr, nil
}

// Create 创建数据
func (a *Account) Create(ctx context.Context, item schema.Account) error {
	sAccount := entity.SchemaAccount(item)
	result := entity.GetAccountDB(ctx, a.DB).Create(sAccount.ToNewAccount())
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//// Get 查询指定数据
//func (a *Account) Get(ctx context.Context, id string, opts ...schema.UserQueryOptions) (*schema.Account, error) {
//	var item entity.User
//	ok, err := FindOne(ctx, entity.GetUserDB(ctx, a.DB).Where("id=?", id), &item)
//	if err != nil {
//		return nil, errors.WithStack(err)
//	} else if !ok {
//		return nil, nil
//	}
//
//	return item.ToSchemaUser(), nil
//}
//
//
//// Update 更新数据
//func (a *Account) Update(ctx context.Context, id string, item schema.Account) error {
//	eitem := entity.SchemaUser(item).ToUser()
//	result := entity.GetUserDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// Delete 删除数据
//func (a *Account) Delete(ctx context.Context, id string) error {
//	result := entity.GetUserDB(ctx, a.DB).Where("id=?", id).Delete(entity.User{})
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// UpdateStatus 更新状态
//func (a *Account) UpdateStatus(ctx context.Context, id string, status int) error {
//	result := entity.GetUserDB(ctx, a.DB).Where("id=?", id).Update("status", status)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// UpdatePassword 更新密码
//func (a *Account) UpdatePassword(ctx context.Context, id, password string) error {
//	result := entity.GetUserDB(ctx, a.DB).Where("id=?", id).Update("password", password)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
