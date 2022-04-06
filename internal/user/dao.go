package user

import (
	"encoding/json"
	"maoim/pkg/redis"
	"strconv"
)

var _ Dao = (*dao)(nil)

type Dao interface {
	SaveUser(u *User) error
	LoadUser(username string) (*User, error)
	DeleteUser(userId string) error
}

type dao struct {
	rdb *redis.Redis
}

func NewDao(rdb *redis.Redis) Dao {
	return &dao{rdb: rdb}
}

func (d *dao) SaveUser(u *User) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return d.rdb.HSet(CACHE_USER_MAP, strconv.FormatInt(u.ID, 10), string(data))
}

func (d *dao) LoadUser(username string) (*User, error) {
	data, err := d.rdb.HGet(CACHE_USER_MAP, username)
	if err != nil {
		return nil, err
	}
	u := &User{}
	err = json.Unmarshal([]byte(data), u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (d *dao) DeleteUser(userId string) error {
	return d.rdb.HDel(CACHE_USER_MAP, userId)
}
