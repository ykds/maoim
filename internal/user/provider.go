package user

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewDao, NewService, NewApi)
