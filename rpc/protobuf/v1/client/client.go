package main

import (
	"fmt"
	"log"

	hello "github.com/ccsunnyfd/concurrency-programming/src/rpc/protobuf/v1/proto"
)

func main() {
	client, err := hello.DialHelloService("tcp", "localhost:11234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	var reply hello.String
	err = client.Hello(&hello.String{Value: "hello"}, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
