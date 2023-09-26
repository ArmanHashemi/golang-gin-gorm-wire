package server

import (
	"application/config"
	"application/src/controller"
	"application/src/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewServer creates a new HTTP server and set up all routes.
func NewGinServer(
	cfg *config.ViperConfig,
	logger *zap.Logger,
	healthzSvc *controller.HealthzService,
	userSvc *controller.UserService,
) *gin.Engine {

	// gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	router.RegisterHealthzRoutes(healthzSvc, engine)
	router.RegisterUserRoutes(userSvc, engine)

	return engine

}
