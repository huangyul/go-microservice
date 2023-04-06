package main

import (
	"fmt"
	"micro-server/new_helloworld/client_proxy"
)

func main() {
	client := client_proxy.NewClientServiceClient("tcp", ":5555")
	reply := new(string)
	client.Hello("huang", reply)
	fmt.Println(*reply)
}
