package server

import (
	"context"
	"log"
	"net"
	"omdb/config"
	"omdb/movie"
	"omdb/proto"

	pbwrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
)

type rpcServer struct {
	grpc     *grpc.Server
	config   *config.Config
	service  movie.Service
	listener net.Listener
	proto.UnimplementedOMDBServiceServer
}

func NewRPCServer(grpc *grpc.Server, conf *config.Config, listener net.Listener, service movie.Service) Server {
	return &rpcServer{
		config:   conf,
		grpc:     grpc,
		service:  service,
		listener: listener,
	}
}
func (s *rpcServer) Run() {
	log.Printf("RPC Server is starting on Port:%d\n", s.config.RPCPort)
	proto.RegisterOMDBServiceServer(s.grpc, s)

	if err := s.grpc.Serve(s.listener); err != nil {
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
