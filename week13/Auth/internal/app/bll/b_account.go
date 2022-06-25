package bll

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
)

// IAccount 账号管理业务逻辑接口
type IAccount interface {
	// 查询数据
	Query(ctx context.Context, params schema.AccountQueryParam, opts ...schema.AccountQueryOptions) (*schema.AccountQueryResult, error)
	// 创建数据
	Create(ctx context.Context, item schema.AccountCreateParam) (string, error)
}
