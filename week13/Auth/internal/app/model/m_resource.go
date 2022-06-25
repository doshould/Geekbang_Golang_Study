package model

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
)

// IResource 应用资源管理存储接口
type IResource interface {
	// 查询数据
	Query(ctx context.Context, params schema.ResourceQueryParam, opts ...schema.ResourceQueryOptions) (*schema.ResourceQueryResult, error)
	//// 查询指定数据
	//Get(ctx context.Context, id string, opts ...schema.MenuActionQueryOptions) (*schema.MenuAction, error)
	//// 创建数据
	//Create(ctx context.Context, item schema.MenuAction) error
	//// 更新数据
	//Update(ctx context.Context, id string, item schema.MenuAction) error
	//// 删除数据
	//Delete(ctx context.Context, id string) error
	//// 根据菜单ID删除数据
	//DeleteByMenuID(ctx context.Context, menuID string) error
}
