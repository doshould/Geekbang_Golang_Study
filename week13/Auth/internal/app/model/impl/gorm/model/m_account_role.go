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

// AccountRole 账号角色存储
type AccountRole struct {
	DB *gorm.DB
}

var _ model.IAccountRole = (*AccountRole)(nil)

// AccountRoleSet 注入AccountRoleSet
var AccountRoleSet = wire.NewSet(wire.Struct(new(AccountRole), "*"), wire.Bind(new(model.IAccountRole), new(*AccountRole)))

func (a *AccountRole) getQueryOption(opts ...schema.AccountRoleQueryOptions) schema.AccountRoleQueryOptions {
	var opt schema.AccountRoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}
//Query(ctx context.Context, params schema.AccountRoleQueryParam, opts ...schema.AccountRoleQueryOptions) (*schema.AccountRoleQueryResult, error)
// Query 查询数据
func (a *AccountRole) Query(ctx context.Context, params schema.AccountRoleQueryParam, opts ...schema.AccountRoleQueryOptions) (*schema.AccountRoleQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetAccountRoleDB(ctx, a.DB)
	if v := params.AccountKey; v != "" {
		db = db.Where("accountKey = ?", v)
	}
	if v := params.AccountKeys; len(v) > 0 {
		db = db.Where("accountKey IN (?)", v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.AccountRoles
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.AccountRoleQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaAccountRoles(),
	}

	return qr, nil
}

//// Get 查询指定数据
//func (a *UserRole) Get(ctx context.Context, id string, opts ...schema.UserRoleQueryOptions) (*schema.UserRole, error) {
//	db := entity.GetUserRoleDB(ctx, a.DB).Where("id=?", id)
//	var item entity.UserRole
//	ok, err := FindOne(ctx, db, &item)
//	if err != nil {
//		return nil, errors.WithStack(err)
//	} else if !ok {
//		return nil, nil
//	}
//
//	return item.ToSchemaUserRole(), nil
//}
//
//// Create 创建数据
//func (a *UserRole) Create(ctx context.Context, item schema.UserRole) error {
//	eitem := entity.SchemaUserRole(item).ToUserRole()
//	result := entity.GetUserRoleDB(ctx, a.DB).Create(eitem)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// Update 更新数据
//func (a *UserRole) Update(ctx context.Context, id string, item schema.UserRole) error {
//	eitem := entity.SchemaUserRole(item).ToUserRole()
//	result := entity.GetUserRoleDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// Delete 删除数据
//func (a *UserRole) Delete(ctx context.Context, id string) error {
//	result := entity.GetUserRoleDB(ctx, a.DB).Where("id=?", id).Delete(entity.UserRole{})
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// DeleteByUserID 根据用户ID删除数据
//func (a *UserRole) DeleteByUserID(ctx context.Context, userID string) error {
//	result := entity.GetUserRoleDB(ctx, a.DB).Where("user_id=?", userID).Delete(entity.UserRole{})
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
