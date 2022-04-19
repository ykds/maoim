package user

import (
	"maoim/pkg/merror"
)

var (
	HasRegisterErr = merror.New(100000, "该账号已被注册")
	PasswordErrorErr = merror.New(100001, "密码错误")
	LoginFailErr = merror.New(100002, "登录失败")
	TokenEmptyErr = merror.New(100003, "Token不能为空")
	TokenCheckFailErr = merror.New(100004, "Token校验不通过")

	AlreadyFriendErr = merror.New(101000, "已是好友")

)
