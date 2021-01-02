package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/ccsunnyfd/concurrency-programming/src/rpc/helloWorld/v1/hello"
)

func main() {
	rpc.RegisterName("HelloService", new(hello.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error: ", err)
	}

	rpc.ServeConn(conn)
}
