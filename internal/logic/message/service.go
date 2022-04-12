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
	PushMsg(bo *PushMsgBo) error
}

type service struct {
	d Dao
}

func NewService(d Dao) Service {
	return &service{d: d}
}

func (s *service) PushMsg(bo *PushMsgBo) error {
	for _, k := range bo.Keys {
		isFriend, err := s.d.IsFriend(context.Background(), bo.u.Username, k)
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