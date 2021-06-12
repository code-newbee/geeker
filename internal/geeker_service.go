package internal

import (
	"context"
	"github.com/code-newbee/geeker/entity/errcode"
	"github.com/code-newbee/protocol/geeker"
	"log"
)

func (g GeekerServiceImpl) SayHello(ctx context.Context,
	req *geeker.SayHelloRequest) (*geeker.SayHelloResponse, error) {

	log.Println(req)

	resId := req.RequestId

	res := geeker.SayHelloResponse{
		Code: errcode.SUCCESS,
		Message: resId,
	}

	return &res, nil
}