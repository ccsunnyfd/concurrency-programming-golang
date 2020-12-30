package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/ccsunnyfd/concurrency-programming/rpc/helloWorld/v2/hello"
)

// HelloService is
type HelloService struct{}

// Hello is
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}

func main() {
	hello.RegisterHelloService(new(HelloService))

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
