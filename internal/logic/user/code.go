package user

import (
	"maoim/pkg/merror"
)

var (
	HasRegisterErr    = merror.New(100000, "该账号已被注册")
	RegisterErr       = merror.New(100001, "注册失败")
	UserNotFound      = merror.New(100002, "无此用户")
	PasswordErrorErr  = merror.New(100003, "密码错误")
	LoginFailErr      = merror.New(100004, "登录失败")
	TokenEmptyErr     = merror.New(100005, "Token不能为空")
	TokenCheckFailErr = merror.New(100006, "Token校验不通过")

	AlreadyFriendErr = merror.New(101000, "已是好友")
)
