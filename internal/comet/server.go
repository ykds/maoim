package comet

import (
	pb "maoim/api/user"
	"maoim/internal/comet/conf"
	"maoim/internal/pkg/grpc/user"
	"maoim/pkg/redis"
)

type Server struct {
	c *conf.Config
	rdb    *redis.Redis
	bucket *Bucket
	userClient pb.UserClient
}

func NewServer(rdb *redis.Redis) *Server {
	return &Server{
		rdb:    rdb,
		bucket: NewBucket(1024),
		userClient: user.NewUserGrpcClient(),
	}
}

func (s *Server) Bucket() *Bucket {
	return s.bucket
}

