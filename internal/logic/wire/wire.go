package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maoim/internal/logic"
	"maoim/internal/logic/conf"
	"maoim/internal/logic/message"
	"maoim/internal/logic/user"
	"maoim/internal/pkg/grpc/comet"
	"maoim/pkg/mysql"
	"maoim/pkg/redis"
)

func Inital(c *conf.Config, rdb *redis.Redis, db *mysql.Mysql, g *gin.Engine) *logic.Server {
	panic(wire.Build(message.Provider, user.Provider, comet.NewCometGrpcClient, logic.New))
}
