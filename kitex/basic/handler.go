package main

import (
	"context"
	"fmt"
	"github.wxx.example/kitex/basic/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	addr, err := getLocalIPAddress()
	if err != nil {
		return nil, nil
	}
	return &api.Response{
		Message: fmt.Sprintf("%s:%s", "Echo from", addr),
	}, nil
}
