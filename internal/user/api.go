package user

import (
	"github.com/gin-gonic/gin"
)

type Api struct {
	g *gin.Engine
	srv Service
}

func NewApi(srv Service) *Api {
	a := &Api{srv: srv}
	a.InitRouter(a.g)
	return a
}

func (a *Api) Register(c *gin.Context) {
	var (
		arg struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
	)

	err := c.ShouldBind(&arg)
	if err != nil || arg.Username == "" || arg.Password == "" {
		c.JSON(200, gin.H{"code": 400, "message": "username或password为空"})
		return
	}

	u, err := a.srv.Register(arg.Username, arg.Password)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": "register success", "data": u.ID})
}

func (a *Api) Login(c *gin.Context) {
	var (
		arg struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
	)

	err := c.ShouldBind(&arg)
	if err != nil || arg.Username == "" || arg.Password == "" {
		c.JSON(200, gin.H{"code": 400, "message": "username或password为空"})
		return
	}

	exists, err := a.srv.Exists(arg.Username)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	if exists {
		c.JSON(200, gin.H{"code": 0, "message": "用户名已存在"})
		return
	}

	cookie, err := a.srv.Login(arg.Username, arg.Password)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": "success", "data": cookie})
}

func (a *Api) Logout(c *gin.Context) {

}
