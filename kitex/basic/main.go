package main

import (
	"errors"
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
	registry, err := zkregistry.NewZookeeperRegistry([]string{"192.168.2.101:2181"}, 40*time.Second)
	if err != nil {
		fmt.Println("create zookeeper registry error: ", err)
		return
	}
	info := &kregistry.Info{
		ServiceName: "Echo",
		Weight: 10,
	}
	ip, err := getLocalIPAddress()
	if err != nil {
		fmt.Println("get local host ip error: ", err)
		return
	}
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", ip,"8888"))
	if err != nil {
		fmt.Println("create tcp address error: ", err)
		return
	}
	fmt.Println("local ip: ", addr)
	svr := api.NewServer(new(EchoImpl), server.WithRegistry(registry), server.WithRegistryInfo(info),
				server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func getLocalIPAddress() (string, error){
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				// 返回第一个有效的IP地址
				return ipnet.IP.String(), err
			}
		}
	}
	return "", errors.New("No valid ip address")
}

func main() {
	// testBasic()
	// testLimit()
	testRegistryZookeeper()
}
