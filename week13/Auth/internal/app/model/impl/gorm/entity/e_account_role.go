package entity

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/util"
	"github.com/jinzhu/gorm"
)

// AccountRole 用户角色关联实体
type AccountRole struct {
	Model
	CommonModel
	AccountKey 	string `gorm:"column:accountKey;size:30;not null;"` // 账号标识
	RoleID 		int `gorm:"column:roleID;size:11;not null;"` 		// 角色ID
}

// GetAccountRoleDB 获取账号角色关联存储
func GetAccountRoleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(AccountRole))
}

// TableName 表名
func (a AccountRole) TableName() string {
	return a.Model.TableName("Account_Role")
}

// ToSchemaAccountRole 转换为账号角色对象
func (a AccountRole) ToSchemaAccountRole() *schema.AccountRole {
	item := new(schema.AccountRole)
	util.StructMapToStruct(a, item)
	return item
}

//----------------------------------------------------------
// SchemaAccountRole 账号角色
type SchemaAccountRole schema.AccountRole

// ToAccountRole 转换为账号角色实体
func (a SchemaAccountRole) ToAccountRole() *AccountRole {
	item := new(AccountRole)
	util.StructMapToStruct(a, item)
	return item
}

//----------------------------------------------------------
// AccountRoles 用户角色关联列表
type AccountRoles []*AccountRole

// ToSchemaAccountRoles 转换为用户角色对象列表
func (a AccountRoles) ToSchemaAccountRoles() []*schema.AccountRole {
	list := make([]*schema.AccountRole, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaAccountRole()
	}
	return list
}
