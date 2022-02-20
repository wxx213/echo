package main

import (
	example "github.wxx.example/kitex/grpc/kitex_gen/example/echo"
	"log"
)

func main() {
	svr := example.NewServer(new(EchoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
