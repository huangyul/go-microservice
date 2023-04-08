package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"micro-server/stream_grpc/proto"
	"sync"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:5555", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 服务端流模式
	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "hhh"})
	for i := 0; i < 10; i++ {
		data, _ := res.Recv()
		fmt.Println(data)
	}
	// 客户端流模式
	putS, _ := c.PostStream(context.Background())
	for i := 0; i < 10; i++ {
		_ = putS.Send(&proto.StreamReqData{Data: fmt.Sprintf("客户端发送的数据：%d\n", i)})
	}
	// 双向流模式
	AllS, _ := c.AllStream(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		AllS.Send(&proto.StreamReqData{Data: "来自客户端的数据"})

	}()
	go func() {
		defer wg.Done()
		d, _ := AllS.Recv()
		fmt.Println(d.Data)
	}()
	wg.Wait()

}
