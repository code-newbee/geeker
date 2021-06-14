package internal

import (
	pb "github.com/code-newbee/protocol/geeker"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GeekerServiceImpl struct {}

const (
	port = ":50001"
)

func Run() {
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatal("tcp listen error")
	}
	s := grpc.NewServer()
	pb.RegisterGreekerServer(s, &GeekerServiceImpl{})
	log.Printf("start svr success, listen %v", port)
	if err := s.Serve(lis); err != nil{
		log.Fatalf("svr start error")
	}
}