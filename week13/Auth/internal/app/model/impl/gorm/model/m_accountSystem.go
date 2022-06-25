package model

import (
	"context"
	"github.com/LyricTian/gin-admin/v6/internal/app/model/impl/gorm/entity"

	"github.com/LyricTian/gin-admin/v6/internal/app/model"
	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/errors"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var _ model.IAccountSystem = (*AccountSystem)(nil)

// AccountSystemSet 注入AccountSystem
var AccountSystemSet = wire.NewSet(wire.Struct(new(AccountSystem), "*"), wire.Bind(new(model.IAccountSystem), new(*AccountSystem)))

// AccountSystem 账号体系存储
type AccountSystem struct {
	DB *gorm.DB
}

func (a *AccountSystem) getQueryOption(opts ...schema.QueryOptions) schema.QueryOptions {
	var opt schema.QueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *AccountSystem) Query(ctx context.Context, params schema.AccountSystemQueryParam, opts ...schema.QueryOptions) (*schema.AccountSystemQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetAccountSystemDB(ctx, a.DB)
	if v := params.AccountType; v != "" {
		db = db.Where("accountType = ?", v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByASC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.AccountSystems
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.AccountSystemQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaAccountSystems(),
	}

	return qr, nil
}

//// Get 查询指定数据
//func (a *MenuAction) Get(ctx context.Context, id string, opts ...schema.MenuActionQueryOptions) (*schema.MenuAction, error) {
//	db := entity.GetMenuActionDB(ctx, a.DB).Where("id=?", id)
//	var item entity.MenuAction
//	ok, err := FindOne(ctx, db, &item)
//	if err != nil {
//		return nil, errors.WithStack(err)
//	} else if !ok {
//		return nil, nil
//	}
//
//	return item.ToSchemaMenuAction(), nil
//}
//
//// Create 创建数据
//func (a *MenuAction) Create(ctx context.Context, item schema.MenuAction) error {
//	eitem := entity.SchemaMenuAction(item).ToMenuAction()
//	result := entity.GetMenuActionDB(ctx, a.DB).Create(eitem)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// Update 更新数据
//func (a *MenuAction) Update(ctx context.Context, id string, item schema.MenuAction) error {
//	eitem := entity.SchemaMenuAction(item).ToMenuAction()
//	result := entity.GetMenuActionDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// Delete 删除数据
//func (a *MenuAction) Delete(ctx context.Context, id string) error {
//	result := entity.GetMenuActionDB(ctx, a.DB).Where("id=?", id).Delete(entity.MenuAction{})
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// DeleteByMenuID 根据菜单ID删除数据
//func (a *MenuAction) DeleteByMenuID(ctx context.Context, menuID string) error {
//	result := entity.GetMenuActionDB(ctx, a.DB).Where("menu_id=?", menuID).Delete(entity.MenuAction{})
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
