package entity

import (
	"context"
	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/util"
	"github.com/jinzhu/gorm"
)

// Role 角色实体
type Role struct {
	Model
	CommonModel
	AppKey  	string  `gorm:"column:appKey;size:30;not null;"` 	// 应用标识
	Name     	string  `gorm:"column:name;size:30;not null;"` 		// 角色名称
}

// GetRoleDB 获取角色存储
func GetRoleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Role))
}

//----------------------------------------------------------
// SchemaRole 角色对象
type SchemaRole schema.Role

// ToRole 转换为角色实体
func (a SchemaRole) ToRole() *Role {
	item := new(Role)
	util.StructMapToStruct(a, item)
	return item
}

// TableName 表名
func (a Role) TableName() string {
	return a.Model.TableName("Role")
}

// ToSchemaRole 转换为角色对象
func (a Role) ToSchemaRole() *schema.Role {
	item := new(schema.Role)
	util.StructMapToStruct(a, item)
	return item
}

//----------------------------------------------------------
// Roles 角色实体列表
type Roles []*Role

// ToSchemaRoles 转换为角色对象列表
func (a Roles) ToSchemaRoles() []*schema.Role {
	list := make([]*schema.Role, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaRole()
	}
	return list
}
