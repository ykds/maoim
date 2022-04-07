package user

import (
	"fmt"
	"maoim/internal/pkg/utils"
	"math/rand"
)

var _ Service = &service{}

type Service interface {
	Register(username, password string) (*User, error)
	Login(username, password string) (string, error)
	Logout(userId string) error
	Exists(username string) (bool, error)
	GetUser(username string) (*User, error)
	Auth(token string) (*User, error)
}

type service struct {
	dao Dao
}


func NewService(d Dao) Service {
	return &service{dao: d}
}

func (s *service) Register(username, password string) (*User, error) {
	u, err := s.dao.LoadUser(username)
	if err != nil {
		return nil, err
	}
	if u.ID != 0 {
		return nil, fmt.Errorf("user has registered")
	}

	u = &User{
		ID:       rand.Int63(),
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
	return user != nil && user.ID != 0, err
}

func (s *service) GetUser(username string) (*User, error) {
	return s.dao.LoadUser(username)
}

func (s *service) Auth(token string) (*User, error) {
	if token == "" {
		return nil, fmt.Errorf("token不能为空")
	}
	username, err := utils.ValidToken(token)
	if err != nil {
		return nil, fmt.Errorf("token错误")
	}
	return s.GetUser(username)
}