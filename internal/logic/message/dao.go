package message

import (
	"context"
	"github.com/pkg/errors"
	cpb "maoim/api/comet"
	upb "maoim/api/user"
)

var _ Dao = new(dao)

type Dao interface {
	IsFriend(ctx context.Context, username, friendName string) (bool, error)
	PushMsg(ctx context.Context, req *cpb.PushMsgReq) error
}

type dao struct {
	cometClient cpb.CometClient
	userClient  upb.UserClient
}

func (d *dao) IsFriend(ctx context.Context, username, friendName string) (bool, error) {
	friendReply, err := d.userClient.IsFriend(ctx, &upb.IsFriendReq{Username: username, Friendname: friendName})
	if err != nil {
		return false, errors.Wrap(err,"Internal Error")
	}
	return friendReply.IsFriend, nil
}

func (d *dao) PushMsg(ctx context.Context, req *cpb.PushMsgReq) error {
	_, err := d.cometClient.PushMsg(ctx, req)
	return err
}

func NewDao(cometClient cpb.CometClient, userClient upb.UserClient) Dao {
	return &dao {
		cometClient: cometClient,
		userClient: userClient,
	}
}
