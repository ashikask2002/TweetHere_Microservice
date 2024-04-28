package server

import (
	"Tweethere-Auth/pkg/config"
	"Tweethere-Auth/pkg/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	server  *grpc.Server
	listner net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.AuthServiceServer) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newserver := grpc.NewServer()
	pb.RegisterAuthServiceServer(newserver, server)
	return &Server{
		server:  newserver,
		listner: lis,
	}, nil

}
func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50002 ")
	return c.server.Serve(c.listner)
}
