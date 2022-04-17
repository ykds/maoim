package message

import (
	"gorm.io/gorm"
)

type MessageIndex struct {
	gorm.Model
	ID        string `gorm:"primarykey"`
	UserId    string `json:"user_id"`
	OtherSideUserId string `json:"other_side_user_id"`
	Box           int8   `json:"box" comment:"0:发，1:收"`
	Read          int8   `json:"read"`
	MsgId         string `json:"msg_id"`
}

func (m *MessageIndex) TableName() string {
	return "tbl_message_index"
}

type MessageContent struct {
	gorm.Model
	ID        string `gorm:"primarykey"`
	MsgId       string    `json:"content_id"`
	Content     string    `json:"content"`
	ContentType int8      `json:"content_type"`
}

func (mc *MessageContent) TableName() string {
	return "tbl_message_content"
}
