package comet

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "maoim/api/comet"
	"time"
)

func NewCometGrpcClient() pb.CometClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	dial, err := grpc.DialContext(ctx, ":9000", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}...)
	if err != nil {
		panic(err)
	}
	return pb.NewCometClient(dial)
}
