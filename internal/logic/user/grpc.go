package user

import (
	"context"
	"errors"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	pb "maoim/api/user"
	"net"
)

type Server struct {
	pb.UnimplementedUserServer

	srv Service
}

func NewUserGrpcServer(srv Service) *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor())),
	)
	pb.RegisterUserServer(server, &Server{srv: srv})
	lis, err := net.Listen("tcp", ":9001")
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
	u, err := s.srv.GetUserByUsername(req.GetUsername())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Id:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}, nil
}

func (s *Server) Connect(ctx context.Context, req *pb.ConnectReq) (reply *pb.ConnectReply, err error) {
	if req.GetUserId() == "" {
		return nil, errors.New("参数错误")
	}

	err = s.srv.Connect(req.GetUserId())
	if err != nil {
		return nil, err
	}
	user, err := s.srv.GetUser(req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pb.ConnectReply{
		UserId:   user.ID,
		UserName: user.Username,
	}, nil
}

func (s *Server) Disconnect(ctx context.Context, req *pb.DisconnectReq) (*pb.DisconnectReply, error) {
	if req.GetUserId() == "" {
		return nil, errors.New("参数错误")
	}
	return &pb.DisconnectReply{}, s.srv.Disconnect(req.GetUserId())
}

//func (s *Server) Auth(ctx context.Context, req *pb.AuthReq) (*pb.AuthReply, merror) {
//	u, err := s.srv.Auth(req.Token)
//	if err != nil {
//		return nil, err
//	}
//	return &pb.AuthReply{
//		Id: strconv.FormatInt(u.ID, 10),
//		Username: u.Username,
//		Password: u.Password,
//	}, nil
//}
//
//func (s *Server) IsFriend(ctx context.Context, req *pb.IsFriendReq) (*pb.IsFriendReply, merror) {
//	if req.GetFriendname() == "" || req.GetUsername() == ""{
//		return nil, fmt.Errorf("参数错误")
//	}
//	is, err := s.srv.IsFriend(req.GetUsername(), req.GetFriendname())
//	if err != nil {
//		return nil, err
//	}
//	return &pb.IsFriendReply{
//		IsFriend: is,
//	}, nil
//}
