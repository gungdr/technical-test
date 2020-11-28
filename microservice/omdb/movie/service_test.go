package movie

import (
	"context"
	"omdb/config"
	"testing"
)

func Test_service_Search(t *testing.T) {
	cfg := config.New()
	client := NewClient(cfg)
	service := NewService(client)
	param := Param{
		Keyword: "batman",
		Page:    1,
	}
	_, err := service.Search(context.Background(), param)
	if err != nil {
		t.Error(err)
	}
	param = Param{
		Keyword: "batman",
		Page:    0,
	}
	_, err = service.Search(context.Background(), param)
	if err == nil {
		t.Error("should have error")
	}
}

func Test_service_Get(t *testing.T) {
	cfg := config.New()
	client := NewClient(cfg)
	service := NewService(client)
	id := IMDBID("tt2975590")
	_, err := service.Get(context.Background(), id)
	if err != nil {
		t.Error(err)
	}

	id = IMDBID("dadass")
	_, err = service.Get(context.Background(), id)
	if err == nil {
		t.Error("should have error")
	}
}
