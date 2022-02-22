package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.wxx.example/thrift/gen-go/demohello"
)

type Greeter struct {
}

func (g *Greeter) SayHello(ctx context.Context, username string) (r string, err error) {
	return "Hello from server.", nil
}

func main() {
	addr := "localhost:9090"

	//protocol
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	//transport
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	//handler
	handler := &Greeter{}

	//transport,no secure
	var err error
	var transport thrift.TServerTransport
	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		fmt.Println("error running server:", err)
	}

	//processor
	processor := demohello.NewHelloWorldServiceProcessor(handler)

	fmt.Println("Starting the simple server... on ", addr)

	//start tcp server
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	err = server.Serve()

	if err != nil {
		fmt.Println("error running server:", err)
	}
}
