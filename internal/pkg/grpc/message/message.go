package message

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "maoim/api/message"
	"time"
)

func NewMessageGrpcClient() pb.MessageClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dial, err := grpc.DialContext(ctx, ":9002", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}...)
	if err != nil {
		panic(err)
	}
	return pb.NewMessageClient(dial)
}
