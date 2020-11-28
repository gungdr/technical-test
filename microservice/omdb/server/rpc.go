package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"omdb/config"
	"omdb/movie"
	"omdb/proto"

	pbwrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
)

type rpcServer struct {
	grpc    *grpc.Server
	config  *config.Config
	service movie.Service
	proto.UnimplementedOMDBServiceServer
}

func NewRPCServer(grpc *grpc.Server, conf *config.Config, service movie.Service) Server {
	return &rpcServer{
		config:  conf,
		grpc:    grpc,
		service: service,
	}
}
func (s *rpcServer) Run() {
	log.Printf("RPC Server is starting on Port:%d\n", s.config.RPCPort)
	proto.RegisterOMDBServiceServer(s.grpc, s)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.RPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *rpcServer) Search(ctx context.Context, param *proto.Param) (*proto.SearchResult, error) {
	searchParam := new(movie.Param)
	searchParam.FromProto(param)
	result, err := s.service.Search(ctx, *searchParam)
	if err != nil {
		return nil, err
	}
	return result.ToProto(), nil
}

func (s *rpcServer) Get(ctx context.Context, id *pbwrapper.StringValue) (*proto.MovieDetailResult, error) {
	result, err := s.service.Get(ctx, movie.IMDBID((*id).Value))
	if err != nil {
		return nil, err
	}
	return result.ToProto(), nil
}
