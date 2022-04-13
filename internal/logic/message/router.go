package message

import (
	"github.com/gin-gonic/gin"
	"maoim/internal/logic/user"
)

func (a *Api) InitRouter(g *gin.Engine) {
	base := g.Group("/msg", user.Auth(a.srv.GetUserService()))
	base.POST("/push", a.PushMsg)
}
