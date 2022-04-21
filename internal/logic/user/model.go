package user

import (
	"maoim/pkg/mysql"
)

type ApplyStatus int8

const (
	WAIT_VERY = iota + 1
	PASS
	EXPIRE
)

type User struct {
	mysql.BaseModel
	Username string `json:"username" gorm:"size:20;unique;not null;comment:用户名"`
	Nickname string `json:"nickname" gorm:"size:20;comment:昵称"`
	Mobile   string `json:"mobile" gorm:"size:11;comment:手机号"`
	Avatar   string `json:"avatar" gorm:"size:255;comment:头像"`
	Password string `json:"password" omitempty:"password" gorm:"size:100;comment:密码"`
}

func (f *User) TableName() string {
	return "tbl_users"
}

type FriendShipApply struct {
	mysql.BaseModel
	UserId        string      `json:"user_id"`
	Username      string      `json:"username"`
	OtherUserId   string      `json:"other_user_id"`
	OtherUsername string      `json:"other_username"`
	Remark        string      `json:"remark"`
	Status        ApplyStatus `json:"status"`
}

func (f *FriendShipApply) TableName() string {
	return "tbl_friendship_apply"
}

type FriendShip struct {
	mysql.BaseModel
	UserId  string `json:"user_id"`
	FUserId string `json:"f_user_id"`
}

func (f *FriendShip) TableName() string {
	return "tbl_friendship"
}
