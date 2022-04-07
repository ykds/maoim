package user

import (
	"github.com/gin-gonic/gin"
)

func (a *Api) InitRouter(c *gin.Engine) {
	base := c.Group("/users")
	base.POST("/register", a.Register)
	base.POST("/login", a.Login)

	group := base.Group("/friends", a.auth())
	group.POST("/add", a.AddFriend)
	group.POST("/del", a.DeleteFriend)
	group.GET("/list/", a.GetFriends)
}
