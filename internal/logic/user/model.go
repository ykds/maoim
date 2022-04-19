package user

import (
	"maoim/pkg/mysql"
)

type User struct {
	mysql.BaseModel
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Password string `json:"password" omitempty:"password"`
}

func (f *User) TableName() string {
	return "tbl_users"
}


type FriendShipApply struct {
	mysql.BaseModel
	UserId      string `json:"user_id"`
	Username    string `json:"username"`
	OtherUserId string `json:"other_user_id"`
	Remark      string `json:"remark"`
	Agree       bool   `json:"agree"`
}

func (f *FriendShipApply) TableName() string {
	return "tbl_friendship_apply"
}


type FriendShip struct {
	mysql.BaseModel
	UserId string `json:"user_id"`
	FUserId string `json:"f_user_id"`
}

func (f *FriendShip) TableName() string {
	return "tbl_friendship"
}