package server

import (
	"log"
	"net"
	"omdb/config"
	"omdb/proto"

	"google.golang.org/grpc"
)

type rpcServer struct {
	grpc     *grpc.Server
	config   *config.Config
	listener net.Listener
	handler  proto.OMDBServiceServer
}

func NewRPCServer(
	grpc *grpc.Server,
	config *config.Config,
	listener net.Listener,
	rpcHandler proto.OMDBServiceServer) Server {
	return &rpcServer{
		grpc:     grpc,
		config:   config,
		listener: listener,
		handler:  rpcHandler,
	}
}
func (s *rpcServer) Run() {
	log.Printf("RPC Server is starting on Port:%d\n", s.config.RPCPort)
	proto.RegisterOMDBServiceServer(s.grpc, s.handler)

	if err := s.grpc.Serve(s.listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
