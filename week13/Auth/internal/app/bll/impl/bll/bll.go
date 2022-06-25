package bll

import "github.com/google/wire"

// BllSet bll注入
var BllSet = wire.NewSet(
	AccountSet,
	AppSet,
	LoginSet,
	ResourceSet,
	RoleSet,
)
