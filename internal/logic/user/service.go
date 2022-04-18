package user

import (
	"fmt"
	"maoim/internal/pkg/utils"
	"math/rand"
	"strconv"
)

type UserVo struct {
	ID string `json:"id"`
	Username string `json:"username"`
}

var _ Service = &service{}

type Service interface {
	Register(username, password string) (*User, error)
	Login(username, password string) (string, error)
	Logout(userId string) error
	Exists(username string) (bool, error)
	GetUser(username string) (*User, error)
	Auth(token string) (*User, error)

	AddFriend(username, friendName string) error
	RemoveFriend(username, friendName string) error
	GetFriends(username string) ([]*UserVo, error)
	IsFriend(userId, friendId string) (bool, error)

	Connect(username string) error
	Disconnect(username string) error

	IsOnline(userId string) (bool, error)
}

type service struct {
	dao Dao
}

func (s *service) IsOnline(userId string) (bool, error) {
	return s.dao.IsOnline(userId)
}

func NewService(d Dao) Service {
	return &service{dao: d}
}

func (s *service) Register(username, password string) (*User, error) {
	u, err := s.dao.LoadUser(username)
	if err != nil {
		return nil, err
	}
	if u.ID != "" {
		return nil, fmt.Errorf("user has registered")
	}

	u = &User{
		ID:       strconv.FormatInt(rand.Int63(), 10),
		Username: username,
		Password: password,
	}
	err = s.dao.SaveUser(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *service) Login(username, password string) (string, error) {
	u, err := s.dao.LoadUser(username)
	if err != nil {
		return "", err
	}
	if password != u.Password {
		return "", fmt.Errorf("密码错误")
	}
	return utils.GenToken(u.ID, u.Username)
}

func (s *service) Logout(userId string) error {
	return s.dao.DeleteUser(userId)
}

func (s *service) Exists(username string) (bool, error) {
	user, err := s.dao.LoadUser(username)
	return user != nil && user.ID != "", err
}

func (s *service) GetUser(username string) (*User, error) {
	return s.dao.LoadUser(username)
}

func (s *service) Auth(token string) (*User, error) {
	if token == "" {
		return nil, fmt.Errorf("token不能为空")
	}
	_, username, err := utils.ValidToken(token)
	if err != nil {
		return nil, fmt.Errorf("token错误")
	}
	return s.GetUser(username)
}

func (s *service) AddFriend(username, friendName string) error {
	user, err := s.GetUser(friendName)
	if err != nil {
		return err
	}
	if user.ID == "" {
		return fmt.Errorf("不存在该用户")
	}

	ok, err := s.IsFriend(username, friendName)
	if err != nil {
		return err
	}
	if ok {
		return fmt.Errorf("已添加该好友")
	}

	err1 := s.dao.AddFriend(username, friendName)
	err2 := s.dao.AddFriend(friendName, username)
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func (s *service) RemoveFriend(username, friendName string) error {
	user, err := s.GetUser(friendName)
	if err != nil {
		return err
	}
	if user.ID == "" {
		return fmt.Errorf("不存在该用户")
	}

	ok, err := s.IsFriend(username, friendName)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("无此好友")
	}

	return s.dao.RemoveFriend(username, friendName)
}

func (s *service) GetFriends(username string) (vos []*UserVo, err error) {
	friends, err := s.dao.GetFriends(username)
	if err != nil {
		return
	}

	vos = make([]*UserVo, 0)
	for _, f := range friends {
		user, err := s.GetUser(f)
		if err != nil {
			continue
		}
		vos = append(vos, &UserVo{
			ID: user.ID,
			Username: user.Username,
		})
	}
	return
}

func (s *service) IsFriend(userId, friendId string) (bool, error) {
	return s.dao.IsFriend(userId, friendId)
}


func (s *service) Disconnect(username string) error {
	return s.dao.SetOffline(username)
}

func (s *service) Connect(username string) error {
	return s.dao.SetOnline(username)
}
