package user

import (
	"encoding/json"
	"maoim/pkg/mysql"
	"maoim/pkg/redis"
	"strings"
)

var _ Dao = (*dao)(nil)

type Dao interface {
	SaveUser(u *User) error
	LoadUser(username string) (*User, error)
	DeleteUser(userId string) error

	AddFriend(username, friendName string) error
	RemoveFriend(username, friendName string) error
	GetFriends(username string) ([]string, error)
	IsFriend(username, friendName string) (bool, error)
}

type dao struct {
	rdb *redis.Redis
	db *mysql.Mysql
}

func NewDao(rdb *redis.Redis, db *mysql.Mysql) Dao {
	return &dao{
		rdb: rdb,
		db: db,
	}
}

func (d *dao) SaveUser(u *User) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return d.rdb.HSet(CACHE_USER_MAP, u.Username, string(data))
}

func (d *dao) LoadUser(username string) (*User, error) {
	data, err := d.rdb.HGet(CACHE_USER_MAP, username)
	if err != nil {
		if strings.Contains(err.Error(), "redis: nil") {
			return &User{}, nil
		}
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
	return d.rdb.SAdd(CACHE_FRIENT_LIST+ ":" + username, friendName)
}

func (d *dao) RemoveFriend(username, friendName string) error {
	return d.rdb.SRem(CACHE_FRIENT_LIST+ ":" + username, friendName)
}

func (d *dao) GetFriends(username string) ([]string, error) {
	return d.rdb.SMembers(CACHE_FRIENT_LIST + ":" + username)
}


func (d *dao) IsFriend(username, friendName string) (bool, error) {
	return d.rdb.SIsMember(CACHE_FRIENT_LIST+ ":" + username, friendName)
}