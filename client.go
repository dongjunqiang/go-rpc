package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"os/signal"
)

type Args struct {
	A, B int
}

func main() {
	go start()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	<-sigc
}

func start() {
	client, err := rpc.DialHTTP("tcp", "0.0.0.0:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := Args{7, 8}
	var retVal int
	err = client.Call("RPCObject.GetFirstNumber", args, &retVal)
	if err != nil {
		log.Fatal("RPCObject error:", err)
	}
	fmt.Printf("Return Value first number: %d\n", retVal)
	err = client.Call("RPCObject.GetSecondNumber", args, &retVal)
	if err != nil {
		log.Fatal("RPCObject error:", err)
	}
	fmt.Printf("Return Value second number: %d\n", retVal)
}
