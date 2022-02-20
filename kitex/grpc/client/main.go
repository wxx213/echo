package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.wxx.example/kitex/grpc/kitex_gen/example"
	"github.wxx.example/kitex/grpc/kitex_gen/example/echo"
	"log"
	"time"
)

func main()  {
	c, err := echo.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &example.Request{Message: "my request"}
	resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
