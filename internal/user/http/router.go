package http

import (
	"github.com/gin-gonic/gin"
	"maoim/internal/pkg/middleware"
)

func (a *Api) InitRouter(c *gin.Engine) {
	group := c.Group("/users", middleware.Auth(a.srv.GetUser))
	group.POST("/register", a.Register)
	group.POST("/login", a.Login)
}
