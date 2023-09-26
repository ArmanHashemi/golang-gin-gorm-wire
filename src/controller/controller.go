package controller

import (
	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	NewHealthzService,
	NewUserService,
)
