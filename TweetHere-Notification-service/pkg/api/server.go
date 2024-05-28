package server

import (
	"fmt"
	"net"
	"tweethere-Notification/pkg/config"
	pb "tweethere-Notification/pkg/pb/noti"

	"google.golang.org/grpc"
)

type Server struct {
	server  *grpc.Server
	listner net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.NotificationServiceServer) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newserver := grpc.NewServer()
	pb.RegisterNotificationServiceServer(newserver, server)
	return &Server{
		server:  newserver,
		listner: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on 50057")
	return c.server.Serve(c.listner)
}
