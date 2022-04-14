package message

import (
	"gorm.io/gorm"
)

type MessageIndex struct {
	gorm.Model
	ID        string `gorm:"primarykey"`
	SendUserId    string `json:"send_user_id"`
	ReceiveUserId string `json:"receive_user_id"`
	Box           int8   `json:"box"`
	Read          int8   `json:"read"`
	MsgId         string `json:"msg_content_id"`
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
