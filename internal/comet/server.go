package comet

import (
	"encoding/json"
	"maoim/pkg/redis"
	"strconv"
)

type Server struct {
	rdb *redis.Redis
	bucket *Bucket
}

func NewServer(rdb *redis.Redis) *Server {
	return &Server{
		rdb: rdb,
		bucket: NewBucket(1024),
	}
}

func (s *Server) Bucket() *Bucket {
	return s.bucket
}

func (s *Server) SaveUser(u *User) error {
	d, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return s.rdb.HSet(CACHE_USER_MAP, strconv.FormatInt(u.ID, 10), string(d))
}

func (s *Server) LoadUser(userId string) (*User, error) {
	d, err := s.rdb.HGet(CACHE_USER_MAP, userId)
	if err != nil {
		return nil, err
	}
	u := &User{}
	err = json.Unmarshal([]byte(d), u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Server) DeleteUser(userId string) error {
	return s.rdb.HDel(CACHE_USER_MAP, userId)
}
