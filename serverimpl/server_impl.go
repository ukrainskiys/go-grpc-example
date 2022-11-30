package serverimpl

import (
	"github.com/ukrainskiys/go-grpc-example/proto"
	"golang.org/x/net/context"
)

var storage = make(map[int64]Entity)

type Entity struct {
	Id   int64
	Name string
}

type Server struct {
	proto.UnimplementedEntityServiceServer
}

func (s *Server) Add(_ context.Context, in *proto.AddEntityRequest) (*proto.AddEntityResponse, error) {
	id := int64(len(storage) + 1)
	storage[id] = Entity{Name: in.GetName(), Id: id}
	return &proto.AddEntityResponse{Id: id}, nil
}

func (s *Server) Get(_ context.Context, in *proto.GetEntityRequest) (*proto.GetEntityResponse, error) {
	entity := storage[in.GetId()]
	return &proto.GetEntityResponse{
		Entity: &proto.Entity{Id: entity.Id, Name: entity.Name},
	}, nil
}

func (s *Server) GetAll(context.Context, *proto.GetAllRequest) (*proto.GetAllResponse, error) {
	var entities []*proto.Entity
	for _, e := range storage {
		entities = append(entities, &proto.Entity{Id: e.Id, Name: e.Name})
	}
	return &proto.GetAllResponse{Entities: entities}, nil
}
