package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.wxx.example/thriftgo/gen-go/demohello"
)

func main() {
	addr := "localhost:9090"
	var transport thrift.TTransport
	var err error
	var ctx = context.Background()

	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
	}

	//protocol
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	//no buffered
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("error running client:", err)
	}

	if err := transport.Open(); err != nil {
		fmt.Println("error running client:", err)
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := demohello.NewHelloWorldServiceClient(thrift.NewTStandardClient(iprot, oprot))
	res, err := client.SayHello(ctx, "test_thriftgo")
	if err != nil {
		fmt.Println("client error")
	} else {
		fmt.Println("result: ", res)
	}
}
