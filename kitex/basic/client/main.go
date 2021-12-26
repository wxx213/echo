package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.wxx.example/kitex/basic/kitex_gen/api"
	"github.wxx.example/kitex/basic/kitex_gen/api/echo"
	"log"
	"time"
)

func testBasic() {
	c, err := echo.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.Request{Message: "my request"}
	resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

func keyFunc(ri rpcinfo.RPCInfo) string {
	return ""
}

func testCircuitBreaker() {
	var opts []client.Option

	cbs := circuitbreak.NewCBSuite(nil)
	cbconfig := circuitbreak.CBConfig{
		Enable: true,
		MinSample: 3,
		ErrRate: 0.5,
	}
	cbs.UpdateInstanceCBConfig(cbconfig)
	opts = append(opts, client.WithHostPorts("0.0.0.0:8888"))
	opts = append(opts, client.WithCircuitBreaker(cbs))

	c, err := echo.NewClient("example", opts...)
	if err != nil {
		log.Fatal(err)
	}
	req := &api.Request{Message: "my request"}

	for {
		resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
		if err != nil {
			fmt.Println("call Echo error: ", err)
			time.Sleep(1 * time.Second)
			continue
		}
		log.Println(resp)
	}
}

func main() {
	// testBasic()
	testCircuitBreaker()
}
