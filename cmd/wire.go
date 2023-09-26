//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"application/src/controller"
	"application/src/repository"
	"application/src/server"
	"application/src/usecase"
	"github.com/google/wire"

	"application/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func wireApp(cfg *config.ViperConfig, logger *zap.Logger) (*gin.Engine, error) {
	panic(wire.Build(
		repository.DataProviderSet,
		server.ServerProviderSet,
		controller.ServiceProviderSet,
		usecase.BizProviderSet,
	))
}
