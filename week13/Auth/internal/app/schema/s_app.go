package schema

import "time"

// App 应用对象
type App struct {
	ID        		int    		`json:"id"`                                    // 主键ID
	AppKey    		string    	`json:"appKey" binding:"required"`               // 编号
	Name      		string    	`json:"name" binding:"required"`               // 名称
	AccountType   	string    	`json:"accountType"`                     // 账号类型
	CreateTime		time.Time	`json:"createTime"`                		// 创建时间
	UpdateTime		time.Time	`json:"updateTime"`                		// 更新时间
}

// AppQueryParam 查询条件
type AppQueryParam struct {
	PaginationParam
	AppKey 			string `form:"-"`          // 应用标识
	AccountType 	string `form:"-"`          // 账号类型
}

// AppQueryOptions 应用对象查询可选参数项
type AppQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// AppQueryResult 示例对象查询结果
type AppQueryResult struct {
	Data       []*App
	PageResult *PaginationResult
}