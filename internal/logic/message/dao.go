package message

import (
	"context"
	"gorm.io/gorm"
	cpb "maoim/api/comet"
	"maoim/pkg/mysql"
)

var _ Dao = new(dao)

type SaveMsgDo struct {
	SendUserId string
	ReceiveUserId string
	Content string
	ContentType int8
}

type Dao interface {
	PushMsg(ctx context.Context, req *cpb.PushMsgReq) error
	SaveMsg(ctx context.Context, do *SaveMsgDo) error
}

type dao struct {
	cometClient cpb.CometClient
	db *mysql.Mysql
}

func (d *dao) SaveMsg(ctx context.Context, do *SaveMsgDo) error {
	return d.db.GetDB().Transaction(func(tx *gorm.DB) (err error) {
		defer func() {
			if err != nil {
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}()

		mc := MessageContent{
			Content: do.Content,
			ContentType: do.ContentType,
		}
		err = d.db.GetDB().Create(&mc).Error
		if err != nil {
			return
		}

		sMi := MessageIndex{
			SendUserId: do.SendUserId,
			ReceiveUserId: do.ReceiveUserId,
			Box: 0,
			MsgId: mc.ID,
		}
		rMi := MessageIndex{
			SendUserId: do.ReceiveUserId,
			ReceiveUserId: do.SendUserId,
			Box: 1,
			MsgId: mc.ID,
		}
		mis := []MessageIndex{sMi, rMi}
		return tx.Create(&mis).Error
	})
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
