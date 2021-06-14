package internal

import (
	"context"
	"github.com/code-newbee/protocol/geeker"
	"log"
)

func (g GeekerServiceImpl) SayHello(ctx context.Context,
	req *geeker.HelloRequest) (*geeker.HelloReply, error) {
	log.Println(req)
	resId := req.Name
	return &geeker.HelloReply{Message: resId}, nil
}