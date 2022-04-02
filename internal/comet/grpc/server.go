package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"net"

	"maoim/internal/comet"
	pb "maoim/api/comet"
)

var _ pb.CometServer = &server{}


type server struct {
	pb.UnimplementedCometServer

	srv *comet.Server
}

func New(s *comet.Server) *grpc.Server {
	srv := grpc.NewServer()
	pb.RegisterCometServer(srv, &server{srv: s})
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		panic(err)
	}
	go func() {
		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return srv
}

func (s *server) PushMsg(ctx context.Context, req *pb.PushMsgReq) (*pb.PushMsgReply, error)  {
	if len(req.Keys) == 0 || req.Proto == nil {
		return nil, errors.New("req params is invalid")
	}
	for _, key := range req.Keys {
		ch, err := s.srv.Bucket().GetChannel(key)
		if err != nil {
			return nil, err
		}
		if err := ch.Push(req.Proto); err != nil {
			return nil, err
		}
	}
	return &pb.PushMsgReply{}, nil
}


