package controller

import (
	"application/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HealthzService struct {
	uc     usecase.HealthzUseCaseInterface
	logger *zap.Logger
}

// New
func NewHealthzService(uc usecase.HealthzUseCaseInterface, logger *zap.Logger) *HealthzService {
	return &HealthzService{
		uc:     uc,
		logger: logger.With(zap.String("type", "HealthzService"), zap.String("version", "v1")),
	}
}

// Healthz Liveness
func (s *HealthzService) HealthzLiveness(c *gin.Context) {

	s.logger.Debug("HealthzLiveness")
	s.uc.Liveness(GetContextFromGin(c))
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// Healthz Readiness
func (s *HealthzService) HealthzReadiness(c *gin.Context) {
	s.logger.Debug("HealthzReadiness")
	s.uc.Readiness(GetContextFromGin(c))
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
