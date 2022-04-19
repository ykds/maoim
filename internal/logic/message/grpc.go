package message

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	pb "maoim/api/message"
	"net"
)

type Server struct {
	pb.UnimplementedMessageServer

	srv Service
}

func NewMessageGrpcServer(srv Service) *grpc.Server {
	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_recovery.UnaryServerInterceptor())))
	pb.RegisterMessageServer(server, &Server{srv: srv})
	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		panic(err)
	}
	go func() {
		err2 := server.Serve(lis)
		if err2 != nil {
			panic(err)
		}
	}()
	return server
}

func (s *Server) AckMsg(ctx context.Context, req *pb.AckReq) (*pb.AckReply, error) {
	return &pb.AckReply{}, s.srv.AckMsg(req.UserId, req.MsgId)
}
