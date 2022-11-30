package main

import (
	"fmt"
	"github.com/ukrainskiys/go-grpc-example/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func main() {
	dial, err := grpc.Dial("localhost:5300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	client := proto.NewEntityServiceClient(dial)

	var ctx = context.Background()
	add, _ := client.Add(ctx, &proto.AddEntityRequest{Name: "test"})
	fmt.Println("Add", add.GetId())
	get, _ := client.Get(ctx, &proto.GetEntityRequest{Id: add.GetId()})
	fmt.Println("Get", get.GetEntity())
	getAll, _ := client.GetAll(ctx, &proto.GetAllRequest{})
	fmt.Println("GetAll", getAll.GetEntities())

	_ = dial.Close()
}
