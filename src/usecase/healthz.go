package usecase

import (
	"context"

	"go.uber.org/zap"
)

type HealthzRepoInterface interface {
	Readiness(ctx context.Context) error
	Liveness(ctx context.Context) error
}

type HealthzUseCaseInterface interface {
	Readiness(ctx context.Context) error
	Liveness(ctx context.Context) error
}

type HealthzUseCase struct {
	repo   HealthzRepoInterface
	logger *zap.Logger
}

// New Usecase
func NewHealthzUseCase(repo HealthzRepoInterface, logger *zap.Logger) HealthzUseCaseInterface {
	return &HealthzUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (uc *HealthzUseCase) Readiness(ctx context.Context) error {
	uc.logger.Debug("Readiness UC", zap.Object("ctx", ContextKeyLogger(ctx)))
	return uc.repo.Readiness(ctx)
}

func (uc *HealthzUseCase) Liveness(ctx context.Context) error {
	uc.logger.Debug("Liveness UC", zap.Object("ctx", ContextKeyLogger(ctx)))
	return uc.repo.Liveness(ctx)
}
