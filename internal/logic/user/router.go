package user

import (
	"github.com/gin-gonic/gin"
)

func (a *Api) InitRouter(g *gin.Engine) {
	base := g.Group("/users")
	base.POST("/register", a.Register)
	base.POST("/login", a.Login)

	group := base.Group("/friends", Auth(a.srv))
	group.POST("/add", a.AddFriend)
	group.POST("/del", a.DeleteFriend)
	group.GET("/list", a.GetFriends)
}
