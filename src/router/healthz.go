package router

import (
	"application/src/controller"
	"github.com/gin-gonic/gin"
)

func RegisterHealthzRoutes(s *controller.HealthzService, e *gin.Engine) {
	r := e.Group("/healthz")
	r.GET("/liveness", s.HealthzLiveness)
	r.GET("/readiness", s.HealthzReadiness)
}
