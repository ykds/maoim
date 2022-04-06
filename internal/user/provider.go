package user

import (
	"github.com/google/wire"
	"maoim/internal/user/http"
)

var Provider = wire.NewSet(NewDao, NewService, http.NewApi)
