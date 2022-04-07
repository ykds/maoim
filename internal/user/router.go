package user

import (
	"github.com/gin-gonic/gin"
)

func (a *Api) InitRouter(c *gin.Engine) {
	group := c.Group("/users")
	group.POST("/register", a.Register)
	group.POST("/login", a.Login)
}
