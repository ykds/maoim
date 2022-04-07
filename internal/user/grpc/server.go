package grpc

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"maoim/internal/user"
	"net"
	"strconv"

	pb "maoim/api/user"
)

type Server struct {
	pb.UnimplementedUserServer

	srv user.Service
}


func NewUserGrpcServer(srv user.Service) *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor())),
		)
	pb.RegisterUserServer(server, &Server{srv: srv})
	lis, err := net.Listen("tcp", ":8003")
	if err != nil {
		panic(err)
	}
	go func() {
		if err := server.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return server
}

func (s *Server) GetUserByUsername(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {
	if req.GetUsername() == "" {
		return nil, fmt.Errorf("用户名不能为空")
	}
	u, err := s.srv.GetUser(req.GetUsername())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Id: strconv.FormatInt(u.ID, 10),
		Username: u.Username,
		Password: u.Password,
	}, nil
}

func (s *Server) Auth(ctx context.Context, req *pb.AuthReq) (*pb.AuthReply, error) {
	u, err := s.srv.Auth(req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.AuthReply{
		Id: strconv.FormatInt(u.ID, 10),
		Username: u.Username,
		Password: u.Password,
	}, nil
}