package movie

import (
	"omdb/config"
	"omdb/proto"
	"testing"

	"google.golang.org/grpc"
)

var svr Server

func init() {
	cfg := config.New()
	client := NewClient(cfg)
	service := NewService(client)
	rpcServer := grpc.NewServer()
	svr = NewServer(rpcServer, service)
	svr.Run()
}

func Test_server_Search(t *testing.T) {
	param := &proto.Param{
		Keyword: "kill",
		Page:    10,
	}
	_, err := svr.Search(param)
	if err != nil {
		t.Error(err)
	}
	param = &proto.Param{
		Keyword: "kill",
		Page:    0,
	}
	_, err = svr.Search(param)
	if err == nil {
		t.Error("should have error")
	}
}
