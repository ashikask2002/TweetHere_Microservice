package server

import (
	"fmt"
	"net"
	"tweethere-chat/pkg/config"
	pb "tweethere-chat/pkg/pb/chat"

	"google.golang.org/grpc"
)

type Server struct {
	server  *grpc.Server
	listner net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.ChatServiceServer) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}
	newserver := grpc.NewServer()
	pb.RegisterChatServiceServer(newserver, server)
	return &Server{
		server:  newserver,
		listner: lis,
	}, nil

}
func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50056")
	return c.server.Serve(c.listner)
}
