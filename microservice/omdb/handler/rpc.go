package handler

import (
	"context"
	"errors"
	"omdb/movie"
	"omdb/proto"

	pbwrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
)

type RPCHandler interface {
}

type rpcHandler struct {
	grpc    *grpc.Server
	service movie.Service
}

func NewRPCHandler(grpc *grpc.Server, service movie.Service) RPCHandler {
	return &rpcHandler{
		grpc:    grpc,
		service: service,
	}
}

func (r *rpcHandler) Search(ctx context.Context, param *proto.Param) (*proto.SearchResult, error) {
	searchParam := new(movie.Param)
	searchParam.FromProto(param)
	if (*searchParam).Valid() == false {
		return nil, errors.New("Invalid Param")
	}
	result, err := r.service.Search(ctx, *searchParam)
	if err != nil {
		return nil, err
	}
	return result.ToProto(), nil
}

func (r *rpcHandler) Get(ctx context.Context, id *pbwrapper.StringValue) (*proto.MovieDetailResult, error) {
	result, err := r.service.Get(ctx, movie.IMDBID((*id).Value))
	if err != nil {
		return nil, err
	}
	return result.ToProto(), nil
}
