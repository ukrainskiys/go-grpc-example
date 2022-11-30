package main

import (
	"github.com/ukrainskiys/go-grpc-example/proto"
	"github.com/ukrainskiys/go-grpc-example/serverimpl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	instance := serverimpl.Server{}

	proto.RegisterEntityServiceServer(server, &instance)

	if err = server.Serve(listener); err != nil {
		return
	}
}
