package bll

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/bll"
	"github.com/LyricTian/gin-admin/v6/internal/app/model"
	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/google/wire"
)

// App 应用管理
type App struct {
	AppModel model.IApp
}

var _ bll.IApp = (*App)(nil)

// AppSet 注入App
var AppSet = wire.NewSet(wire.Struct(new(App), "*"), wire.Bind(new(bll.IApp), new(*App)))

// Query 查询数据
func (a *App) Query(ctx context.Context, params schema.AppQueryParam, opts ...schema.AppQueryOptions) (*schema.AppQueryResult, error) {
	return a.AppModel.Query(ctx, params, opts...)
}

//// Get 查询指定数据
//func (a *Demo) Get(ctx context.Context, id string, opts ...schema.DemoQueryOptions) (*schema.Demo, error) {
//	item, err := a.DemoModel.Get(ctx, id, opts...)
//	if err != nil {
//		return nil, err
//	} else if item == nil {
//		return nil, errors.ErrNotFound
//	}
//
//	return item, nil
//}
//
//func (a *Demo) checkCode(ctx context.Context, code string) error {
//	result, err := a.DemoModel.Query(ctx, schema.DemoQueryParam{
//		PaginationParam: schema.PaginationParam{
//			OnlyCount: true,
//		},
//		Code: code,
//	})
//	if err != nil {
//		return err
//	} else if result.PageResult.Total > 0 {
//		return errors.New400Response("编号已经存在")
//	}
//
//	return nil
//}
//
//// Create 创建数据
//func (a *Demo) Create(ctx context.Context, item schema.Demo) (*schema.IDResult, error) {
//	err := a.checkCode(ctx, item.Code)
//	if err != nil {
//		return nil, err
//	}
//
//	item.ID = iutil.NewID()
//	err = a.DemoModel.Create(ctx, item)
//	if err != nil {
//		return nil, err
//	}
//
//	return schema.NewIDResult(item.ID), nil
//}
//
//// Update 更新数据
//func (a *Demo) Update(ctx context.Context, id string, item schema.Demo) error {
//	oldItem, err := a.DemoModel.Get(ctx, id)
//	if err != nil {
//		return err
//	} else if oldItem == nil {
//		return errors.ErrNotFound
//	} else if oldItem.Code != item.Code {
//		if err := a.checkCode(ctx, item.Code); err != nil {
//			return err
//		}
//	}
//	item.ID = oldItem.ID
//	item.Creator = oldItem.Creator
//	item.CreatedAt = oldItem.CreatedAt
//
//	return a.DemoModel.Update(ctx, id, item)
//}
//
//// Delete 删除数据
//func (a *Demo) Delete(ctx context.Context, id string) error {
//	oldItem, err := a.DemoModel.Get(ctx, id)
//	if err != nil {
//		return err
//	} else if oldItem == nil {
//		return errors.ErrNotFound
//	}
//
//	return a.DemoModel.Delete(ctx, id)
//}
//
//// UpdateStatus 更新状态
//func (a *Demo) UpdateStatus(ctx context.Context, id string, status int) error {
//	oldItem, err := a.DemoModel.Get(ctx, id)
//	if err != nil {
//		return err
//	} else if oldItem == nil {
//		return errors.ErrNotFound
//	}
//
//	return a.DemoModel.UpdateStatus(ctx, id, status)
//}
