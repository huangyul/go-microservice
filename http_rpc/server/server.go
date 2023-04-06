package server

type HelloService struct {
}

func (h *HelloService) Hello(name string, reply *string) error {
	*reply = "hello, " + name
	return nil
}

func main() {}
