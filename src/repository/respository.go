package repository

import (
	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(
	NewDataSource,
	NewHealthzRepo,
	NewUserRepo,
)
