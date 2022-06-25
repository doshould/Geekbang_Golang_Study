package entity

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/util"
	"github.com/jinzhu/gorm"
)

// Resource 资源实体
type Resource struct {
	Model
	CommonModel
	AppKey    	string  `gorm:"column:appKey;size:30;not null;"` 				// 应用标识
	Name   		string  `gorm:"column:name;size:30;not null;"`      			// 名称
	Type     	string 	`gorm:"column:type;size:30;not null;"`                  // 资源类型：api
	Feature    	string 	`gorm:"column:feature;size:255;not null;"`       		// 特征值：如：uri
	Method    	string 	`gorm:"column:method;size:255;not null;"`       		// 请求类型
	Pid  	 	int 	`gorm:"column:pid;size:11;not null;default:0"`       	// 父级资源
	Sequence 	int 	`gorm:"column:sequence;size:11;not null;default:0"`    	// 资源排序，值越大越靠前
}

// GetResourceDB 获取资源存储
func GetResourceDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Resource))
}

//----------------------------------------------------------
// SchemaResource 资源对象
type SchemaResource schema.Resource

// ToResource 转换为资源实体
func (a SchemaResource) ToResource() *Resource {
	item := new(Resource)
	util.StructMapToStruct(a, item)
	return item
}

// TableName 表名
func (a Resource) TableName() string {
	return a.Model.TableName("Resource")
}

// ToSchemaResource 转换为资源对象
func (a Resource) ToSchemaResource() *schema.Resource {
	item := new(schema.Resource)
	util.StructMapToStruct(a, item)
	return item
}

//----------------------------------------------------------
// Resources 资源实体列表
type Resources []*Resource

// ToSchemaResources 转换为菜单对象列表
func (a Resources) ToSchemaResources() []*schema.Resource {
	list := make([]*schema.Resource, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaResource()
	}
	return list
}
