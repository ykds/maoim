package message

import (
	"context"
	"github.com/pkg/errors"
	pb "maoim/api/comet"
	"maoim/api/protocal"
	"maoim/internal/logic/user"
	"strconv"
)

type PushMsgBo struct {
	Keys []string
	Op   int32
	Seq  int32
	Body string

	u *user.User
}

type Service interface {
	GetUserService() user.Service
	PushMsg(bo *PushMsgBo) error
}

type service struct {
	userSrv user.Service
	d Dao
}

func NewService(d Dao, userSrv user.Service) Service {
	return &service{
		d: d,
		userSrv: userSrv,
	}
}

func (s *service) GetUserService() user.Service {
	return s.userSrv
}

func (s *service) PushMsg(bo *PushMsgBo) error {
	for _, k := range bo.Keys {
		isFriend, err := s.userSrv.IsFriend(bo.u.Username, k)
		if err != nil {
			return err
		}
		if !isFriend {
			return errors.New(k + "不是好友")
		}
	}

	req := &pb.PushMsgReq{
		Keys: bo.Keys,
		PushMsg: &pb.PushMsg{
			FromKey: strconv.FormatInt(bo.u.ID, 10),
			FromWho: bo.u.Username,
			Proto: &protocal.Proto{
				Op:   bo.Op,
				Seq:  bo.Seq,
				Body: bo.Body,
			},
		},
	}

	return s.d.PushMsg(context.Background(), req)
}