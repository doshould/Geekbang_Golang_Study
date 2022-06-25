package schema

import (
	"context"
	"time"

	"github.com/LyricTian/gin-admin/v6/internal/app/config"
	"github.com/LyricTian/gin-admin/v6/pkg/util"
)

// Account 账号对象
type Account struct {
	AccountKey			string		`json:"accountKey" binding:"required"`	// 唯一标识
	Username  			string  	`json:"username" binding:"required"`   	// 用户名
	Password  			string 		`json:"password" binding:"required"`    // 密码
	AccountType     	string  	`json:"accountType" binding:"required"`	// 账号类型
	MobilePhone     	string  	`json:"mobilePhone"`		            // 手机号
	Email     			string  	`json:"email"`                         	// 邮箱
	CreateTime			time.Time	`json:"createTime"`                		// 创建时间
	UpdateTime			time.Time	`json:"updateTime"`                		// 更新时间
	//UserRoles 			UserRoles 	`json:"user_roles"`    				// 角色授权
}

// Account 账号对象
type AccountCreateParam struct {
	AccountType     	string  	`json:"accountType" binding:"required"`	// 账号类型
	Username  			string  	`json:"username" binding:"required"`   	// 用户名
	Password  			string 		`json:"password" binding:"required"`    // 密码
}

// GetRootAccount 获取root账号
func GetRootAccount() *Account {
	root := config.C.Root
	return &Account{
		AccountKey: root.Username,
		Username: root.Username,
		Password: util.MD5HashString(root.Password),
	}
}

// CheckIsRootAccount 检查是否是root账号
func CheckIsRootUser(ctx context.Context, AccountKey string) bool {
	return GetRootAccount().AccountKey == AccountKey
}

func (a *Account) String() string {
	return util.JSONMarshalToString(a)
}

// CleanSecure 清理安全数据
func (a *Account) CleanSecure() *Account {
	a.Password = ""
	return a
}

// AccountQueryParam 查询条件
type AccountQueryParam struct {
	PaginationParam
	AccountType string   `form:"accountType"`   // 账号类型
	AccountKey	string   `form:"accountKey"`   	// 账号标识
	Username   	string   `form:"username"`   	// 账号名
	Password   	string   `form:"password"`   	// 密码


	AppKey		string   `form:"appKey"`   		// 应用标识
	RoleIDs   	[]string `form:"-"`          	// 角色ID列表
}

// AccountQueryOptions 查询可选参数项
type AccountQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// Users 用户对象列表
type Accounts []*Account

// AccountQueryResult 查询结果
type AccountQueryResult struct {
	Data       Accounts
	PageResult *PaginationResult
}

//// ToShowResult 转换为显示结果
//func (a AccountQueryResult) ToShowResult(mUserRoles map[string]UserRoles, mRoles map[string]*Role) *UserShowQueryResult {
//	return &UserShowQueryResult{
//		PageResult: a.PageResult,
//		Data:       a.Data.ToUserShows(mUserRoles, mRoles),
//	}
//}
//
//
//// ToIDs 转换为唯一标识列表
//func (a Accounts) ToIDs() []string {
//	idList := make([]string, len(a))
//	for i, item := range a {
//		idList[i] = item.AccountKey
//	}
//	return idList
//}
//
//// ToUserShows 转换为用户显示列表
//func (a Accounts) ToUserShows(mUserRoles map[string]UserRoles, mRoles map[string]*Role) UserShows {
//	list := make(UserShows, len(a))
//	for i, item := range a {
//		showItem := new(UserShow)
//		util.StructMapToStruct(item, showItem)
//		for _, roleID := range mUserRoles[item.AccountKey].ToRoleIDs() {
//			if v, ok := mRoles[roleID]; ok {
//				showItem.Roles = append(showItem.Roles, v)
//			}
//		}
//		list[i] = showItem
//	}
//
//	return list
//}

// ----------------------------------------AccountSystem--------------------------------------
// AccountSystem 账号体系
type AccountSystem struct {
	ID     		string `json:"id"`      	// 主键ID
	AccountType string `json:"accountType"` // 账号类型
}

// AccountSystemQueryParam 查询条件
type AccountSystemQueryParam struct {
	PaginationParam
	AccountType  string   // 账号类型
}

// AccountSystemQueryResult 查询结果
type AccountSystemQueryResult struct {
	Data       AccountSystems
	PageResult *PaginationResult
}

// AccountSystems 账号体系列表
type AccountSystems []*AccountSystem

// ----------------------------------------AccountRole--------------------------------------

// AccountRole 账号角色
type AccountRole struct {
	ID     		int `json:"id"`      	// 唯一标识
	AccountKey 	string `json:"accountKey"` 	// 账号标识
	RoleID 		int `json:"roleID"` 		// 角色ID
}

// AccountRoleQueryParam 查询条件
type AccountRoleQueryParam struct {
	PaginationParam
	AccountKey  string   // 账号标识
	AccountKeys []string // 账号标识列表
}

// AccountRoleQueryOptions 查询可选参数项
type AccountRoleQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// AccountRoleQueryResult 查询结果
type AccountRoleQueryResult struct {
	Data       AccountRoles
	PageResult *PaginationResult
}

// AccountRoles 角色菜单列表
type AccountRoles []*AccountRole

//// ToMap 转换为map
//func (a AccountRoles) ToMap() map[string]*AccountRole {
//	m := make(map[string]*AccountRole)
//	for _, item := range a {
//		m[item.RoleID] = item
//	}
//	return m
//}

//// ToRoleIDs 转换为角色ID列表
//func (a UserRoles) ToRoleIDs() []string {
//	list := make([]string, len(a))
//	for i, item := range a {
//		list[i] = item.RoleID
//	}
//	return list
//}
//

// ToUserIDMap 转换为用户ID映射
func (a AccountRoles) ToAccountKeyMap() map[string]AccountRoles {
	m := make(map[string]AccountRoles)
	for _, item := range a {
		m[item.AccountKey] = append(m[item.AccountKey], item)
	}
	return m
}

//// ----------------------------------------UserShow--------------------------------------
//
//// UserShow 用户显示项
//type UserShow struct {
//	ID        string    `json:"id"`         // 唯一标识
//	UserName  string    `json:"user_name"`  // 用户名
//	RealName  string    `json:"real_name"`  // 真实姓名
//	Phone     string    `json:"phone"`      // 手机号
//	Email     string    `json:"email"`      // 邮箱
//	Status    int       `json:"status"`     // 用户状态(1:启用 2:停用)
//	CreatedAt time.Time `json:"created_at"` // 创建时间
//	Roles     []*Role   `json:"roles"`      // 授权角色列表
//}
//
//// UserShows 用户显示项列表
//type UserShows []*UserShow
//
//// UserShowQueryResult 用户显示项查询结果
//type UserShowQueryResult struct {
//	Data       UserShows
//	PageResult *PaginationResult
//}
