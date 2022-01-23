package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/netpoll"
)

var eventLoop netpoll.EventLoop

func main() {

	listener, err := netpoll.CreateListener("tcp", "0.0.0.0:8888")
	if err != nil {
		panic("create netpoll listener failed")
	}

	eventLoop, _ = netpoll.NewEventLoop(handler)
	eventLoop.Serve(listener)
}

func handler (ctx context.Context, connection netpoll.Connection) error {
	reader, writer := connection.Reader(), connection.Writer()
	// reading
	buf, _ := reader.Next(17)
	// parse the read data
	fmt.Println("data from client: ", string(buf))
	reader.Release()

	// writing
	var write_data = []byte("hello from server")
	// make the write data
	alloc, _ := writer.Malloc(len(write_data))
	copy(alloc, write_data) // write data
	writer.Flush()
	return nil
}


