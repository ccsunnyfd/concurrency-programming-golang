package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/ccsunnyfd/concurrency-programming/rpc/protobuf/v1/hello"
)

// HelloService is
type HelloService struct{}

// Hello is
func (p *HelloService) Hello(request *hello.String, reply *hello.String) error {
	*reply = hello.String{Value: "hello: " + request.GetValue()}
	return nil
}

func main() {
	server := rpc.NewServer()
	hello.RegisterHelloService(server, &HelloService{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go rpc.ServeConn(conn)
	}
}
