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

var _ model.IApp = (*App)(nil)

// AppSet 注入App
var AppSet = wire.NewSet(wire.Struct(new(App), "*"), wire.Bind(new(model.IApp), new(*App)))

// App 应用存储
type App struct {
	DB *gorm.DB
}

func (a *App) getQueryOption(opts ...schema.AppQueryOptions) schema.AppQueryOptions {
	var opt schema.AppQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *App) Query(ctx context.Context, params schema.AppQueryParam, opts ...schema.AppQueryOptions) (*schema.AppQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetAppDB(ctx, a.DB)
	if v := params.AppKey; v != "" {
		db = db.Where("appKey = ?", v)
	}
	if v := params.AccountType; v != "" {
		db = db.Where("accountType = ?", v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.Apps
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.AppQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaApps(),
	}

	return qr, nil
}

//// Get 查询指定数据
//func (a *App) Get(ctx context.Context, id string, opts ...schema.DemoQueryOptions) (*schema.Demo, error) {
//	db := entity.GetDemoDB(ctx, a.DB).Where("id=?", id)
//	var item entity.Demo
//	ok, err := FindOne(ctx, db, &item)
//	if err != nil {
//		return nil, errors.WithStack(err)
//	} else if !ok {
//		return nil, nil
//	}
//
//	return item.ToSchemaDemo(), nil
//}
//
//// Create 创建数据
//func (a *App) Create(ctx context.Context, item schema.Demo) error {
//	eitem := entity.SchemaDemo(item).ToDemo()
//	result := entity.GetDemoDB(ctx, a.DB).Create(eitem)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// Update 更新数据
//func (a *App) Update(ctx context.Context, id string, item schema.Demo) error {
//	eitem := entity.SchemaDemo(item).ToDemo()
//	result := entity.GetDemoDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// Delete 删除数据
//func (a *App) Delete(ctx context.Context, id string) error {
//	result := entity.GetDemoDB(ctx, a.DB).Where("id=?", id).Delete(entity.Demo{})
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//// UpdateStatus 更新状态
//func (a *App) UpdateStatus(ctx context.Context, id string, status int) error {
//	result := entity.GetDemoDB(ctx, a.DB).Where("id=?", id).Update("status", status)
//	if err := result.Error; err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
