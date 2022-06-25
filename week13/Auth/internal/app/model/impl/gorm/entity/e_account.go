package entity

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/util"
	"github.com/jinzhu/gorm"
)

// NewAccount 创建用户实体
type NewAccount struct {
	Model
	AccountType string  `gorm:"column:accountType;size:30;not null;"` 				// 用户名
	AccountKey 	string  `gorm:"column:accountKey;size:255;not null;"` 				// 用户名
	Username 	string  `gorm:"column:username;size:255;not null;"` 				// 用户名
	Email    	string 	`gorm:"column:email;size:255;default:null;null;"`      	// 邮箱
	MobilePhone string 	`gorm:"column:mobilePhone;size:255;default:null;null;"`   // 手机号
	Password 	string  `gorm:"column:password;size:255;not null;"`      			// 密码md5签名
}

// Account 用户实体
type Account struct {
	CommonModel
	NewAccount
}

// GetAccountDB 获取账号存储
func GetAccountDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Account))
}

// TableName 表名
func (a NewAccount) TableName() string {
	return a.Model.TableName("Account")
}

// ToSchemaAccount 转换为用户对象
func (a Account) ToSchemaAccount() *schema.Account {
	item := new(schema.Account)
	util.StructMapToStruct(a, item)
	return item
}

//----------------------------------------------------------
// SchemaAccount 用户对象
type SchemaAccount schema.Account

// ToNewAccount 转换为创建用户实体
func (a SchemaAccount) ToNewAccount() *NewAccount {
	item := new(NewAccount)
	util.StructMapToStruct(a, item)
	return item
}

// ToNewAccount 转换为创建用户实体
func (a SchemaAccount) ToAccount() *Account {
	item := new(Account)
	util.StructMapToStruct(a, item)
	return item
}

//----------------------------------------------------------
// Accounts 用户实体列表
type Accounts []*Account

// ToSchemaAccounts 转换为用户对象列表
func (a Accounts) ToSchemaAccounts() []*schema.Account {
	list := make([]*schema.Account, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaAccount()
	}
	return list
}
