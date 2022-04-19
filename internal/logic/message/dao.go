package message

import (
	"context"
	"gorm.io/gorm"
	cpb "maoim/api/comet"
	"maoim/pkg/mysql"
)

var _ Dao = new(dao)

type SaveMsgDo struct {
	SendUserId    string
	ReceiveUserId string
	Content       string
	ContentType   int8
}

type PullMsgDo struct {
	SendUserId  string
	MsgId       string
	Content     string
	ContentType int8
}

type Dao interface {
	PushMsg(ctx context.Context, req *cpb.PushMsgReq) error
	SaveMsg(ctx context.Context, do *SaveMsgDo) error
	AckMsg(userId string, msgId []string) error
	ListUnReadMsg(userId string) ([]*PullMsgDo, error)
}

type dao struct {
	cometClient cpb.CometClient
	db          *mysql.Mysql
}

func (d *dao) ListUnReadMsg(userId string) ([]*PullMsgDo, error) {
	mis := make([]MessageIndex, 0)
	err := d.db.GetDB().Where("user_id = ? AND box = 1 AND is_read = 0", userId).Find(&mis).Error
	if err != nil {
		return nil, err
	}

	result := make([]*PullMsgDo, 0)
	for _, mi := range mis {
		mc := MessageContent{}
		err = d.db.GetDB().Where("id = ?", mi.MsgId).Find(&mc).Error
		if err != nil {
			return nil, err
		}
		result = append(result, &PullMsgDo{
			SendUserId:  mi.OtherSideUserId,
			MsgId:       mi.MsgId,
			Content:     mc.Content,
			ContentType: mc.ContentType,
		})
	}
	return result, nil
}

func (d *dao) SaveMsg(ctx context.Context, do *SaveMsgDo) error {
	return d.db.GetDB().Transaction(func(tx *gorm.DB) (err error) {
		mc := MessageContent{
			Content:     do.Content,
			ContentType: do.ContentType,
		}
		err = tx.Create(&mc).Error
		if err != nil {
			return
		}

		sMi := MessageIndex{
			UserId:          do.SendUserId,
			OtherSideUserId: do.ReceiveUserId,
			Box:             0,
			IsRead:            1,
			MsgId:           mc.ID,
		}
		rMi := MessageIndex{
			UserId:          do.ReceiveUserId,
			OtherSideUserId: do.SendUserId,
			Box:             1,
			MsgId:           mc.ID,
		}
		mis := []MessageIndex{sMi, rMi}
		return tx.Create(&mis).Error
	})
}

func (d *dao) PushMsg(ctx context.Context, req *cpb.PushMsgReq) error {
	_, err := d.cometClient.PushMsg(ctx, req)
	return err
}

func (d *dao) AckMsg(userId string, msgId []string) error {
	return d.db.GetDB().Model(&MessageIndex{}).Where("user_id = ? AND box = 1 AND msg_id in ?", userId, msgId).Update("is_read", 1).Error
}

func NewDao(cometClient cpb.CometClient, db *mysql.Mysql) Dao {
	_ = db.GetDB().AutoMigrate(&MessageIndex{}, &MessageContent{})

	return &dao{
		cometClient: cometClient,
		db:          db,
	}
}
