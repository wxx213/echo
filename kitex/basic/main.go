package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/server"
	api "github.wxx.example/kitex/basic/kitex_gen/api/echo"
	"log"
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

func main() {
	// testBasic()
	testLimit()
}
