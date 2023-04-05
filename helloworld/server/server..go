package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (h *HelloService) Hello(s string, reply *string) error {
	*reply = "hello, " + s
	return nil
}
func main() {
	// 1. 启动一个server
	listen, _ := net.Listen("tcp", ":5555")
	// 2. 注册一个服务
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 3. 启动服务
	conn, _ := listen.Accept() // 当接收到一个新的请求
	// 交给rpc处理
	rpc.ServeConn(conn)
}
