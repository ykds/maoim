package user

import (
	"github.com/gin-gonic/gin"
)

func (a *Api) InitRouter(g *gin.Engine) {
	g.Static("/static", "./static")
	g.POST("/upload", Auth(a.srv), a.UploadFile)

	base := g.Group("/users")
	base.POST("/register", a.Register)
	base.POST("/login", a.Login)
	base.GET("/info", Auth(a.srv), a.GetUserInfo)
	base.GET("/mine", Auth(a.srv), a.MineInfo)
	base.PUT("/update/mine", Auth(a.srv), a.UpdateMineInfo)

	group := base.Group("/friends", Auth(a.srv))
	//group.POST("/add", a.AddFriend)
	group.POST("/del", a.DeleteFriend)
	group.GET("/list", a.GetFriends)
	group.POST("/apply", a.ApplyFriend)
	group.GET("/apply/list", a.ListApplyRecord)
	group.GET("/apply/list/offset", a.ListOffsetApplyRecord)
	group.POST("/apply/agree", a.AgreeFriendShipApply)
}
