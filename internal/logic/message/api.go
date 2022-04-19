package message

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"maoim/internal/logic/user"
)

type Api struct {
	srv  Service
	grpc *grpc.Server
}

func NewApi(srv Service, g *gin.Engine) *Api {
	a := &Api{
		srv:  srv,
		grpc: NewMessageGrpcServer(srv),
	}
	a.InitRouter(g)
	return a
}

func (a *Api) PushMsg(c *gin.Context) {
	var (
		arg struct {
			UserId  string `json:"user_id"`
			Op   int32  `json:"op"`
			Seq  int32  `json:"seq"`
			Body string `json:"body"`
		}
	)
	err := c.BindJSON(&arg)
	if err != nil {
		log.Println(arg)
		c.JSON(400, gin.H{"code": 400, "message": "参数格式错误"})
		return
	}

	u, exists := c.Get("user")
	if !exists {
		log.Println(arg)
		c.JSON(401, gin.H{"code": 401, "message": "no auth"})
		return
	}
	us, _ := u.(*user.User)

	err = a.srv.PushMsg(&PushMsgBo{
		Key:  arg.UserId,
		Op:   arg.Op,
		Seq:  arg.Seq,
		Body: arg.Body,
		u:    us,
	})
	if err != nil {
		log.Println(arg)
		c.JSON(500, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": "success"})
}

func (a *Api) PullMsg(c *gin.Context) {
	u, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"code": 401, "message": "no auth"})
		return
	}
	us, _ := u.(*user.User)

	msg, err := a.srv.PullMsg(us.ID)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 200, "message": "success", "data": msg})
}
