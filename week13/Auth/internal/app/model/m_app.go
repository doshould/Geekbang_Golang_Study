package model

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
)

// IApp 应用管理存储接口
type IApp interface {
	// 查询数据
	Query(ctx context.Context, params schema.AppQueryParam, opts ...schema.AppQueryOptions) (*schema.AppQueryResult, error)
	//// 查询指定数据
	//Get(ctx context.Context, id string, opts ...schema.AppQueryOptions) (*schema.App, error)
	//// 创建数据
	//Create(ctx context.Context, item schema.App) error
	//// 更新数据
	//Update(ctx context.Context, id string, item schema.App) error
	//// 删除数据
	//Delete(ctx context.Context, id string) error
	//// 更新父级路径
	//UpdateParentPath(ctx context.Context, id, parentPath string) error
	//// 更新状态
	//UpdateStatus(ctx context.Context, id string, status int) error
}
