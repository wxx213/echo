package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/server"
	zkregistry "github.com/kitex-contrib/registry-zookeeper/registry"
	api "github.wxx.example/kitex/basic/kitex_gen/api/echo"
	"log"
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
	svr := api.NewServer(new(EchoImpl), server.WithRegistry(registry))

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
