package comet

import (
	"maoim/pkg/redis"
)

type Server struct {
	rdb    *redis.Redis
	bucket *Bucket
}

func NewServer(rdb *redis.Redis) *Server {
	return &Server{
		rdb:    rdb,
		bucket: NewBucket(1024),
	}
}

func (s *Server) Bucket() *Bucket {
	return s.bucket
}
