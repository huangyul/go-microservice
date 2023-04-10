package main

import (
	"context"
	"google.golang.org/grpc"
	"micro-server/interpreter/proto"
	"net"
)

type Server struct {
}

func (s *Server) Hello(ctx context.Context, req *proto.ServerRequest) (*proto.ServerResponse, error) {
	return &proto.ServerResponse{Reply: "hello," + req.Name}, nil
}

func main() {
	listen, _ := net.Listen("tcp", ":5555")

	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	_ = g.Serve(listen)
}
