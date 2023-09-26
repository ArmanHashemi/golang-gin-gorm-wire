package repository

import (
	"application/src/usecase"
	"context"

	"go.uber.org/zap"
)

type HealthzRepo struct {
	logger *zap.Logger
	ds     *DataSource
}

func NewHealthzRepo(logger *zap.Logger, ds *DataSource) usecase.HealthzRepoInterface {
	return &HealthzRepo{
		logger: logger.With(zap.String("type", "Repo"), zap.String("controller", "Healthz")),
		ds:     ds,
	}
}

func (r *HealthzRepo) Readiness(ctx context.Context) error {
	r.logger.Debug("repo Readiness")
	return nil
}

func (r *HealthzRepo) Liveness(ctx context.Context) error {
	r.logger.Debug("repo Liveness")
	return nil
}
