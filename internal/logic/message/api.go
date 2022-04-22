package message

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"maoim/internal/logic/user"
	"maoim/internal/pkg/resp"
	"net/http"
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
			UserId string `json:"user_id"`
			Op     int32  `json:"op"`
			Seq    int32  `json:"seq"`
			Body   string `json:"body"`
		}
	)
	err := c.BindJSON(&arg)
	if err != nil {
		resp.Response(c, http.StatusBadRequest, "参数格式错误", nil)
		return
	}

	u, _ := c.Get("user")
	us := u.(*user.User)

	err = a.srv.PushMsg(&PushMsgBo{
		Key:  arg.UserId,
		Op:   arg.Op,
		Seq:  arg.Seq,
		Body: arg.Body,
		u:    us,
	})
	if err != nil {
		resp.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	resp.SuccessResponse(c, nil)
}

func (a *Api) PullMsg(c *gin.Context) {
	userId := c.Query("user_id")
	if userId == "" {
		resp.Response(c, http.StatusBadRequest, "缺少userId", nil)
		return
	}

	u, _ := c.Get("user")
	us, _ := u.(*user.User)

	msg, err := a.srv.PullMsg(us.ID, userId)
	if err != nil {
		resp.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	resp.SuccessResponse(c, msg)
}

func (a *Api) Shutdown() error {
	a.grpc.GracefulStop()
	return nil
}
