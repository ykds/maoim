package message

import (
	"context"
	cpb "maoim/api/comet"
)

var _ Dao = new(dao)

type Dao interface {
	PushMsg(ctx context.Context, req *cpb.PushMsgReq) error
}

type dao struct {
	cometClient cpb.CometClient
}


func (d *dao) PushMsg(ctx context.Context, req *cpb.PushMsgReq) error {
	_, err := d.cometClient.PushMsg(ctx, req)
	return err
}

func NewDao(cometClient cpb.CometClient) Dao {
	return &dao {
		cometClient: cometClient,
	}
}
