package usecase

import (
	"github.com/google/wire"
)

var BizProviderSet = wire.NewSet(
	NewHealthzUseCase,
	NewUserUseCase,
)
