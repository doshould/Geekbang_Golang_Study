package bll

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
)

// IApp应用逻辑接口
type IApp interface {
	// 查询数据
	Query(ctx context.Context, params schema.AppQueryParam, opts ...schema.AppQueryOptions) (*schema.AppQueryResult, error)
	//// 查询指定数据
	//Get(ctx context.Context, id string, opts ...schema.AppQueryOptions) (*schema.App, error)
	//// 创建数据
	//Create(ctx context.Context, item schema.App) (*schema.IDResult, error)
	//// 更新数据
	//Update(ctx context.Context, id string, item schema.App) error
}
