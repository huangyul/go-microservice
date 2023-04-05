package rpc

type Company struct {
	Name    string
	Address string
}

type Employee struct {
	Name    string
	company Company
}

// 模拟rpc打印过程
func rpcPrintln(employee Employee) {
	/*
		客户端：
		1. 建立链接tcp/http
		2. 将employee对象序列化成json字符串-序列化
		3. 发送json字段-调用成功后实际上你接收到的是一个二进制的数据
		4. 等待服务器发送结果
		5. 将服务器返回的数据解析-反序列化
		客户端：
		1. 监听端口
		2. 读取数据 - 二进制
		3. 对数据进行反序列化
		4. 开始业务处理
		5. 将结果序列化
		6. 将二进制数据返回

		在过程中，传输协议不一定是json，还有protobuf
	*/
}

func main() {

}
