package entity

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/util"
	"github.com/jinzhu/gorm"
)

// Demo demo实体
type App struct {
	Model
	CommonModel
	AppKey 	string  `gorm:"column:appKey;size:30;not null;"`  	// 应用标识
	Name    string  `gorm:"column:name;size:30;not null;"` 		// 名称
}

// GetAppDB 获取应用存储
func GetAppDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(App))
}

// TableName 表名
func (a App) TableName() string {
	return a.Model.TableName("App")
}

// ToSchemaApp 转换为App对象
func (a App) ToSchemaApp() *schema.App {
	item := new(schema.App)
	util.StructMapToStruct(a, item)
	return item
}

//----------------------------------------------------------
// SchemaApp 应用对象
type SchemaApp schema.App

// ToApp 转换为App实体
func (a SchemaApp) ToApp() *App {
	item := new(App)
	util.StructMapToStruct(a, item)
	return item
}

//----------------------------------------------------------
// Apps App列表
type Apps []*App

// ToSchemaApps 转换为app对象列表
func (a Apps) ToSchemaApps() []*schema.App {
	list := make([]*schema.App, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaApp()
	}
	return list
}
