package entity

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/util"
	"github.com/jinzhu/gorm"
)

// AccountSystem 账号体系实体
type AccountSystem struct {
	Model
	CommonModel
	AccountType string `gorm:"column:accountType;size:30;not null;"` 	// 账号类型
}

// GetAccountSystemDB 菜单动作
func GetAccountSystemDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(AccountSystem))
}

// TableName 表名
func (a AccountSystem) TableName() string {
	return a.Model.TableName("AccountSystem")
}

// ToSchemaAccountSystem 转换为账号体系对象
func (a AccountSystem) ToSchemaAccountSystem() *schema.AccountSystem {
	item := new(schema.AccountSystem)
	util.StructMapToStruct(a, item)
	return item
}

//----------------------------------------------------------
// SchemaAccountSystem 账号体系
type SchemaAccountSystem schema.AccountSystem

// ToAccountSystem 转换为账号体系实体
func (a SchemaAccountSystem) ToAccountSystem() *AccountSystem {
	item := new(AccountSystem)
	util.StructMapToStruct(a, item)
	return item
}

//----------------------------------------------------------
// AccountSystems 账号体系列表
type AccountSystems []*AccountSystem

// ToSchemaAccountSystems 转换为账号体系对象列表
func (a AccountSystems) ToSchemaAccountSystems() []*schema.AccountSystem {
	list := make([]*schema.AccountSystem, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaAccountSystem()
	}
	return list
}
