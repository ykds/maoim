package user

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"maoim/api/comet"
	"maoim/internal/pkg/resp"
	"net/http"
	"sort"
)

type Api struct {
	srv         Service
	grpc        *grpc.Server
	cometClient comet.CometClient
}

func NewApi(srv Service, g *gin.Engine) *Api {
	a := &Api{
		srv:  srv,
		grpc: NewUserGrpcServer(srv),
	}
	a.InitRouter(g)
	return a
}

func (a *Api) Register(c *gin.Context) {
	var (
		arg struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		res struct {
			UserId string `json:"user_id"`
		}
	)

	err := c.ShouldBind(&arg)
	if err != nil || arg.Username == "" || arg.Password == "" {
		resp.Response(c, http.StatusBadRequest, "username或password为空", nil)
		return
	}

	u, err := a.srv.Register(arg.Username, arg.Password)
	if err != nil {
		resp.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	res.UserId = u.ID
	resp.SuccessResponse(c, res)
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
		resp.Response(c, http.StatusBadRequest,  "username或password为空", nil)
		return
	}

	token, err := a.srv.Login(arg.Username, arg.Password)
	if err != nil {
		resp.Response(c, http.StatusInternalServerError,  err.Error(), nil)
		return
	}
	resp.SuccessResponse(c, token)
}

//func (a *Api) AddFriend(c *gin.Context) {
//	var (
//		arg struct {
//			Friendname string `json:"friendname"`
//		}
//	)
//
//	user, _ := c.Get("user")
//	u := user.(*User)
//
//	err := c.ShouldBind(&arg)
//	if err != nil || arg.Friendname == "" {
//		c.JSON(200, gin.H{"code": 400, "message": "Friendname为空"})
//		return
//	}
//
//	err = a.srv.AddFriend(u.Username, arg.Friendname)
//	if err != nil {
//		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
//		return
//	}
//	c.JSON(200, gin.H{"code": 200, "message": "success"})
//}

func (a *Api) DeleteFriend(c *gin.Context) {
	var (
		arg struct {
			FUserId string `json:"f_user_id"`
		}
	)

	user, _ := c.Get("user")
	u := user.(*User)

	err := c.ShouldBind(&arg)
	if err != nil || arg.FUserId == "" {
		resp.Response(c, http.StatusBadRequest, "f_user_id为空", nil)
		return
	}

	err = a.srv.RemoveFriend(u.ID, arg.FUserId)
	if err != nil {
		resp.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	resp.SuccessResponse(c, nil)
}

func (a *Api) GetFriends(c *gin.Context) {
	user, _ := c.Get("user")
	u := user.(*User)
	friends, err := a.srv.GetFriends(u.ID)
	if err != nil {
		resp.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	resp.SuccessResponse(c, friends)
}

func (a *Api) GetUserService() Service {
	return a.srv
}

func (a *Api) ApplyFriend(c *gin.Context) {
	var (
		arg struct {
			OtherUsername string `json:"other_username"`
			Remark      string `json:"remark"`
		}
	)

	err := c.ShouldBind(&arg)
	if err != nil {
		resp.Response(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	get, _ := c.Get("user")
	u := get.(*User)

	err = a.srv.ApplyFriend(u.ID, arg.OtherUsername, arg.Remark)
	if err != nil {
		resp.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	resp.SuccessResponse(c, nil)
}

type ApplyRecord struct {
	ID string `json:"id"`
	ApplyUserId string `json:"apply_user_id"`
	ApplyUsername string `json:"apply_username"`
	AppliedUserId string `json:"applied_user_id"`
	Remark string `json:"remark"`
	ApplyTime string `json:"apply_time"`
	ApplyType int `json:"apply_type"`
}

func (a *Api) ListApplyRecord(c *gin.Context) {
	get, _ := c.Get("user")
	u := get.(*User)

	applyingList, err := a.srv.ListApplyRecord(u.ID, true)
	if err != nil {
		resp.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	applyedList, err := a.srv.ListApplyRecord(u.ID, false)
	if err != nil {
		resp.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	record := make([]*FriendShipApply, 0)
	record = append(record, applyingList...)
	record = append(record, applyedList...)
	sort.Slice(record, func(i, j int) bool {
		return record[i].CreatedAt.After(record[j].CreatedAt)
	})

	result := make([]ApplyRecord, len(record))
	for i, r := range record {
		applyType := 0
		if u.ID == r.OtherUserId {
			applyType = 1
		}
		result[i] = ApplyRecord{
			ID: r.ID,
			ApplyUserId: r.UserId,
			ApplyUsername: r.Username,
			AppliedUserId: r.OtherUserId,
			Remark: r.Remark,
			ApplyTime: r.CreatedAt.Format("2006-01-02 15:04:05"),
			ApplyType: applyType,
		}
	}
	resp.SuccessResponse(c, result)
}

func (a *Api) ListOffsetApplyRecord(c *gin.Context) {
	get, _ := c.Get("user")
	u := get.(*User)

	recordId := c.Query("record_id")
	if recordId == "" {
		resp.Response(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	applyingList, err := a.srv.ListOffsetApplyRecord(u.ID, recordId, true)
	if err != nil {
		resp.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	applyedList, err := a.srv.ListOffsetApplyRecord(u.ID, recordId, false)
	if err != nil {
		resp.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	record := make([]*FriendShipApply, 0)
	record = append(record, applyingList...)
	record = append(record, applyedList...)
	sort.Slice(record, func(i, j int) bool {
		return record[i].CreatedAt.After(record[j].CreatedAt)
	})
	result := make([]ApplyRecord, len(record))
	for i, r := range record {
		applyType := 0
		if u.ID == r.OtherUserId {
			applyType = 1
		}
		result[i] = ApplyRecord{
			ID: r.ID,
			ApplyUserId: r.UserId,
			ApplyUsername: r.Username,
			AppliedUserId: r.OtherUserId,
			Remark: r.Remark,
			ApplyTime: r.CreatedAt.Format("2006-01-02 15:04:05"),
			ApplyType: applyType,
		}
	}
	resp.SuccessResponse(c, result)
}

func (a *Api) AgreeFriendShipApply(c *gin.Context) {
	var (
		req struct {
			RecordId string `json:"record_id"`
		}
	)
	err := c.BindJSON(&req)
	if req.RecordId == "" {
		resp.Response(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	user, _ := c.Get("user")
	u := user.(*User)

	err = a.srv.AgreeFriendShipApply(u.ID, req.RecordId)
	if err != nil {
		resp.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	resp.SuccessResponse(c, nil)
}

func (a *Api) Shutdown() error {
	a.grpc.GracefulStop()
	return nil
}
