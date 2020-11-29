package handler

import (
	"context"
	"omdb/config"
	"omdb/movie"
	"omdb/proto"

	pbwrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type rpcHandler struct {
	config  *config.Config
	service movie.Service
	proto.UnimplementedOMDBServiceServer
}

func NewRPCHandler(
	config *config.Config,
	service movie.Service) proto.OMDBServiceServer {
	return &rpcHandler{
		config:  config,
		service: service,
	}
}
func (s *rpcHandler) Search(ctx context.Context, param *proto.Param) (*proto.SearchResult, error) {
	searchParam := new(movie.Param)
	searchParam.FromProto(param)
	if searchParam.Valid() == false {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Search Param")
	}
	result, err := s.service.Search(ctx, *searchParam)
	if err != nil {
		return nil, err
	}
	return result.ToProto(), nil
}

func (s *rpcHandler) Get(ctx context.Context, id *pbwrapper.StringValue) (*proto.MovieDetailResult, error) {
	result, err := s.service.Get(ctx, movie.IMDBID((*id).Value))
	if err != nil {
		return nil, err
	}
	return result.ToProto(), nil
}
