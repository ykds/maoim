package user

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "maoim/api/user"
	"time"
)

func NewUserGrpcClient() pb.UserClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dial, err := grpc.DialContext(ctx, ":9001", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}...)
	if err != nil {
		panic(err)
	}
	return pb.NewUserClient(dial)
}
