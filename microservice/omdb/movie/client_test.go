package movie

import (
	"omdb/config"
	"testing"
)

func Test_client_Search(t *testing.T) {
	cfg := config.New()
	client := NewClient(cfg)
	param := Param{
		Keyword: "batman",
		Page:    1,
	}
	_, err := client.Search(param.ToMap())
	if err != nil {
		t.Error(err)
	}

	param = Param{
		Keyword: "batman",
		Page:    0,
	}
	_, err = client.Search(param.ToMap())
	if err == nil {
		t.Error("should have error")
	}
}

func Test_client_Get(t *testing.T) {
	cfg := config.New()
	client := NewClient(cfg)
	id := IMDBID("tt2975590")
	_, err := client.Get(id.ToMap())
	if err != nil {
		t.Error(err)
	}

	id = IMDBID("dadass")
	_, err = client.Get(id.ToMap())
	if err == nil {
		t.Error("should have error")
	}
}
