package bll

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
)

// IResource 资源管理业务逻辑接口
type IResource interface {
	//// 初始化菜单数据
	//InitData(ctx context.Context, dataFile string) error
	// 查询数据
	Query(ctx context.Context, params schema.ResourceQueryParam, opts ...schema.ResourceQueryOptions) (*schema.ResourceQueryResult, error)
	//// 查询指定数据
	//Get(ctx context.Context, id string, opts ...schema.ResourceQueryOptions) (*schema.Resource, error)
	//// 创建数据
	//Create(ctx context.Context, item schema.Resource) (*schema.IDResult, error)
	//// 更新数据
	//Update(ctx context.Context, id string, item schema.Resource) error
	//// 删除数据
	//Delete(ctx context.Context, id string) error
	//// 更新状态
	//UpdateStatus(ctx context.Context, id string, status int) error
}
