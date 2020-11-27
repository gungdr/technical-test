package omdb

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type Client interface {
	Search(param map[string]interface{}) (*SearchResult, error)
	Get(id string) (*Movie, error)
}

type client struct {
	resty *resty.Client
}

func NewClient(url string) Client {
	return &client{
		resty: resty.New().
			SetHostURL(url).
			SetRetryCount(3).
			SetTimeout(30 * time.Second).
			SetRetryWaitTime(100 * time.Millisecond).
			SetRetryMaxWaitTime(500 * time.Millisecond).
			SetDebug(true),
	}
}

func (c *client) Search(param map[string]interface{}) (*SearchResult, error) {
	return nil, nil
}

func (c *client) Get(id string) (*Movie, error) {
	return nil, nil
}
