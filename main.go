package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"strconv"
)

type Args struct {
	A, B int
}

type RPCObject struct {
	number1 int
	number2 string
}

func (t *RPCObject) GetFirstNumber(args *Args, retVal *int) error {
	fmt.Printf("Serve number1: %d\n", t.number1)
	*retVal = t.number1
	return nil
}

func (t *RPCObject) GetSecondNumber(args *Args, retVal *int) error {
	fmt.Printf("Serve number2: %s\n", t.number2)
	tmp, err := strconv.Atoi(t.number2)
	if err != nil {
		log.Fatal(err)
	}
	*retVal = tmp
	return nil
}

func main() {
	go start()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	<-sigc
}

func start() {
	rpcobject := new(RPCObject)
	rpcobject.number1 = 13
	rpcobject.number2 = "27"
	rpc.Register(rpcobject)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "0.0.0.0:1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}
