package server_proxy

import (
	"micro-server/new_helloworld/handler"
	"net/rpc"
)

// 因为注册的服务需要有hello方法，所以可以使用接口定义
type HelloServicer interface {
	Hello(name string, reply *string) error
}

func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(handler.HelloName, srv)
}
