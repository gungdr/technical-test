package movie

import (
	"context"
	"omdb/proto"

	pbwrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
)

type Server interface {
	Run()
}
type server struct {
	grpc    *grpc.Server
	service Service
	proto.UnimplementedOMDBServiceServer
}

func NewServer(grpc *grpc.Server, service Service) Server {
	return &server{
		grpc:    grpc,
		service: service,
	}
}
func (s *server) Run() {
	proto.RegisterOMDBServiceServer(s.grpc, s)
}

func (s *server) Search(ctx context.Context, param *proto.Param) (*proto.SearchResult, error) {
	searchParam := new(Param)
	searchParam.FromProto(param)
	result, err := s.service.Search(*searchParam)
	if err != nil {
		return nil, err
	}
	return result.ToProto(), nil
}

func (s *server) Get(ctx context.Context, id *pbwrapper.StringValue) (*proto.MovieDetailResult, error) {
	result, err := s.service.Get(IMDBID((*id).Value))
	if err != nil {
		return nil, err
	}
	return result.ToProto(), nil
}
