package comet

import (
	"maoim/api/message"
	pb "maoim/api/user"
	"maoim/internal/comet/conf"
	mess "maoim/internal/pkg/grpc/message"
	"maoim/internal/pkg/grpc/user"
	"maoim/pkg/redis"
)

type Server struct {
	c *conf.Config
	rdb    *redis.Redis
	bucket *Bucket
	userClient pb.UserClient
	messageClient message.MessageClient
}

func NewServer(rdb *redis.Redis) *Server {
	return &Server{
		rdb:    rdb,
		bucket: NewBucket(1024),
		userClient: user.NewUserGrpcClient(),
		messageClient: mess.NewMessageGrpcClient(),
	}
}

func (s *Server) Bucket() *Bucket {
	return s.bucket
}

