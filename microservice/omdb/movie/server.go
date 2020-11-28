package movie

import (
	"context"
	"log"
	"omdb/config"
	"omdb/proto"

	pbwrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
)

type Server interface {
	Run()
}
type server struct {
	grpc    *grpc.Server
	config  *config.Config
	service Service
	proto.UnimplementedOMDBServiceServer
}

func NewServer(grpc *grpc.Server, conf *config.Config, service Service) Server {
	return &server{
		config:  conf,
		grpc:    grpc,
		service: service,
	}
}
func (s *server) Run() {
	log.Printf("Server is starting on Port...%d\n", s.config.Port)
	proto.RegisterOMDBServiceServer(s.grpc, s)
}

func (s *server) Search(ctx context.Context, param *proto.Param) (*proto.SearchResult, error) {
	searchParam := new(Param)
	searchParam.FromProto(param)
	result, err := s.service.Search(ctx, *searchParam)
	if err != nil {
		return nil, err
	}
	return result.ToProto(), nil
}

func (s *server) Get(ctx context.Context, id *pbwrapper.StringValue) (*proto.MovieDetailResult, error) {
	result, err := s.service.Get(ctx, IMDBID((*id).Value))
	if err != nil {
		return nil, err
	}
	return result.ToProto(), nil
}
