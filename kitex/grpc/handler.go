package main

import (
	"context"
	"fmt"
	"github.wxx.example/kitex/grpc/kitex_gen/example"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *example.Request) (resp *example.Response, err error) {
	// TODO: Your code here...
	return &example.Response{
		Message: fmt.Sprintf("Echo from grpc server"),
	}, nil
}
