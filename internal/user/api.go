package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
	g   *gin.Engine
	http *http.Server
	srv Service
}

func NewApi(srv Service) *Api {
	g := gin.Default()
	httpServer := &http.Server{
		Addr: ":8004",
		Handler: g,
	}
	a := &Api{
		srv: srv,
		g: g,
		http: httpServer,
	}
	a.InitRouter(a.g)
	return a
}

func NewApiDebug(srv Service, g *gin.Engine) *Api {
	a := &Api{
		srv: srv,
		g: g,
	}
	a.InitRouter(a.g)
	return a
}

func (a *Api) Start() {
	err := a.http.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (a *Api) Shutdown(ctx context.Context) error {
	return a.http.Shutdown(ctx)
}

func (a *Api) auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.Abort()
			return
		}

		u, err := a.srv.Auth(token)
		if err != nil {
			c.Abort()
			return
		}

		c.Set("user", u)
		c.Next()
	}
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

	exists, err := a.srv.Exists(arg.Username)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	if exists {
		c.JSON(200, gin.H{"code": 0, "message": "用户名已存在"})
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

	token, err := a.srv.Login(arg.Username, arg.Password)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": "success", "data": token})
}

func (a *Api) AddFriend(c *gin.Context) {
	var (
		arg struct {
			Friendname string `json:"friendname"`
		}
	)

	user, _ := c.Get("user")
	u := user.(*User)

	err := c.ShouldBind(&arg)
	if err != nil || arg.Friendname == "" {
		c.JSON(200, gin.H{"code": 400, "message": "Friendname为空"})
		return
	}

	err = a.srv.AddFriend(u.Username, arg.Friendname)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": "success"})
}

func (a *Api) DeleteFriend(c *gin.Context) {
	var (
		arg struct {
			FriendName string `json:"friendname"`
		}
	)

	user, _ := c.Get("user")
	u := user.(*User)

	err := c.ShouldBind(&arg)
	if err != nil ||  arg.FriendName == "" {
		c.JSON(200, gin.H{"code": 400, "message": "friendname为空"})
		return
	}

	err = a.srv.RemoveFriend(u.Username, arg.FriendName)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": "success"})
}

func (a *Api) GetFriends(c *gin.Context) {
	user, _ := c.Get("user")
	u := user.(*User)
	friends, err := a.srv.GetFriends(u.Username)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": "success", "data": friends})
}
