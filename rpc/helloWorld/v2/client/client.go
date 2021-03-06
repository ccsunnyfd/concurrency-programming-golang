package main

import (
	"fmt"
	"log"

	"github.com/ccsunnyfd/concurrency-programming/src/rpc/helloWorld/v2/hello"
)

func main() {
	client, err := hello.DialHelloService("tcp", "localhost:11234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
