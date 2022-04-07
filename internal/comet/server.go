package comet

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "maoim/api/user"
	"maoim/pkg/redis"
	"time"
)

type Server struct {
	rdb    *redis.Redis
	bucket *Bucket
	userClient pb.UserClient
}

func NewServer(rdb *redis.Redis) *Server {
	return &Server{
		rdb:    rdb,
		bucket: NewBucket(1024),
		userClient: newUserGrpcClient(),
	}
}

func (s *Server) Bucket() *Bucket {
	return s.bucket
}

func newUserGrpcClient() pb.UserClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dial, err := grpc.DialContext(ctx, ":8003", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}...)
	if err != nil {
		panic(err)
	}
	return pb.NewUserClient(dial)
}