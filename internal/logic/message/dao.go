package message

import (
	"context"
	cpb "maoim/api/comet"
	"maoim/pkg/mysql"
)

var _ Dao = new(dao)

type Dao interface {
	PushMsg(ctx context.Context, req *cpb.PushMsgReq) error
}

type dao struct {
	cometClient cpb.CometClient
	db *mysql.Mysql
}


func (d *dao) PushMsg(ctx context.Context, req *cpb.PushMsgReq) error {
	_, err := d.cometClient.PushMsg(ctx, req)
	return err
}

func NewDao(cometClient cpb.CometClient, db *mysql.Mysql) Dao {
	return &dao {
		cometClient: cometClient,
		db: db,
	}
}
