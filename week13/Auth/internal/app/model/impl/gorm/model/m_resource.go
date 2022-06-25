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

var _ model.IResource = (*Resource)(nil)

// ResourceSet 注入Resource
var ResourceSet = wire.NewSet(wire.Struct(new(Resource), "*"), wire.Bind(new(model.IResource), new(*Resource)))

// Resource 资源存储
type Resource struct {
	DB *gorm.DB
}

func (a *Resource) getQueryOption(opts ...schema.ResourceQueryOptions) schema.ResourceQueryOptions {
	var opt schema.ResourceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Resource) Query(ctx context.Context, params schema.ResourceQueryParam, opts ...schema.ResourceQueryOptions) (*schema.ResourceQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetResourceDB(ctx, a.DB)
	if v := params.IDs; len(v) > 0 {
		db = db.Where("id IN (?)", v)
	}
	if v := params.Name; v != "" {
		db = db.Where("name=?", v)
	}
	if v := params.Feature; v != "" {
		db = db.Where("feature=?", v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.Resources
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.ResourceQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaResources(),
	}

	return qr, nil
}

//// Get 查询指定数据
//func (a *Menu) Get(ctx context.Context, id string, opts ...schema.MenuQueryOptions) (*schema.Menu, error) {
//	var item entity.Menu
//	ok, err := FindOne(ctx, entity.GetMenuDB(ctx, a.DB).Where("id=?", id), &item)
//	if err != nil {
//		return nil, errors.WithStack(err)
//	} else if !ok {
//		return nil, nil
//	}
//
//	return item.ToSchemaMenu(), nil
//}
//
//// Create 创建数据
//func (a *Menu) Create(ctx context.Context, item schema.Menu) error {
//	eitem := entity.SchemaMenu(item).ToMenu()
//	result := entity.GetMenuDB(ctx, a.DB).Create(eitem)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// Update 更新数据
//func (a *Menu) Update(ctx context.Context, id string, item schema.Menu) error {
//	eitem := entity.SchemaMenu(item).ToMenu()
//	result := entity.GetMenuDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// UpdateParentPath 更新父级路径
//func (a *Menu) UpdateParentPath(ctx context.Context, id, parentPath string) error {
//	result := entity.GetMenuDB(ctx, a.DB).Where("id=?", id).Update("parent_path", parentPath)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// Delete 删除数据
//func (a *Menu) Delete(ctx context.Context, id string) error {
//	result := entity.GetMenuDB(ctx, a.DB).Where("id=?", id).Delete(entity.Menu{})
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// UpdateStatus 更新状态
//func (a *Menu) UpdateStatus(ctx context.Context, id string, status int) error {
//	result := entity.GetMenuDB(ctx, a.DB).Where("id=?", id).Update("status", status)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
