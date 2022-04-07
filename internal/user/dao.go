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

	AddFriend(username, friendName string) error
	RemoveFriend(username, friendName string) error
	GetFriends(username string) ([]string, error)
	HasFriend(username, friendName string) (bool, error)
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

func (d *dao) DeleteUser(username string) error {
	return d.rdb.HDel(CACHE_USER_MAP, username)
}

func (d *dao) AddFriend(username, friendName string) error {
	return d.rdb.SAdd(CACHE_FRIENT_LIST + ":" + username, friendName)
}

func (d *dao) RemoveFriend(username, friendName string) error {
	return d.rdb.SRem(CACHE_FRIENT_LIST + ":" + username, friendName)
}

func (d *dao) GetFriends(username string) ([]string, error) {
	return d.rdb.SMembers(CACHE_FRIENT_LIST + ":" + username)
}


func (d *dao) HasFriend(username, friendName string) (bool, error) {
	return d.rdb.SIsMember(CACHE_FRIENT_LIST + ":" + username, friendName)
}