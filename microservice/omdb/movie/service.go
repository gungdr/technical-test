package movie

import "context"

type Service interface {
	Get(ctx context.Context, id IMDBID) (*MovieDetailResult, error)
	Search(ctx context.Context, param Param) (*SearchResult, error)
}

type service struct {
	client Client
}

func NewService(client Client) Service {
	return &service{
		client: client,
	}
}

func (s *service) Get(ctx context.Context, id IMDBID) (*MovieDetailResult, error) {
	result, err := s.client.Get(ctx, id.ToMap())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *service) Search(ctx context.Context, param Param) (*SearchResult, error) {
	result, err := s.client.Search(ctx, param.ToMap())
	if err != nil {
		return nil, err
	}
	return result, nil
}
