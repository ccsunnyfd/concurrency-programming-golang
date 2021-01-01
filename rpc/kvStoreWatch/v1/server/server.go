package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/ccsunnyfd/concurrency-programming/rpc/kvStoreWatch/v1/kvstore"
)

func main() {
	rpc.RegisterName("KVStoreService", kvstore.NewKVStoreService())

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
