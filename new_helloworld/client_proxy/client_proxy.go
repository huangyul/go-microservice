package client_proxy

import (
	"micro-server/new_helloworld/handler"
	"net/rpc"
)

// HelloServiceStub 代理，使得拥有rpc的方法
type HelloServiceStub struct {
	*rpc.Client
}

func NewClientServiceClient(protcol, address string) HelloServiceStub {
	client, _ := rpc.Dial(protcol, address)
	return HelloServiceStub{client}
}

func (c *HelloServiceStub) Hello(name string, reply *string) error {
	// 这一步是真的去执行hello方法，跟原来client使用的方式一致
	_ = c.Call(handler.HelloName+".Hello", name, reply)
	return nil
}
