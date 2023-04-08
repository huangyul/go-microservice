package main

import (
	"fmt"
	"google.golang.org/grpc"
	"micro-server/stream_grpc/proto"
	"net"
	"strconv"
	"sync"
)

const PORT = ":5555"
const NETWORK = "tcp"

// 实现三种流模式的方法

type server struct{}

// GetStream 服务流
func (s *server) GetStream(req *proto.StreamReqData, streamServer proto.Greeter_GetStreamServer) error {
	for i := 0; i < 10; i++ {
		streamServer.Send(&proto.StreamResData{
			Data: strconv.Itoa(i),
		})
	}
	return nil
}

// PostStream  客户端流
func (s *server) PostStream(client proto.Greeter_PostStreamServer) error {
	for {
		data, err := client.Recv()
		if err != nil {
			break
		} else {
			fmt.Println(data.Data)
		}
	}
	return nil
}

// AllStream 双向流
func (s *server) AllStream(client proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		res, _ := client.Recv()
		fmt.Println("接收到客户端的数据")
		fmt.Println(res.Data)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("发送数据给客户端")
		client.Send(&proto.StreamResData{Data: "来自服务端的数据"})
	}()
	wg.Wait()
	return nil
}

func main() {
	listen, err := net.Listen(NETWORK, PORT)
	if err != nil {
		fmt.Println("network error")
		panic(err)
	}
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &server{})
	err = g.Serve(listen)
	if err != nil {
		panic("grpc error")
	}
}
