package main

import (
	"context"
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
	return nil
}


