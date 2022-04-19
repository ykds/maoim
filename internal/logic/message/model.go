package message

import (
	"maoim/pkg/mysql"
)

type MessageIndex struct {
	mysql.BaseModel
	UserId          string `json:"user_id"`
	OtherSideUserId string `json:"other_side_user_id"`
	Box             int8   `json:"box" comment:"0:发，1:收"`
	IsRead            int8   `json:"is_read"`
	MsgId           string `json:"msg_id"`
}

func (m *MessageIndex) TableName() string {
	return "tbl_message_index"
}

type MessageContent struct {
	mysql.BaseModel
	MsgId       string `json:"content_id"`
	Content     string `json:"content"`
	ContentType int8   `json:"content_type"`
}

func (mc *MessageContent) TableName() string {
	return "tbl_message_content"
}
