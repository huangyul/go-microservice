package handler

const HelloName = "handler/HelloService"

type HelloService struct {
}

func (s *HelloService) Hello(name string, reply *string) error {
	*reply = "hello, " + name
	return nil
}
