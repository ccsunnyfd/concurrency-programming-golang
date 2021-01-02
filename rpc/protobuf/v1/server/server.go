package main

import (
	"log"
	"net"
	"net/rpc"

	hello "github.com/ccsunnyfd/concurrency-programming/src/rpc/protobuf/v1/proto"
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

	listener, err := net.Listen("tcp", ":11234")
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
