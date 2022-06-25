package model

import "github.com/google/wire"

// ModelSet model注入
var ModelSet = wire.NewSet(
	AccountSet,
	AccountRoleSet,
	AccountSystemSet,
	AppSet,
	ResourceSet,
	RoleSet,
	TransSet,
	RoleResourceSet,
)
