package entity

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/util"
	"github.com/jinzhu/gorm"
)

// RoleResource 角色资源实体
type RoleResource struct {
	Model
	CommonModel
	RoleID   	int `gorm:"column:roleID;size:11;not null;"`   		// 角色ID
	ResourceID 	int `gorm:"column:resourceID;size:11;not null;"`   	// 资源ID
}

// GetRoleResourceDB 角色资源存储
func GetRoleResourceDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(RoleResource))
}

// TableName 表名
func (a RoleResource) TableName() string {
	return a.Model.TableName("Role_Resource")
}

// ToSchemaRoleResource 转换为角色资源对象
func (a RoleResource) ToSchemaRoleResource() *schema.RoleResource {
	item := new(schema.RoleResource)
	util.StructMapToStruct(a, item)
	return item
}

//----------------------------------------------------------
// SchemaRoleResource 角色资源
type SchemaRoleResource schema.RoleResource

// ToRoleResource 转换为角色资源实体
func (a SchemaRoleResource) ToRoleResource() *RoleResource {
	item := new(RoleResource)
	util.StructMapToStruct(a, item)
	return item
}

//----------------------------------------------------------
// RoleResources 角角色资源列表
type RoleResources []*RoleResource

// ToSchemaRoleResources 转换为角色资源对象列表
func (a RoleResources) ToSchemaRoleResources() []*schema.RoleResource {
	list := make([]*schema.RoleResource, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaRoleResource()
	}
	return list
}
