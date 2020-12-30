package main

import (
	"fmt"
	"log"

	"github.com/ccsunnyfd/concurrency-programming/rpc/protobuf/v1/hello"
)

func main() {
	client, err := hello.DialHelloService("tcp", "localhost:1234")
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
