package movie

import (
	"encoding/json"
	"errors"
	"log"
	"omdb/config"
	"time"

	resty "github.com/go-resty/resty/v2"
)

type Client interface {
	Search(param map[string]string) (*SearchResult, error)
	Get(id map[string]string) (*MovieDetailResult, error)
}

type client struct {
	resty  *resty.Client
	config *config.Config
}

func NewClient(conf *config.Config) Client {
	return &client{
		config: conf,
		resty: resty.New().
			SetHostURL(conf.OMDBHost).
			SetRetryCount(3).
			SetTimeout(30 * time.Second).
			SetRetryWaitTime(100 * time.Millisecond).
			SetRetryMaxWaitTime(500 * time.Millisecond).
			SetDebug(true),
	}
}

func (c *client) Search(param map[string]string) (*SearchResult, error) {
	param["apikey"] = c.config.OMDBAPIKey
	response, err := c.resty.R().SetQueryParams(param).
		Get("/")
	if err != nil {
		log.Println("Client Search:", err)
		return nil, err
	}
	result := &SearchResult{}
	err = json.Unmarshal([]byte(response.String()), &result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, errors.New(*result.Error)
	}
	return result, nil
}

func (c *client) Get(param map[string]string) (*MovieDetailResult, error) {
	param["apikey"] = c.config.OMDBAPIKey
	response, err := c.resty.R().SetQueryParams(param).
		Get("/")
	if err != nil {
		log.Println("Client Get:", err)
		return nil, err
	}
	result := &MovieDetailResult{}
	err = json.Unmarshal([]byte(response.String()), &result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, errors.New(*result.Error)
	}
	return result, nil
}
