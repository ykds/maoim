package user

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
)

var _ Service = &service{}

type Service interface {
	Register(username, password string) (*User, error)
	Login(username, password string) (string, error)
	Logout(userId string) error
	Exists(username string) (bool, error)
	GetUser(username string) (*User, error)
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
	cookie := map[string]string{"username": u.Username}
	ck, err := json.Marshal(cookie)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ck), nil
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
