//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"hello/internal/biz"
	"hello/internal/conf"
	"hello/internal/data"
	"hello/internal/server"
	"hello/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
}
