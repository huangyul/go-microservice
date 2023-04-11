package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"micro-server/interpreter/proto"
	"net"
)

type Server struct {
}

func (s *Server) Hello(ctx context.Context, req *proto.ServerRequest) (*proto.ServerResponse, error) {
	return &proto.ServerResponse{Reply: "hello," + req.Name}, nil
}

func main() {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("拦截器响应")
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "没有验证信息")
		}
		// 从metadata中取出验证信息
		var (
			appid  string
			appkey string
		)
		if va1, ok := md["appid"]; ok {
			appid = va1[0]
		}
		if va2, ok := md["appkey"]; ok {
			appkey = va2[0]
		}
		if appid != "10109" || appkey != "huang" {
			return resp, status.Error(codes.Unauthenticated, "信息不正确")
		}
		// 原样执行
		return handler(ctx, req)
	}
	// 添加拦截器
	opt := grpc.UnaryInterceptor(interceptor)
	listen, _ := net.Listen("tcp", "127.0.0.1:5555")

	g := grpc.NewServer(opt) // 多个拦截器  opt,opt2,opt3
	proto.RegisterGreeterServer(g, &Server{})
	_ = g.Serve(listen)
}
