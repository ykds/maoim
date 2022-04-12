package message

import (
	"github.com/gin-gonic/gin"
	"log"
	user2 "maoim/internal/logic/user"
)

type Api struct {
	srv Service
}

func (a *Api) PushMsg(c *gin.Context)  {
	var (
		arg struct {
			Keys []string `json:"keys"`
			Op   int32    `json:"op"`
			Seq  int32    `json:"seq"`
			Body string   `json:"body"`
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
	us, _ := u.(*user2.User)

	err = a.srv.PushMsg(&PushMsgBo{
		Keys: arg.Keys,
		Op: arg.Op,
		Seq: arg.Seq,
		Body: arg.Body,
		u: us,
	})
	if err != nil {
		log.Println(arg)
		c.JSON(500, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": "success"})
}