package server

import (
	"fmt"
	"net"
	"tweet-service/pkg/config"
	pb "tweet-service/pkg/pb/tweet"

	"google.golang.org/grpc"
)

type Server struct {
	server  *grpc.Server
	listner net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.TweetServiceServer) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newserver := grpc.NewServer()
	pb.RegisterTweetServiceServer(newserver, server)
	return &Server{
		server:  newserver,
		listner: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50055")
	return c.server.Serve(c.listner)
}
