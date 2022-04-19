package message

import (
	"context"
	"github.com/pkg/errors"
	pb "maoim/api/comet"
	"maoim/api/protocal"
	"maoim/internal/logic/user"
)

const (
	Text int8 = iota + 1
	Notice
)

type PushMsgBo struct {
	Key  string
	Op   int32
	Seq  int32
	Body string

	u *user.User
}

type PullMsgBo struct {
	SendUserId  string `json:"send_user_id"`
	MsgId       string `json:"msg_id"`
	Content     string `json:"content"`
	ContentType int8   `json:"content_type"`
}

type Service interface {
	GetUserService() user.Service
	PushMsg(bo *PushMsgBo) error
	AckMsg(userId string, msgId []string) error
	SaveMsg(do *SaveMsgDo) error
	PullMsg(userId string) ([]*PullMsgBo, error)
}

type service struct {
	userSrv user.Service
	d       Dao
}

func (s *service) PullMsg(userId string) ([]*PullMsgBo, error) {
	unReadMsg, err := s.d.ListUnReadMsg(userId)
	if err != nil {
		return nil, err
	}
	result := make([]*PullMsgBo, 0, len(unReadMsg))
	for _, msg := range unReadMsg {
		result = append(result, &PullMsgBo{
			SendUserId:  msg.SendUserId,
			MsgId:       msg.MsgId,
			Content:     msg.Content,
			ContentType: msg.ContentType,
		})
	}
	return result, nil
}

func NewService(d Dao, userSrv user.Service) Service {
	return &service{
		d:       d,
		userSrv: userSrv,
	}
}

func (s *service) GetUserService() user.Service {
	return s.userSrv
}

func (s *service) SaveMsg(do *SaveMsgDo) error {
	return s.d.SaveMsg(context.Background(), do)
}

func (s *service) canPush(userId, friendId, msg string) (bool, error) {
	// 查询是否好友
	isFriend, err := s.userSrv.IsFriend(userId, friendId)
	if err != nil {
		return false, err
	}
	if !isFriend {
		return false, errors.New(friendId + "不是好友")
	}

	err = s.SaveMsg(&SaveMsgDo{
		SendUserId:    userId,
		ReceiveUserId: friendId,
		Content:       msg,
		ContentType:   Text,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *service) PushMsg(bo *PushMsgBo) error {
	ok, err := s.canPush(bo.u.ID, bo.Key, bo.Body)
	if err != nil {
		return err
	}

	if !ok {
		return nil
	}

	req := &pb.PushMsgReq{
		Keys: []string{bo.Key},
		Proto: &protocal.Proto{
			Op:   bo.Op,
			Seq:  bo.Seq,
			Body: []byte(bo.Body),
		},
	}

	return s.d.PushMsg(context.Background(), req)
}

func (s *service) AckMsg(userId string, msgId []string) error {
	return s.d.AckMsg(userId, msgId)
}
