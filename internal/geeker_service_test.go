package internal

import (
	"context"
	"fmt"
	"github.com/code-newbee/protocol/geeker"
	"testing"
)

func TestGeekerServiceImpl_SayHello(t *testing.T) {

	req := geeker.SayHelloRequest{
		RequestId: "hello",
	}
	g := GeekerServiceImpl{}
	res, _ := g.SayHello(context.Background(), &req)

	fmt.Print(res)
}