package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/server"
	kregistry "github.com/cloudwego/kitex/pkg/registry"
	zkregistry "github.com/kitex-contrib/registry-zookeeper/registry"
	api "github.wxx.example/kitex/basic/kitex_gen/api/echo"
	"log"
	"net"
	"time"
)

func testBasic() {
	svr := api.NewServer(new(EchoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func testLimit()  {
	lopt := &limit.Option{
		MaxConnections: 100,
		MaxQPS: 100,
	}
	svr := api.NewServer(new(EchoImpl), server.WithLimit(lopt))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func testRegistryZookeeper() {
	registry, err := zkregistry.NewZookeeperRegistry([]string{"127.0.0.1:2181"}, 40*time.Second)
	if err != nil {
		fmt.Println("create zookeeper registry error: ", err)
		return
	}
	info := &kregistry.Info{
		ServiceName: "Echo",
		Weight: 10,
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("create tcp address error: ", err)
		return
	}
	svr := api.NewServer(new(EchoImpl), server.WithRegistry(registry), server.WithRegistryInfo(info),
				server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	// testBasic()
	// testLimit()
	testRegistryZookeeper()
}
