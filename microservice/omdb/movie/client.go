package movie

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"omdb/config"
	"time"

	resty "github.com/go-resty/resty/v2"
)

type Client interface {
	Search(ctx context.Context, param map[string]string) (*SearchResult, error)
	Get(ctx context.Context, id map[string]string) (*MovieDetailResult, error)
}

type client struct {
	resty  *resty.Client
	config *config.Config
	cache  map[string]string
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
			SetDebug(false),
		cache: make(map[string]string),
	}
}

func (c *client) Search(ctx context.Context, param map[string]string) (*SearchResult, error) {
	param["apikey"] = c.config.OMDBAPIKey
	result := &SearchResult{}
	cacheKey := getKey(param)
	if cacheKey != "" {
		if val, found := c.cache[cacheKey]; found {
			err := json.Unmarshal([]byte(val), &result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
	}
	response, err := c.resty.R().SetQueryParams(param).
		Get("/")
	if err != nil {
		log.Println("Client Search:", err)
		return nil, err
	}

	err = json.Unmarshal([]byte(response.String()), &result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, errors.New(*result.Error)
	}
	c.cache[cacheKey] = response.String()
	return result, nil
}

func (c *client) Get(ctx context.Context, param map[string]string) (*MovieDetailResult, error) {
	param["apikey"] = c.config.OMDBAPIKey
	result := &MovieDetailResult{}
	cacheKey := getKey(param)
	if cacheKey != "" {
		if val, found := c.cache[cacheKey]; found {
			err := json.Unmarshal([]byte(val), &result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
	}
	response, err := c.resty.R().SetQueryParams(param).
		Get("/")
	if err != nil {
		log.Println("Client Get:", err)
		return nil, err
	}
	body := response.String()
	log.Println(body)
	err = json.Unmarshal([]byte(response.String()), &result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, errors.New(*result.Error)
	}
	c.cache[cacheKey] = response.String()
	return result, nil
}

func getKey(param map[string]string) string {
	if val, found := param["s"]; found {
		page := param["page"]
		return fmt.Sprintf("%s:%s", val, page)
	}
	if val, found := param["i"]; found {
		return val
	}
	return ""
}
