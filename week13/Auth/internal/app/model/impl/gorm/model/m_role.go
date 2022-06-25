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

var _ model.IRole = (*Role)(nil)

// RoleSet 注入Role
var RoleSet = wire.NewSet(wire.Struct(new(Role), "*"), wire.Bind(new(model.IRole), new(*Role)))

// Role 角色存储
type Role struct {
	DB *gorm.DB
}

func (a *Role) getQueryOption(opts ...schema.RoleQueryOptions) schema.RoleQueryOptions {
	var opt schema.RoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Role) Query(ctx context.Context, params schema.RoleQueryParam, opts ...schema.RoleQueryOptions) (*schema.RoleQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetRoleDB(ctx, a.DB)
	if v := params.AppKey; v != "" {
		db = db.Where("appKey = ?", v)
	}
	if v := params.IDs; len(v) > 0 {
		db = db.Where("id IN (?)", v)
	}
	if v := params.Name; v != "" {
		db = db.Where("name = ?", v)
	}
	if v := params.AccountKey; v != "" {
		subQuery := entity.GetAccountRoleDB(ctx, a.DB).
			Where("accountKey = ?", v).
			Select("roleID").SubQuery()
		db = db.Where("id IN ?", subQuery)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.Roles
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.RoleQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaRoles(),
	}

	return qr, nil
}

//// Get 查询指定数据
//func (a *Role) Get(ctx context.Context, id string, opts ...schema.RoleQueryOptions) (*schema.Role, error) {
//	var role entity.Role
//	ok, err := FindOne(ctx, entity.GetRoleDB(ctx, a.DB).Where("id=?", id), &role)
//	if err != nil {
//		return nil, errors.WithStack(err)
//	} else if !ok {
//		return nil, nil
//	}
//
//	return role.ToSchemaRole(), nil
//}
//
//// Create 创建数据
//func (a *Role) Create(ctx context.Context, item schema.Role) error {
//	eitem := entity.SchemaRole(item).ToRole()
//	result := entity.GetRoleDB(ctx, a.DB).Create(eitem)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// Update 更新数据
//func (a *Role) Update(ctx context.Context, id string, item schema.Role) error {
//	eitem := entity.SchemaRole(item).ToRole()
//	result := entity.GetRoleDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// Delete 删除数据
//func (a *Role) Delete(ctx context.Context, id string) error {
//	result := entity.GetRoleDB(ctx, a.DB).Where("id=?", id).Delete(entity.Role{})
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// UpdateStatus 更新状态
//func (a *Role) UpdateStatus(ctx context.Context, id string, status int) error {
//	result := entity.GetRoleDB(ctx, a.DB).Where("id=?", id).Update("status", status)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
