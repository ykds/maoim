package logic

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	pb "maoim/api/comet"
	"maoim/api/protocal"
	"maoim/internal/user"
	"strconv"
)

func (s *Server) PushMsg(c *gin.Context) {
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
	us, _ := u.(*user.User)
	req := &pb.PushMsgReq{
		Keys: arg.Keys,
		PushMsg: &pb.PushMsg{
			FromKey: strconv.FormatInt(us.ID, 10),
			FromWho: us.Username,
			Proto: &protocal.Proto{
				Op:   arg.Op,
				Seq:  arg.Seq,
				Body: arg.Body,
			},
		},
	}

	_, err = s.comet.PushMsg(context.Background(), req)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"code": 500, "message": "Internal Error"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": "success"})
}